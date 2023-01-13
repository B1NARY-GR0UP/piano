package core

const (
	defaultAddr = ":7246"
)

type Option struct {
	F func(o *Options)
}

type Options struct {
	Addr string
}

// NewOptions for PIANO engine
func NewOptions(opts ...Option) *Options {
	options := &Options{
		Addr: defaultAddr,
	}
	options.apply(opts...)
	return options
}

func (o *Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt.F(o)
	}
}

// WithHostAddr used to define addr you prefer
func WithHostAddr(addr string) Option {
	return Option{F: func(o *Options) {
		o.Addr = addr
	}}
}
