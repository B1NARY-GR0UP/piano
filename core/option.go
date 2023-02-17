package core

const (
	defaultAddr = ":7246"
)

type Option func(o *Options)

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
		opt(o)
	}
}

// WithHostAddr used to define addr you prefer
func WithHostAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}
