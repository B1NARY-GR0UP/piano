package main

import (
	"context"
	"net/http"

	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
)

func main() {
	p := bin.Default()
	// static route or common route
	p.GET("/ping", func(ctx context.Context, pk *core.PianoKey) {
		pk.JSON(http.StatusOK, core.M{
			"ping": "pong",
		})
	})
	// param route
	p.GET("/param/:username", func(ctx context.Context, pk *core.PianoKey) {
		pk.JSON(http.StatusOK, core.M{
			"username": pk.Param("username"),
		})
	})
	// wildcard route
	p.GET("/wild/*", func(ctx context.Context, pk *core.PianoKey) {
		pk.JSON(http.StatusOK, core.M{
			"route": "wildcard route",
		})
	})
	p.Play()
}
