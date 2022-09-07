package rpctracing

import (
	"context"
	"github.com/spacemonkeygo/monkit/v3"
)

const (
	annotation = "tracing"
	value      = "local"
)

// WithoutDistributedTracing disables distributed tracing for the current span.
func WithoutDistributedTracing(ctx context.Context) {
	monkit.SpanFromCtx(ctx).Annotate(annotation, value)
}

// IsExcluded check if span shouldn't be reported to remote location.
func IsExcluded(span *monkit.Span) bool {
	for _, a := range span.Annotations() {
		if a.Name == annotation && a.Value == value {
			return true
		}
	}
	return false
}
