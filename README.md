![piano](images/PIANO.png)

PIANO is a simple and lightweight HTTP framework.

[![Go Report Card](https://goreportcard.com/badge/github.com/B1NARY-GR0UP/piano)](https://goreportcard.com/report/github.com/B1NARY-GR0UP/piano)

## Install

```shell
go get github.com/B1NARY-GR0UP/piano
```

## Quick Start

```go
package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/core/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	p.GET("/hello", func(ctx context.Context, pk *core.PianoKey) {
		pk.String(consts.StatusOK, "piano")
	})
	p.Play()
}
```

Refer to [piano-examples](https://github.com/rainiring/piano-examples) for more information.

## Related Projects

- [PIANO-EXAMPLES](https://github.com/rainiring/piano-examples) | Examples for PIANO | `examples`
- [DREAMEMO](https://github.com/B1NARY-GR0UP/dreamemo) | A distributed cache with out-of-the-box, high-scalability, modular-design features. | `golang` `cache` `distributed`

## Blogs

- [PIANO: A Simple and Lightweight HTTP Framework Implemented in Go](https://dev.to/justlorain/piano-a-simple-and-lightweight-http-framework-implemented-in-go-224p)
- [如何使用 channel 实现一个优雅退出功能？](https://juejin.cn/post/7207423263344427068)

## Acknowledgement

Thanks for the IDE license support by [JetBrains OpenSource Support](https://jb.gg/OpenSourceSupport).

<img src="./images/jb_beam.png" width="300" alt="jb beam"/>

## License

PIANO is distributed under the [Apache License 2.0](./LICENSE). The licenses of third party dependencies of PIANO are explained [here](./licenses).

## ECOLOGY

<p align="center">
<img src="https://github.com/justlorain/justlorain/blob/main/images/BINARY-WEB-ECO.png" alt="BINARY-WEB-ECO"/>
<br/><br/>
PIANO is a Subproject of the <a href="https://github.com/B1NARY-GR0UP">BINARY WEB ECOLOGY</a>
</p>