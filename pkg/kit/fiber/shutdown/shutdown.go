package shutdown

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/ms-kit/pkg/kit/log"
)

func Gracefully(app *fiber.App, timeout time.Duration) chan struct{} {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	serverShutdown := make(chan struct{}, 1)
	go func() {
		<-sig
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		log.Ctx(ctx).Debug("Gracefully shutting down...")
		defer cancel()
		go func() {
			err := app.Shutdown()
			if err != nil {
				log.Ctx(ctx).Error(fmt.Sprintf("Error shutdown server: %v\n", err))
			}
			cancel()
		}()
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				log.Ctx(ctx).Warn("Gracefull shutown timed out! Force shutdown")
			case context.Canceled:
				log.Ctx(ctx).Debug("Gracefull shutdown sucessfull completed")
			}
		}
		serverShutdown <- struct{}{}
	}()
	return serverShutdown
}
