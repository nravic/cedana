package profiling

import (
	"context"
	"reflect"
	"runtime"
	"time"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"github.com/cedana/cedana/internal/metrics"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
)

// RecordDuration records the elapsed time since start into the profiling data.
// Use with defer to record the time spent in a function.
// If no f is provided, uses the caller.
func RecordDuration(ctx context.Context, profiling *daemon.ProfilingData, f ...any) (childCtx context.Context, end func()) {
	start := time.Now()
	childCtx, span := otel.Tracer(metrics.API_TRACER).Start(ctx, "RecordDuration")

	return childCtx, func() {
		duration := time.Since(start)
		span.End()
		profiling.Duration = duration.Nanoseconds()

		if profiling.Name == "" {
			var pc uintptr
			if len(f) == 0 {
				pc, _, _, _ = runtime.Caller(1)
			} else {
				pc = reflect.ValueOf(f[0]).Pointer()
			}

			profiling.Name = utils.FunctionName(pc)
		}

		log.Trace().Str("in", profiling.Name).Msgf("spent %s", duration)
	}
}

// RecordComponentDuration records the elapsed time since start into the profiling data.
// Unlike RecordDuration, this adds the data as a new component of the profiling data.
// Also returns the component that was added.
func RecordDurationComponent(start time.Time, profiling *daemon.ProfilingData, f ...any) *daemon.ProfilingData {
	duration := time.Since(start)

	var pc uintptr
	if len(f) == 0 {
		pc, _, _, _ = runtime.Caller(1)
	} else {
		pc = reflect.ValueOf(f[0]).Pointer()
	}
	name := utils.FunctionName(pc)

	component := &daemon.ProfilingData{
		Name:     name,
		Duration: duration.Nanoseconds(),
	}

	profiling.Duration += duration.Nanoseconds()
	profiling.Components = append(profiling.Components, component)

	log.Trace().Str("in", name).Msgf("spent %s", duration)
	return component
}

// RecordDurationCategory records the elapsed time since start into the profiling data.
// Instead of directly inserting a component like RecordDurationComponent, this adds the data as a nested component,
// with the name matching the category provided. Also returns the component that was added to the category.
func RecordDurationCategory(start time.Time, profiling *daemon.ProfilingData, category string, f ...any) *daemon.ProfilingData {
	var categoryComponent *daemon.ProfilingData
	for _, component := range profiling.Components {
		if component.Name == category {
			categoryComponent = component
			break
		}
	}
	if categoryComponent == nil {
		categoryComponent = &daemon.ProfilingData{
			Name: category,
		}
		profiling.Components = append(profiling.Components, categoryComponent)
	}

	return RecordDurationComponent(start, categoryComponent, f...)
}

// LogDuration logs the elapsed time since start.
// Use with defer to log the time spent in a function
// If no f is provided, uses the caller.
func LogDuration(start time.Time, f ...any) {
	duration := time.Since(start)

	var pc uintptr
	if len(f) == 0 {
		pc, _, _, _ = runtime.Caller(1)
	} else {
		pc = reflect.ValueOf(f[0]).Pointer()
	}

	name := utils.FunctionName(pc)
	log.Trace().Str("in", name).Msgf("spent %s", duration)
}
