package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/ms-kit/pkg/kit/health"
)

type HealthHandler struct {
	r fiber.Router
}

type StatusFunc func() health.Health

func For(app *fiber.App) *HealthHandler {
	g := app.Group("/health")
	g.Get("", getHealth)
	hh := &HealthHandler{
		r: g,
	}
	return hh
}

func getHealth(app *fiber.Ctx) error {
	app.JSON(&health.Health{
		Status: health.UP,
	})
	return nil
}

func (h *HealthHandler) WithReadiness(sf StatusFunc) *HealthHandler {
	h.r.Get("/readiness", func(app *fiber.Ctx) error {
		es := sf()
		for _, v := range es.Components {
			if v.Status == health.DOWN {
				es.Status = health.DOWN
				break
			}
		}
		app.JSON(es)
		return nil
	})
	return h
}

func (h *HealthHandler) WithLiveness(sf StatusFunc) *HealthHandler {
	h.r.Get("/liveness", func(app *fiber.Ctx) error {
		es := sf()
		for _, v := range es.Components {
			if v.Status == health.DOWN {
				es.Status = health.DOWN
				break
			}
		}
		app.JSON(es)
		return nil
	})
	return h
}
