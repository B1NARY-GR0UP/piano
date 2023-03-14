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
	"time"

	"github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default(server.WithShutdownTimeout(time.Second * 3))
	p.OnRun = append(p.OnRun, func(ctx context.Context) error {
		core.Info("hello")
		return nil
	})
	p.OnRun = append(p.OnRun, func(ctx context.Context) error {
		core.Info("world")
		return nil
	})
	p.OnShutdown = append(p.OnShutdown, func(ctx context.Context) {
		core.Info("binary")
	})
	p.OnShutdown = append(p.OnShutdown, func(ctx context.Context) {
		core.Info("piano")
	})
	p.GET("/ping", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "pong")
	})
	p.Play()
}
