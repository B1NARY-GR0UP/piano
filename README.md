# PIANO

[![Go Report Card](https://goreportcard.com/badge/github.com/B1NARY-GR0UP/piano)](https://goreportcard.com/report/github.com/B1NARY-GR0UP/piano)

> Piano will respond to you.

![piano](images/PIANO.png)

PIANO is a simple and lightweight HTTP framework. More features will be supported gradually.

## Install

```shell
go get github.com/B1NARY-GR0UP/piano
```

## Quick Start

### Hello

[example](examples/hello)

```go
package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	p.GET("/hello", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "piano")
	})
	p.Play()
}
```

### Route

[example](examples/route)

```go
package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	// static route or common route
	p.GET("/ping", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"ping": "pong",
		})
	})
	// param route
	p.GET("/param/:username", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"username": pk.Param("username"),
		})
	})
	// wildcard route
	p.GET("/wild/*", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"route": "wildcard route",
		})
	})
	p.Play()
}
```

### Group

[example](examples/group)

```go
package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	auth := p.Group("/auth")
	auth.GET("/ping", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "pong")
	})
	auth.GET("/binary", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "lorain")
	})
	p.Play()
}
```

### Middleware

[example](examples/middleware)

```go
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
```

### Query

[example](examples/query)

```go
package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	p.GET("/query", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"username": pk.Query("username"),
		})
	})
	p.Play()
}
```

### Form

[example](examples/form)

```go
package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	p.POST("/form", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"username": pk.PostForm("username"),
			"password": pk.PostForm("password"),
		})
	})
	p.Play()
}
```

### Hook

[example](examples/hook)

```go
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
```

You can also go through the code for more information.

## Related Projects

- [DREAMEMO](https://github.com/B1NARY-GR0UP/dreamemo) | A distributed cache with out-of-the-box, high-scalability, modular-design features. | `golang` `cache` `distributed`
- [INQUISITOR](https://github.com/B1NARY-GR0UP/inquisitor) | A simple and lightweight log. | `golang` `log`

## Blogs

- [PIANO: A Simple and Lightweight HTTP Framework Implemented in Go](https://dev.to/justlorain/piano-a-simple-and-lightweight-http-framework-implemented-in-go-224p)
- [如何使用 channel 实现一个优雅退出功能？](https://juejin.cn/post/7207423263344427068)

## License

PIANO is distributed under the [Apache License 2.0](./LICENSE). The licenses of third party dependencies of PIANO are explained [here](./licenses).

## ECOLOGY

<p align="center">
<img src="https://github.com/justlorain/justlorain/blob/main/images/BINARY-WEB-ECO.png" width="300"/>
<br/>
PIANO is a subproject of the <a href="https://github.com/B1NARY-GR0UP">BINARY WEB ECOLOGY</a>.
</p>