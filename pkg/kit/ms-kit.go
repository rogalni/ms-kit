package kit

import (
	"github.com/rogalni/ms-kit/internal/config"
	"github.com/rogalni/ms-kit/pkg/kit/fiber"
	"github.com/rogalni/ms-kit/pkg/kit/log"
	"github.com/rogalni/ms-kit/pkg/kit/tracer"
)

type Kit struct {
	Server *fiber.Server
}

func Setup() *Kit {
	log.Setup()
	if config.BEnv(config.EnvTracingEnabled) {
		tracer.Instrument()
	}

	srv := fiber.New()
	return &Kit{
		Server: srv,
	}
}
