package kit

import (
	"github.com/rogalni/ms-kit/pkg/kit/fiber"
	"github.com/rogalni/ms-kit/pkg/kit/log"
	"github.com/rogalni/ms-kit/pkg/kit/tracer"
)

type Kit struct {
	Server *fiber.Server
}

func Setup() *Kit {
	log.Setup()
	tracer.Instrument()
	srv := fiber.New()
	return &Kit{
		Server: srv,
	}
}
