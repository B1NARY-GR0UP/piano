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

package main

import (
	"context"
	"net/http"
	"time"

	log "github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
)

func main() {
	p := bin.Default(core.WithShutdownTimeout(time.Second * 3))
	p.OnRun = append(p.OnRun, func(ctx context.Context) error {
		log.Info("hello")
		return nil
	})
	p.OnRun = append(p.OnRun, func(ctx context.Context) error {
		log.Info("world")
		return nil
	})
	p.OnShutdown = append(p.OnShutdown, func(ctx context.Context) {
		log.Info("binary")
	})
	p.OnShutdown = append(p.OnShutdown, func(ctx context.Context) {
		log.Info("piano")
	})
	p.GET("/ping", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "pong")
	})
	p.Play()
}
