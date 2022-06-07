package fiber

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/ms-kit/pkg/kit/config"
	"github.com/rogalni/ms-kit/pkg/kit/fiber/health"
	"github.com/rogalni/ms-kit/pkg/kit/fiber/shutdown"
	"github.com/rogalni/ms-kit/pkg/kit/log"
)

type Server struct {
	App           *fiber.App
	HealthHandler *health.HealthHandler
}

// Has to be always the last call in main since it is blocking.
func New(conf ...fiber.Config) *Server {
	c := fiber.Config{
		EnablePrintRoutes:     false,
		DisableStartupMessage: true,
	}
	if len(conf) > 0 {
		c = conf[0]
	}
	app := fiber.New(c)
	app.Use(otelfiber.Middleware(config.Kit.ServiceName))
	hh := health.For(app)
	return &Server{
		App:           app,
		HealthHandler: hh,
	}
}

func (s *Server) Run() {
	ctx := context.Background()
	gsc := shutdown.Gracefully(s.App, 2*time.Second)
	port := config.Kit.Port
	log.Ctx(ctx).Info(fmt.Sprintf("Running server on port %s", port))
	if err := s.App.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Ctx(ctx).Error(fmt.Sprintf("%e", err))
	}
	<-gsc
	log.Ctx(ctx).Info("Server stopped")
}
