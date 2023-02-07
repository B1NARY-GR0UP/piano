package recovery

import (
	"context"
	log "github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core"
)

type (
	RecoveryHandler func(ctx context.Context, pk *core.PianoKey, err any, info string)

	Option func(o *Options)

	Options struct {
		recoveryHandler RecoveryHandler
	}
)

func defaultRecoveryHandler(_ context.Context, pk *core.PianoKey, err any, info string) {
	log.Errorf("[PIANO] RECOVERY Err: %v Stack: %v", err, info)
	//pk.BreakWithStatus(http.StatusInternalServerError)
}

func newOptions(opts ...Option) *Options {
	options := &Options{
		recoveryHandler: defaultRecoveryHandler,
	}
	options.apply(opts...)
	return options
}

func (o *Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithRecoveryHandler(handler RecoveryHandler) Option {
	return func(o *Options) {
		o.recoveryHandler = handler
	}
}
