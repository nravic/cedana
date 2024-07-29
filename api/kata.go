package api

// Implements the task service functions for kata container workloads

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"os"

	"github.com/cedana/cedana/api/services/task"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// const (
// 	KATA_OUTPUT_FILE_PATH  string      = "/tmp/log/cedana-output.log"
// 	KATA_OUTPUT_FILE_PERMS os.FileMode = 0o777
// 	KATA_OUTPUT_FILE_FLAGS int         = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
// )

func (s *service) KataDump(ctx context.Context, args *task.DumpArgs) (*task.DumpResp, error) {
	var err error

	ctx, dumpTracer := s.tracer.Start(ctx, "dump-ckpt")
	dumpTracer.SetAttributes(attribute.String("jobID", args.JID))
	defer dumpTracer.End()

	state := &task.ProcessState{}
	kataAgentPid := childPidFromPPid(1)
	pid := childPidFromPPid(kataAgentPid)

	state, err = s.generateState(ctx, pid)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	state.JID = args.JID

	err = s.kataDump(ctx, state, args)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		dumpTracer.RecordError(st.Err())
		return nil, st.Err()
	}

	var resp task.DumpResp

	switch args.Type {
	case task.CRType_LOCAL:
		resp = task.DumpResp{
			Message:      fmt.Sprintf("Dumped process %d to %s", pid, args.Dir),
			CheckpointID: state.CheckpointPath, // XXX: Just return path for ID for now
		}
	}

	err = s.updateState(ctx, state.JID, state)
	if err != nil {
		st := status.New(codes.Internal, fmt.Sprintf("failed to update state with error: %s", err.Error()))
		return nil, st.Err()
	}

	resp.State = state

	return &resp, err
}

func (s *service) KataRestore(ctx context.Context, args *task.RestoreArgs) (*task.RestoreResp, error) {
	ctx, restoreTracer := s.tracer.Start(ctx, "restore-ckpt")
	restoreTracer.SetAttributes(attribute.String("jobID", args.JID))
	defer restoreTracer.End()

	var resp task.RestoreResp
	var pid *int32
	var err error

	if args.CheckpointPath == "" {
		return nil, status.Error(codes.InvalidArgument, "checkpoint path cannot be empty")
	}
	if stat, err := os.Stat(args.CheckpointPath); os.IsNotExist(err) || stat.IsDir() || !strings.HasSuffix(args.CheckpointPath, ".tar") {
		return nil, status.Error(codes.InvalidArgument, "invalid checkpoint path")
	}

	pid, err = s.kataRestore(ctx, args)
	if err != nil {
		staterr := status.Error(codes.Internal, fmt.Sprintf("failed to restore process: %v", err))
		restoreTracer.RecordError(staterr)
		return nil, staterr
	}

	resp = task.RestoreResp{
		Message: fmt.Sprintf("successfully restored process: %v", *pid),
		NewPID:  *pid,
	}

	state, err := s.generateState(ctx, *pid)
	if err != nil {
		s.logger.Warn().Err(err).Msg("failed to generate state after restore")
	}

	// Only update state if it was a managed job
	if args.JID != "" && state != nil {
		state.JobState = task.JobState_JOB_RUNNING
		err = s.updateState(ctx, state.JID, state)
		if err != nil {
			s.logger.Warn().Err(err).Msg("failed to update managed job state after restore")
		}
	}

	resp.State = state

	return &resp, nil
}

//////////////////////////
///// Kata Utils //////
//////////////////////////

func childPidFromPPid(ppid int32) (int32) {
	// Replace PID with the actual parent process ID

	// Run the pgrep command
	cmd := exec.Command("pgrep", "--parent", strconv.Itoa(int(ppid)))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return -1
	}

	// Get the first line of the output
	pgrepOutput := strings.TrimSpace(out.String())
	lines := strings.Split(pgrepOutput, "\n")
	if len(lines) == 0 {
		return -1
	}
	firstLine := lines[0]

	// Convert the first line to an integer (PID of the first child process)
	firstChildPID, err := strconv.Atoi(firstLine)
	if err != nil {
		return -1
	}

	return int32(firstChildPID)
}
