package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/ms-kit/pkg/kit"
	"github.com/rogalni/ms-kit/pkg/kit/config"
	"github.com/rogalni/ms-kit/pkg/kit/fiber/middleware"
	"github.com/rogalni/ms-kit/pkg/kit/log"
)

func main() {
	kit := kit.Setup(config.Default())

	setupRoutes(kit.Server.App)
	kit.Server.Run()
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Use(middleware.Authorized())
	v1 := api.Group("/v1")
	v1.Get("", Index)
}

func Index(c *fiber.Ctx) error {
	log.Ctx(c.UserContext()).Info("Hi")
	return c.SendString("Welcome!")
}
