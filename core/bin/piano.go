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
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core"
	"github.com/B1NARY-GR0UP/piano/middleware/recovery"
)

// Piano will respond to you.
type Piano struct {
	*core.Engine
}

// New a pure PIANO
func New(opts ...core.Option) *Piano {
	options := core.NewOptions(opts...)
	p := &Piano{
		Engine: core.NewEngine(options),
	}
	return p
}

// Default will new a PIANO with recovery middleware
func Default(opts ...core.Option) *Piano {
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
				return errors.New(sig.String())
			case syscall.SIGHUP, syscall.SIGINT:
				// graceful shutdown
				log.Infof("[PIANO] Receive signal: %v", sig)
				return nil
			}
		case err := <-errCh:
			return err
		}
		return nil
	}
	if err := waitSignal(errCh); err != nil {
		log.Errorf("[PIANO] Receive close signal error: %v", err)
		return
	}
	log.Infof("[PIANO] Begin graceful shutdown, wait up to %d seconds", p.Options().ShutdownTimeout/time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), p.Options().ShutdownTimeout)
	defer cancel()
	if err := p.Shutdown(ctx); err != nil {
		log.Errorf("[PIANO] Shutdown err: %v", err)
	}
}
