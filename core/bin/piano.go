package bin

import (
	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/middlewares/recovery"
)

// Piano will respond to you.
type Piano struct {
	*core.Engine
}

// New a pure PIANO
func New(opts ...core.Option) *Piano {
	options := core.NewOptions(opts...)
	p := &Piano{
		Engine: core.NewEngine(options),
	}
	return p
}

// Default will new a PIANO with recovery middleware
func Default(opts ...core.Option) *Piano {
	p := New(opts...)
	p.Use(recovery.New())
	return p
}
