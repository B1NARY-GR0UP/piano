package main

import (
	"context"
	"net/http"

	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
)

func main() {
	p := bin.Default()
	auth := p.GROUP("/auth")
	auth.GET("/ping", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "pong")
	})
	auth.GET("/binary", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(http.StatusOK, "lorain")
	})
	p.Play()
}
