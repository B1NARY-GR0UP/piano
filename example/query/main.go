package main

import (
	"context"
	"net/http"

	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
)

func main() {
	p := bin.Default()
	p.GET("/query", func(ctx context.Context, pk *core.PianoKey) {
		pk.JSON(http.StatusOK, core.M{
			"username": pk.Query("username"),
		})
	})
	p.Play()
}
