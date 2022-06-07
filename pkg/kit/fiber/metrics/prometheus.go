package metrics

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/ms-kit/pkg/kit/config"
)

func Instrument(app *fiber.App) {
	prometheus := fiberprometheus.New(config.Kit.ServiceName)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
}
