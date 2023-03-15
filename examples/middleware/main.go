package main

import (
	"context"

	"github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	p.Use(func(ctx context.Context, pk *server.PianoKey) {
		core.Info("pre-handle")
		pk.Next(ctx)
	}, func(ctx context.Context, pk *server.PianoKey) {
		pk.Next(ctx)
		core.Info("post-handle")
	})
	p.GET("/mw", func(ctx context.Context, pk *server.PianoKey) {
		core.Info("in-handle")
		pk.String(consts.StatusOK, "middleware")
	})
	p.Play()
}
