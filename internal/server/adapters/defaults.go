package adapters

import (
	"context"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"buf.build/gen/go/cedana/criu/protocolbuffers/go/criu"
	"github.com/cedana/cedana/internal/config"
	"github.com/cedana/cedana/pkg/types"
	"google.golang.org/protobuf/proto"
)

// This file contains all the adapters that fill in missing request details
// with defaults

///////////////////////
//// Dump Adapters ////
///////////////////////

// Adapter that fills missing info from the request using config defaults
func FillMissingDumpDefaults(next types.Dump) types.Dump {
	return func(ctx context.Context, server types.ServerOpts, resp *daemon.DumpResp, req *daemon.DumpReq) (chan int, error) {
		if req.GetDir() == "" {
			req.Dir = config.Get(config.STORAGE_DUMP_DIR)
		}

		if req.GetCriu() == nil {
			req.Criu = &criu.CriuOpts{}
		}

		// Only override if unset
		if req.GetCriu().LeaveRunning == nil {
			req.Criu.LeaveRunning = proto.Bool(config.Get(config.CRIU_LEAVE_RUNNING))
		}

		return next(ctx, server, resp, req)
	}
}

//////////////////////////
//// Restore Adapters ////
//////////////////////////

// Adapter that fills missing info from the request using config defaults
func FillMissingRestoreDefaults(next types.Restore) types.Restore {
	return func(ctx context.Context, server types.ServerOpts, resp *daemon.RestoreResp, req *daemon.RestoreReq) (chan int, error) {
		// Nothing to do, yet

		return next(ctx, server, resp, req)
	}
}

////////////////////////
//// Start Adapters ////
////////////////////////

// Adapter that fills missing info from the request using config defaults
func FillMissingStartDefaults(next types.Start) types.Start {
	return func(ctx context.Context, server types.ServerOpts, resp *daemon.StartResp, req *daemon.StartReq) (chan int, error) {
		// Nothing to fill in for now

		return next(ctx, server, resp, req)
	}
}
