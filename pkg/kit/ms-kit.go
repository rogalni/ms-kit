package kit

import (
	"github.com/rogalni/ms-kit/pkg/kit/config"
	"github.com/rogalni/ms-kit/pkg/kit/fiber"
	"github.com/rogalni/ms-kit/pkg/kit/log"
	"github.com/rogalni/ms-kit/pkg/kit/tracer"
)

type Kit struct {
	Server *fiber.Server
}

func Setup(cfg config.KitConfig) *Kit {
	config.Apply(cfg)
	log.Setup()
	if cfg.IsTracingEnabled {
		tracer.Instrument()
	}
	srv := fiber.New()
	return &Kit{
		Server: srv,
	}
}

func (k *Kit) RunServer() {
	k.Server.Run()
}
