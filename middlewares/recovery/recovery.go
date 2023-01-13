package recovery

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core"
)

func New(opts ...Option) core.HandlerFunc {
	cfg := newOptions(opts...)
	return func(ctx context.Context, pk *core.PianoKey) {
		defer func() {
			if err := recover(); err != nil {
				s := stack(3)
				cfg.recoveryHandler(ctx, pk, err, s)
			}
		}()
	}
}

// TODO: enrich stack
func stack(skip int) string {
	return ""
}
