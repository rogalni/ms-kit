package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/rogalni/ms-kit/pkg/kit"
	"github.com/rogalni/ms-kit/pkg/kit/config"
	"github.com/rogalni/ms-kit/pkg/kit/fiber/middleware"
	"github.com/rogalni/ms-kit/pkg/kit/log"
)

func main() {
	kit := kit.Setup(config.Development())

	setupRoutes(kit.Server.App)
	kit.Server.Run()
}

func setupRoutes(app *fiber.App) {
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Ms-Kit Metrics"}))
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("", Index)
	v1.Get("/secure", middleware.Authorized(), Index)
}

func Index(c *fiber.Ctx) error {
	time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
	log.Ctx(c.UserContext()).Info("Hi")
	return c.SendString("Welcome!")
}
