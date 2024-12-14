package types

import (
	"context"
	"sync"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"github.com/cedana/cedana/pkg/criu"
	"github.com/cedana/cedana/pkg/plugins"
)

type (
	// ServerOpts is intended to be passed by **value** to each handler, so that each handler can modify it
	// before passing it to the next handler in the chain, without affecting the original value.
	ServerOpts struct {
		WG           *sync.WaitGroup
		CRIU         *criu.Criu
		CRIUCallback *criu.NotifyCallbackMulti
		Plugins      plugins.Manager
		Lifetime     context.Context
	}

	Dump    func(context.Context, ServerOpts, *daemon.DumpResp, *daemon.DumpReq) (exited chan int, err error)
	Restore func(context.Context, ServerOpts, *daemon.RestoreResp, *daemon.RestoreReq) (exited chan int, err error)
	Run     func(context.Context, ServerOpts, *daemon.RunResp, *daemon.RunReq) (exited chan int, err error)
	Manage  func(context.Context, ServerOpts, *daemon.ManageResp, *daemon.ManageReq) (exited chan int, err error)

	// An adapter is a function that takes a Handler and returns a new Handler
	Adapter[H Dump | Restore | Run | Manage] func(H) H

	// A middleware is simply a chain of adapters
	Middleware[H Dump | Restore | Run | Manage] []Adapter[H]
)

// With is a method on Handler that applies a list of Middleware to the Handler
func (h Dump) With(middleware ...Adapter[Dump]) Dump {
	return adapted(h, middleware...)
}

// With is a method on Handler that applies a list of Middleware to the Handler
func (h Restore) With(middleware ...Adapter[Restore]) Restore {
	return adapted(h, middleware...)
}

// With is a method on Handler that applies a list of Middleware to the Handler
func (h Run) With(middleware ...Adapter[Run]) Run {
	return adapted(h, middleware...)
}

// With is a method on Handler that applies a list of Middleware to the Handler
func (h Manage) With(middleware ...Adapter[Manage]) Manage {
	return adapted(h, middleware...)
}

//////////////////////////
//// Helper Functions ////
//////////////////////////

// Adapted takes a Handler and a list of Adapters, and
// returns a new Handler that applies the adapters in order.
func adapted[H Dump | Restore | Run | Manage](h H, adapters ...Adapter[H]) H {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}
	return h
}