package handlers

// Defines the checkpoint-restore (CRIU) handlers that ship with the server

import (
	"context"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"syscall"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"github.com/cedana/cedana/pkg/criu"
	"github.com/cedana/cedana/pkg/keys"
	"github.com/cedana/cedana/pkg/types"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	CRIU_LOG_VERBOSITY_LEVEL = 4
	CRIU_LOG_FILE            = "criu.log"
	GHOST_FILE_MAX_SIZE      = 10000000 // 10MB
)

// Returns a CRIU dump handler for the server
func DumpCRIU() types.Dump {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.DumpResp, req *daemon.DumpReq) (exited chan int, err error) {
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
		criuOpts.Pid = proto.Int32(int32(resp.GetState().GetPID()))
		criuOpts.NotifyScripts = proto.Bool(true)

		// TODO: Add support for pre-dump
		// TODO: Add support for lazy migration

		log.Debug().Int("CRIU", version).Msg("CRIU dump starting")
		// utils.LogProtoMessage(criuOpts, "CRIU option", zerolog.DebugLevel)

		// Capture internal logs from CRIU
		go utils.LogFromFile(
			log.With().Int("CRIU", version).Logger().WithContext(ctx),
			filepath.Join(criuOpts.GetImagesDir(), CRIU_LOG_FILE),
			zerolog.TraceLevel,
		)

		_, err = server.CRIU.Dump(ctx, criuOpts, nfy)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed CRIU dump: %v", err)
		}

		log.Debug().Int("CRIU", version).Msg("CRIU dump complete")

		return utils.WaitForPid(resp.State.PID), nil
	}
}

// Returns a CRIU restore handler for the server
func RestoreCRIU() types.Restore {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.RestoreResp, req *daemon.RestoreReq) (chan int, error) {
		extraFiles := utils.GetContextValSafe(
			ctx,
			keys.RESTORE_EXTRA_FILES_CONTEXT_KEY,
			[]*os.File{},
		)
		ioFiles := utils.GetContextValSafe(ctx, keys.RESTORE_IO_FILES_CONTEXT_KEY, []*os.File{})

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
		criuOpts.OrphanPtsMaster = proto.Bool(false)

		log.Debug().Int("CRIU", version).Msg("CRIU restore starting")
		// utils.LogProtoMessage(criuOpts, "CRIU option", zerolog.DebugLevel)

		// Capture internal logs from CRIU
		go utils.LogFromFile(
			log.With().Int("CRIU", version).Logger().WithContext(ctx),
			filepath.Join(criuOpts.GetImagesDir(), CRIU_LOG_FILE),
			zerolog.TraceLevel,
		)

		// Attach IO if requested, otherwise log to file
		exitCode := make(chan int, 1)
		var inWriter, outReader, errReader *os.File
		if req.Attach {
			if len(ioFiles) != 3 {
				return nil, status.Error(codes.Internal, "ioFiles did not contain 3 files")
			}
			inWriter, outReader, errReader = ioFiles[0], ioFiles[1], ioFiles[2]
			// Use a random number, since we don't have PID yet
			id := rand.Uint32()
			stdIn, stdOut, stdErr := utils.NewStreamIOSlave(server.Lifetime, id, exitCode)
			defer utils.SetIOSlavePID(id, &resp.PID) // PID should be available then
			go io.Copy(inWriter, stdIn)
			go io.Copy(stdOut, outReader)
			go io.Copy(stdErr, errReader)
		}

		criuResp, err := server.CRIU.Restore(ctx, criuOpts, nfy, extraFiles...)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed CRIU restore: %v", err)
		}
		resp.PID = uint32(*criuResp.Pid)

		// If restoring as child of daemon (RstSibling), we need wait to close the exited channel
		// as their could be goroutines waiting on it.
		exited := make(chan int)
		if criuOpts.GetRstSibling() {
			server.WG.Add(1)
			go func() {
				defer server.WG.Done()
				var status syscall.WaitStatus
				_, err := syscall.Wait4(int(resp.PID), &status, 0, nil)
				if err != nil {
					log.Debug().Err(err).Msg("process Wait4()")
				}
				code := status.ExitStatus()
				log.Debug().Int("code", status.ExitStatus()).Msg("process exited")
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
		} else {
			close(exitCode)
			close(exited)
		}

		log.Debug().Int("CRIU", version).Msg("CRIU restore complete")

		return exited, err
	}
}