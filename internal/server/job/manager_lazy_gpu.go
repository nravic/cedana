package job

// Extends the lazy job manager to support GPU management

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"buf.build/gen/go/cedana/cedana-gpu/grpc/go/gpu/gpugrpc"
	"buf.build/gen/go/cedana/cedana-gpu/protocolbuffers/go/gpu"
	"buf.build/gen/go/cedana/criu/protocolbuffers/go/criu"
	"github.com/cedana/cedana/pkg/plugins"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	GPU_CONTROLLER_HOST               = "localhost"
	GPU_CONTROLLER_LOG_PATH_FORMATTER = "/tmp/cedana-gpu-controller-%s.log"
	GPU_CONTROLLER_LOG_FLAGS          = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	GPU_CONTROLLER_LOG_PERMS          = 0644

	GPU_HEALTH_TIMEOUT  = 30 * time.Second
	GPU_DUMP_TIMEOUT    = 5 * time.Minute
	GPU_RESTORE_TIMEOUT = 5 * time.Minute

	// Signal sent to job when GPU controller exits prematurely. The intercepted job
	// is guaranteed to exit upon receiving this signal, and prints to stderr
	// about the GPU controller's failure.
	GPU_CONTROLLER_PREMATURE_EXIT_SIGNAL = syscall.SIGUSR1
)

type gpuController struct {
	JID    string
	ready  chan bool
	cmd    *exec.Cmd
	client gpugrpc.ControllerClient
	conn   *grpc.ClientConn
	stderr *bytes.Buffer
}

func (m *ManagerLazy) AttachGPU(
	ctx context.Context,
	lifetime context.Context,
	jid string,
) error {
	// Check if GPU plugin is installed
	var gpu *plugins.Plugin
	if gpu = m.plugins.Get("gpu"); gpu.Status != plugins.Installed {
		return fmt.Errorf("Please install the GPU plugin to use GPU support")
	}
	binary := gpu.BinaryPaths()[0]

	if _, err := os.Stat(binary); err != nil {
		return err
	}

	job := m.Get(jid)
	job.SetGPUEnabled(true)
	m.pending <- action{update, job}

	port, err := utils.GetFreePort()
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(lifetime, binary, jid, "--port", strconv.Itoa(port))

	stdout, err := os.OpenFile(
		fmt.Sprintf(GPU_CONTROLLER_LOG_PATH_FORMATTER, jid),
		GPU_CONTROLLER_LOG_FLAGS,
		GPU_CONTROLLER_LOG_PERMS)
	if err != nil {
		return fmt.Errorf("failed to open GPU controller log file: %w", err)
	}

	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGTERM,
	}
	cmd.Cancel = func() error { return cmd.Process.Signal(syscall.SIGTERM) } // NO SIGKILL!!!

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf(
			"failed to start GPU controller: %w",
			utils.GRPCErrorShort(err, stderr.String()),
		)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", GPU_CONTROLLER_HOST, port), opts...)
	if err != nil {
		cmd.Process.Signal(syscall.SIGTERM)
		return fmt.Errorf(
			"failed to create GPU controller client: %w",
			utils.GRPCErrorShort(err, stderr.String()),
		)
	}
	client := gpugrpc.NewControllerClient(conn)

	gpuController := &gpuController{
		JID:    jid,
		cmd:    cmd,
		client: client,
		conn:   conn,
		stderr: stderr,
	}
	m.gpuControllers.Store(jid, gpuController)

	// Cleanup controller on exit, and signal job of its exit

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		err := cmd.Wait()
		if err != nil {
			log.Trace().Err(err).Str("JID", jid).Msg("GPU controller Wait()")
		}
		log.Debug().
			Int("code", cmd.ProcessState.ExitCode()).
			Str("JID", jid).
			Msg("GPU controller exited")

		m.Kill(jid, GPU_CONTROLLER_PREMATURE_EXIT_SIGNAL)
		conn.Close()
		m.gpuControllers.Delete(jid)
	}()

	log.Debug().Str("JID", jid).Int("port", port).Msg("waiting for GPU controller...")

	err = m.checkGPUHealth(ctx, gpuController)
	if err != nil {
		cmd.Process.Signal(syscall.SIGTERM)
		conn.Close()
		return err
	}

	m.addCRIUCallbackGPU(lifetime, jid)

	log.Debug().Str("JID", jid).Msg("GPU controller ready")

	return nil
}

func (m *ManagerLazy) AttachGPUAsync(
	ctx context.Context,
	lifetime context.Context,
	jid string,
) <-chan error {
	err := make(chan error)

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer close(err)
		for {
			select {
			case <-ctx.Done():
				return
			case err <- m.AttachGPU(ctx, lifetime, jid):
				return
			}
		}
	}()

	return err
}

//////////////////////////
//// Helper Functions ////
//////////////////////////

func (m *ManagerLazy) addCRIUCallbackGPU(lifetime context.Context, jid string) {
	job := m.Get(jid)

	// Add pre-dump hook for GPU dump. This ensures that the GPU is dumped before
	// CRIU freezes the process.
	job.CRIUCallback.PreDumpFunc = func(ctx context.Context, opts *criu.CriuOpts) error {
		err := checkCRIUOptsCompatibilityGPU(opts)
		if err != nil {
			return err
		}

		waitCtx, cancel := context.WithTimeout(ctx, GPU_DUMP_TIMEOUT)
		defer cancel()

		controller := m.getGPUController(jid)

		_, err = controller.client.Checkpoint(waitCtx, &gpu.CheckpointRequest{Directory: opts.GetImagesDir()})
		if err != nil {
			log.Error().Err(err).Str("JID", jid).Msg("failed to dump GPU")
			return err
		}
		log.Info().Str("JID", jid).Msg("GPU dump complete")
		return err
	}

	// Add pre-restore hook for GPU restore, that begins GPU restore in parallel
	// to CRIU restore. We instead block at pre-resume, to maximize concurrency.
	restoreErr := make(chan error)
	job.CRIUCallback.PreRestoreFunc = func(ctx context.Context, opts *criu.CriuOpts) error {
		err := checkCRIUOptsCompatibilityGPU(opts)
		if err != nil {
			return err
		}

		err = m.AttachGPU(ctx, lifetime, jid) // Re-attach a GPU to the job
		if err != nil {
			return err
		}

		go func() {
			defer close(restoreErr)

			waitCtx, cancel := context.WithTimeout(ctx, GPU_RESTORE_TIMEOUT)
			defer cancel()

			controller := m.getGPUController(jid)
			_, err = controller.client.Restore(waitCtx, &gpu.RestoreRequest{Directory: opts.GetImagesDir()})
			if err != nil {
				log.Error().Err(err).Str("JID", jid).Msg("failed to restore GPU")
				restoreErr <- err
				return
			}
			log.Info().Str("JID", jid).Msg("GPU restore complete")
		}()
		return nil
	}

	// Wait for GPU restore to finish before resuming the process
	job.CRIUCallback.PreResumeFunc = func(ctx context.Context, pid int32) error {
		return <-restoreErr
	}
}

// Health checks the GPU controller, blocking on connection until ready.
// This can be used as a proxy to wait for the controller to be ready.
func (m *ManagerLazy) checkGPUHealth(ctx context.Context, controller *gpuController) error {
	waitCtx, cancel := context.WithTimeout(ctx, GPU_HEALTH_TIMEOUT)
	defer cancel()

	// Wait for early controller exit, and cancel the blocking health check
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		<-utils.WaitForPidCtx(waitCtx, uint32(controller.cmd.Process.Pid))
		cancel()
	}()

	resp, err := controller.client.HealthCheck(waitCtx, &gpu.HealthCheckRequest{}, grpc.WaitForReady(true))
	if resp != nil {
		log.Debug().
			Str("JID", controller.JID).
			Int32("devices", resp.DeviceCount).
			Str("version", resp.Version).
			Int32("driver", resp.GetAvailableAPIs().GetDriverVersion()).
			Msg("GPU health check")
	}
	if err != nil || !resp.Success {
		controller.cmd.Process.Signal(syscall.SIGTERM)
		controller.conn.Close()
		if err == nil {
			err = status.Errorf(codes.FailedPrecondition, "GPU health check failed")
			controller.stderr.WriteString("GPU health check failed")
		}
		return utils.GRPCErrorShort(err, controller.stderr.String())
	}
	return nil
}

// Certain CRIU options are not compatible with GPU support.
func checkCRIUOptsCompatibilityGPU(opts *criu.CriuOpts) error {
	if opts.GetLeaveRunning() {
		return fmt.Errorf("leave_running is not compatible with GPU support")
	}
	return nil
}