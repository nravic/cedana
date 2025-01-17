package criu

import (
	"context"
	"path/filepath"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"github.com/cedana/cedana/pkg/profiling"
	"github.com/cedana/cedana/pkg/types"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	CRIU_LOG_VERBOSITY_LEVEL = 3
	CRIU_LOG_FILE            = "criu.log"
	GHOST_FILE_MAX_SIZE      = 10000000 // 10MB
)

var Dump types.Dump = dump

// Returns a CRIU dump handler for the server
func dump(ctx context.Context, opts types.Opts, resp *daemon.DumpResp, req *daemon.DumpReq) (exited chan int, err error) {
	if req.GetCriu() == nil {
		return nil, status.Error(codes.InvalidArgument, "criu options is nil")
	}

	version, err := opts.CRIU.GetCriuVersion(ctx)
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
	criuOpts.LogToStderr = proto.Bool(false)

	// TODO: Add support for pre-dump
	// TODO: Add support for lazy migration

	log.Debug().Int("CRIU", version).Interface("opts", criuOpts).Msg("CRIU dump starting")
	// utils.LogProtoMessage(criuOpts, "CRIU option", zerolog.DebugLevel)

	ctx, end := profiling.StartTimingCategory(ctx, "criu", opts.CRIU.Dump)

	_, err = opts.CRIU.Dump(ctx, criuOpts, opts.CRIUCallback)

	end()

	// Capture internal logs from CRIU
	utils.LogFromFile(
		log.With().Int("CRIU", version).Logger().WithContext(ctx),
		filepath.Join(criuOpts.GetImagesDir(), CRIU_LOG_FILE),
		zerolog.TraceLevel,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed CRIU dump: %v", err)
	}

	log.Debug().Int("CRIU", version).Msg("CRIU dump complete")

	return utils.WaitForPid(resp.State.PID), nil
}
