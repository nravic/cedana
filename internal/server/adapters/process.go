package adapters

// Defines all the adapters that manage the process-level details

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	criu_proto "buf.build/gen/go/cedana/criu/protocolbuffers/go/criu"
	"github.com/cedana/cedana/pkg/criu"
	"github.com/cedana/cedana/pkg/keys"
	"github.com/cedana/cedana/pkg/types"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	STATE_FILE                         = "process_state.json"
	DEFAULT_RESTORE_LOG_PATH_FORMATTER = "/var/log/cedana-restore-%d.log"
	RESTORE_LOG_FLAGS                  = os.O_CREATE | os.O_WRONLY | os.O_APPEND // no truncate
	RESTORE_LOG_PERMS                  = 0o644
)

////////////////////////
//// Dump Adapters /////
////////////////////////

// Fills process state in the dump response.
// Requires at least the PID to be present in the DumpResp.State
// Also saves the state to a file in the dump directory, post dump.
func FillProcessStateForDump(next types.Dump) types.Dump {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.DumpResp, req *daemon.DumpReq) (exited chan int, err error) {
		state := resp.GetState()
		if state == nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"missing state. at least PID is required in resp.state",
			)
		}

		if state.PID == 0 {
			return nil, status.Errorf(
				codes.NotFound,
				"missing PID. Ensure an adapter sets this PID in response.",
			)
		}

		err = utils.FillProcessState(ctx, state.PID, state)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to fill process state: %v", err)
		}

		exited, err = next(ctx, server, nfy, resp, req)
		if err != nil {
			return exited, err
		}

		// Post dump, save the state to a file in the dump
		stateFile := filepath.Join(req.GetCriu().GetImagesDir(), STATE_FILE)
		if err := utils.SaveJSONToFile(state, stateFile); err != nil {
			log.Warn().Err(err).Str("file", stateFile).Msg("failed to save process state")
		}

		return exited, nil
	}
}

// Detect and sets shell job option for CRIU
func DetectShellJobForDump(next types.Dump) types.Dump {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.DumpResp, req *daemon.DumpReq) (exited chan int, err error) {
		var isShellJob bool
		if info := resp.GetState().GetInfo(); info != nil {
			for _, f := range info.GetOpenFiles() {
				if strings.Contains(f.Path, "pts") {
					isShellJob = true
					break
				}
			}
		} else {
			log.Warn().Msg("No process info found. it should have been filled by an adapter")
		}

		if req.GetCriu() == nil {
			req.Criu = &criu_proto.CriuOpts{}
		}

		// Only set unless already set
		if req.GetCriu().ShellJob == nil {
			req.Criu.ShellJob = proto.Bool(isShellJob)
		}

		return next(ctx, server, nfy, resp, req)
	}
}

// Close common file descriptors b/w the parent and child process
func CloseCommonFilesForDump(next types.Dump) types.Dump {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.DumpResp, req *daemon.DumpReq) (exited chan int, err error) {
		pid := resp.GetState().GetPID()
		if pid == 0 {
			return nil, status.Errorf(
				codes.NotFound,
				"missing PID. Ensure an adapter sets this PID in response before.",
			)
		}

		err = utils.CloseCommonFds(ctx, int32(os.Getpid()), int32(pid))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to close common fds: %v", err)
		}

		return next(ctx, server, nfy, resp, req)
	}
}

////////////////////////
/// Restore Adapters ///
////////////////////////

// Fill process state in the restore response
func FillProcessStateForRestore(next types.Restore) types.Restore {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.RestoreResp, req *daemon.RestoreReq) (chan int, error) {
		// Check if path is a directory
		path := req.GetCriu().GetImagesDir()
		if path == "" {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"missing path. should have been set by an adapter",
			)
		}

		stateFile := filepath.Join(path, STATE_FILE)
		state := &daemon.ProcessState{}

		if err := utils.LoadJSONFromFile(stateFile, state); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				"failed to load process state from dump: %v",
				err,
			)
		}

		resp.State = state

		exited, err := next(ctx, server, nfy, resp, req)
		if err != nil {
			return exited, err
		}

		// Try to update the process state with the latest information,
		// Only possible if process is still running, otherwise ignore.
		_ = utils.FillProcessState(ctx, state.PID, state)

		return exited, err
	}
}

// Detect and sets shell job option for CRIU
func DetectShellJobForRestore(next types.Restore) types.Restore {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.RestoreResp, req *daemon.RestoreReq) (chan int, error) {
		var isShellJob bool
		if info := resp.GetState().GetInfo(); info != nil {
			for _, f := range info.GetOpenFiles() {
				if strings.Contains(f.Path, "pts") {
					isShellJob = true
					break
				}
			}
		} else {
			log.Warn().Msg("No process info found. it should have been filled by an adapter")
		}

		if req.GetCriu() == nil {
			req.Criu = &criu_proto.CriuOpts{}
		}

		// Only set unless already set
		if req.GetCriu().ShellJob == nil {
			req.Criu.ShellJob = proto.Bool(isShellJob)
		}

		return next(ctx, server, nfy, resp, req)
	}
}

// Open files from the dump state are inherited by the restored process.
// For e.g. the standard streams (stdin, stdout, stderr) are inherited to use
// a log file.
func InheritOpenFilesForRestore(next types.Restore) types.Restore {
	return func(ctx context.Context, server types.ServerOpts, nfy *criu.NotifyCallbackMulti, resp *daemon.RestoreResp, req *daemon.RestoreReq) (chan int, error) {
		extraFiles, _ := ctx.Value(keys.RESTORE_EXTRA_FILES_CONTEXT_KEY).([]*os.File)
		ioFiles, _ := ctx.Value(keys.RESTORE_IO_FILES_CONTEXT_KEY).([]*os.File)
		inheritFds := req.GetCriu().GetInheritFd()

		if info := resp.GetState().GetInfo(); info != nil {
			// In case of attach, we need to create pipes for stdin, stdout, stderr
			if req.Attach {
				inReader, inWriter, err := os.Pipe()
				outReader, outWriter, err := os.Pipe()
				errReader, errWriter, err := os.Pipe()
				if err != nil {
					return nil, status.Errorf(
						codes.Internal,
						"failed to create pipes for attach: %v",
						err,
					)
				}
				ioFiles = append(ioFiles, inWriter, outReader, errReader)
				for _, f := range info.GetOpenFiles() {
					f.Path = strings.TrimPrefix(f.Path, "/")
					if f.Fd == 0 {
						extraFiles = append(extraFiles, inReader)
						inheritFds = append(inheritFds, &criu_proto.InheritFd{
							Fd:  proto.Int32(2 + int32(len(extraFiles))),
							Key: proto.String(f.Path),
						})
						defer inReader.Close()
					} else if f.Fd == 1 {
						extraFiles = append(extraFiles, outWriter)
						inheritFds = append(inheritFds, &criu_proto.InheritFd{
							Fd:  proto.Int32(2 + int32(len(extraFiles))),
							Key: proto.String(f.Path),
						})
						defer outWriter.Close()
					} else if f.Fd == 2 {
						extraFiles = append(extraFiles, errWriter)
						inheritFds = append(inheritFds, &criu_proto.InheritFd{
							Fd:  proto.Int32(2 + int32(len(extraFiles))),
							Key: proto.String(f.Path),
						})
						defer errWriter.Close()
					}
				}

				// In case of log, we need to open a log file
			} else {
				if req.Log == "" {
					req.Log = fmt.Sprintf(DEFAULT_RESTORE_LOG_PATH_FORMATTER, time.Now().Unix())
				}
				restoreLog, err := os.OpenFile(req.Log, RESTORE_LOG_FLAGS, RESTORE_LOG_PERMS)
				defer restoreLog.Close()
				if err != nil {
					return nil, status.Errorf(codes.Internal, "failed to open restore log: %v", err)
				}
				for _, f := range info.GetOpenFiles() {
					if f.Fd == 1 || f.Fd == 2 {
						f.Path = strings.TrimPrefix(f.Path, "/")
						extraFiles = append(extraFiles, restoreLog)
						inheritFds = append(inheritFds, &criu_proto.InheritFd{
							Fd:  proto.Int32(2 + int32(len(extraFiles))),
							Key: proto.String(f.Path),
						})
					} else {
						log.Warn().Msgf("found non-stdio open file %s with fd %d", f.Path, f.Fd)
					}
				}
			}
		} else {
			log.Warn().Msg("No process info found. it should have been filled by an adapter")
		}

		ctx = context.WithValue(ctx, keys.RESTORE_EXTRA_FILES_CONTEXT_KEY, extraFiles)
		ctx = context.WithValue(ctx, keys.RESTORE_IO_FILES_CONTEXT_KEY, ioFiles)

		// Set the inherited fds
		if req.GetCriu() == nil {
			req.Criu = &criu_proto.CriuOpts{}
		}
		req.GetCriu().InheritFd = inheritFds

		return next(ctx, server, nfy, resp, req)
	}
}
