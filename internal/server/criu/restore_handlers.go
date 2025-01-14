package criu

import (
	"context"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"syscall"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	cedana_io "github.com/cedana/cedana/pkg/io"
	"github.com/cedana/cedana/pkg/keys"
	"github.com/cedana/cedana/pkg/profiling"
	"github.com/cedana/cedana/pkg/types"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var Restore types.Restore = restore

// Returns a CRIU restore handler for the server
func restore(ctx context.Context, server types.ServerOpts, resp *daemon.RestoreResp, req *daemon.RestoreReq) (chan int, error) {
	extraFiles, ok := ctx.Value(keys.EXTRA_FILES_CONTEXT_KEY).([]*os.File)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "failed to get extra files from context")
	}

	if req.GetCriu() == nil {
		return nil, status.Error(codes.InvalidArgument, "criu options is nil")
	}

	version, err := server.CRIU.GetCriuVersion(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get CRIU version: %v", err)
	}

	criuOpts := req.GetCriu()

	// Set CRIU server
	criuOpts.LogFile = proto.String(CRIU_LOG_FILE)
	criuOpts.LogLevel = proto.Int32(CRIU_LOG_VERBOSITY_LEVEL)
	criuOpts.GhostLimit = proto.Uint32(GHOST_FILE_MAX_SIZE)
	criuOpts.NotifyScripts = proto.Bool(true)
	criuOpts.LogToStderr = proto.Bool(false)

	log.Debug().Int("CRIU", version).Interface("opts", criuOpts).Msg("CRIU restore starting")
	// utils.LogProtoMessage(criuOpts, "CRIU option", zerolog.DebugLevel)

	// Attach IO if requested, otherwise log to file
	exitCode := make(chan int, 1)
	var stdin io.Reader
	var stdout, stderr io.Writer
	if req.Attachable {
		criuOpts.RstSibling = proto.Bool(true) // restore as child, so we can wait for the exit code
		id := rand.Uint32()                    // Use a random number, since we don't have PID yet
		stdin, stdout, stderr = cedana_io.NewStreamIOSlave(server.Lifetime, server.WG, id, exitCode)
		defer cedana_io.SetIOSlavePID(id, &resp.PID) // PID should be available then
	} else {
		logFile, ok := ctx.Value(keys.LOG_FILE_CONTEXT_KEY).(*os.File)
		if ok {
			stdout, stderr = logFile, logFile
		}
	}

	ctx, end := profiling.StartTimingCategory(ctx, "criu", server.CRIU.Restore)

	criuResp, err := server.CRIU.Restore(
		ctx,
		criuOpts,
		server.CRIUCallback,
		stdin,
		stdout,
		stderr,
		extraFiles...)

	end()

	// Capture internal logs from CRIU
	utils.LogFromFile(
		log.With().Int("CRIU", version).Logger().WithContext(ctx),
		filepath.Join(criuOpts.GetImagesDir(), CRIU_LOG_FILE),
		zerolog.TraceLevel,
	)

	if err != nil {
		// NOTE: It's possible that after the restore failed, the process
		// exists as a zombie process. We need to reap it.
		if pid := resp.GetState().GetPID(); pid != 0 {
			p, _ := os.FindProcess(int(pid))
			p.Wait()
		}

		return nil, status.Errorf(codes.Internal, "failed CRIU restore: %v", err)
	}
	resp.PID = uint32(*criuResp.Pid)

	// If restoring as child of daemon (RstSibling), we need wait to close the exited channel
	// as their could be goroutines waiting on it.
	var exited chan int
	if criuOpts.GetRstSibling() {
		exited = make(chan int)
		server.WG.Add(1)
		go func() {
			defer server.WG.Done()
			p, _ := os.FindProcess(int(resp.PID)) // always succeeds on linux
			status, err := p.Wait()
			if err != nil {
				log.Debug().Err(err).Msg("process Wait()")
			}
			code := status.ExitCode()
			log.Debug().Uint8("code", uint8(code)).Msg("process exited")
			exitCode <- code
			close(exitCode)
			close(exited)
		}()

		// Also kill the process if it's lifetime expires
		server.WG.Add(1)
		go func() {
			defer server.WG.Done()
			<-server.Lifetime.Done()
			syscall.Kill(int(resp.PID), syscall.SIGKILL)
		}()
	}

	log.Debug().Int("CRIU", version).Msg("CRIU restore complete")

	return exited, err
}