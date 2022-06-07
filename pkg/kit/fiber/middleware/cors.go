package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rogalni/ms-kit/pkg/kit/config"
)

var c fiber.Handler

func Cors() fiber.Handler {
	if c == nil {
		setupCors()
	}
	return c
}

func setupCors() {
	envCors := config.Kit.CorsUrls
	if len(envCors) > 0 {
		c = cors.New(cors.Config{
			AllowOrigins: envCors,
		})
	} else {
		c = cors.New()
	}
}
