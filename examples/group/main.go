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

	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
)

func main() {
	p := bin.Default()
	auth := p.Group("/auth")
	auth.GET("/ping", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "pong")
	})
	auth.GET("/binary", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "lorain")
	})
	p.Play()
}
