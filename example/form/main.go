package main

import (
	"context"
	"net/http"

	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
)

func main() {
	p := bin.Default()
	p.POST("/form", func(ctx context.Context, pk *core.PianoKey) {
		pk.JSON(http.StatusOK, core.M{
			"username": pk.PostForm("username"),
			"password": pk.PostForm("password"),
		})
	})
	p.Play()
}
