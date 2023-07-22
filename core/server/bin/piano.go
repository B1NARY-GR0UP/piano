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

package bin

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/middlewares/recovery"
	"github.com/B1NARY-GR0UP/piano/pkg/errors"
)

// Piano will respond to you.
type Piano struct {
	*server.Engine
}

// New a pure PIANO
func New(opts ...server.Option) *Piano {
	options := server.NewOptions(opts...)
	p := &Piano{
		Engine: server.NewEngine(options),
	}
	if !options.HideBanner {
		fmt.Printf(banner, Version)
	}
	return p
}

// Default will new a PIANO with recovery middleware
func Default(opts ...server.Option) *Piano {
	p := New(opts...)
	p.Use(recovery.New())
	return p
}

// Play the PIANO now
func (p *Piano) Play() {
	errCh := make(chan error)
	go func() {
		errCh <- p.Run()
	}()
	waitSignal := func(errCh chan error) error {
		signalToNotify := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}
		if signal.Ignored(syscall.SIGHUP) {
			signalToNotify = signalToNotify[1:]
		}
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, signalToNotify...)
		select {
		case sig := <-signalCh:
			switch sig {
			case syscall.SIGTERM:
				// force exit
				return errors.NewPublic(sig.String())
			case syscall.SIGHUP, syscall.SIGINT:
				// graceful shutdown
				core.Infof("---PIANO--- Receive signal: %v", sig)
				return nil
			}
		case err := <-errCh:
			return err
		}
		return nil
	}
	if err := waitSignal(errCh); err != nil {
		core.Errorf("---PIANO--- Receive close signal error: %v", err)
		return
	}
	core.Infof("---PIANO--- Begin graceful shutdown, wait up to %d seconds", p.Options().ShutdownTimeout/time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), p.Options().ShutdownTimeout)
	defer cancel()
	if err := p.Shutdown(ctx); err != nil {
		core.Errorf("---PIANO--- Shutdown err: %v", err)
	}
}
