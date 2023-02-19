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
