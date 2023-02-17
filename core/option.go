// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
