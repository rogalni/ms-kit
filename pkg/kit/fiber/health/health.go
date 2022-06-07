package health

import "github.com/gofiber/fiber/v2"

type HealthHandler struct {
	r fiber.Router
}

type StatusFunc func() Health

func For(app *fiber.App) *HealthHandler {
	g := app.Group("/health")

	g.Get("", health)

	hh := &HealthHandler{
		r: g,
	}
	return hh
}

func health(app *fiber.Ctx) error {
	app.JSON(&Health{
		Status: UP,
	})
	return nil
}

func (h *HealthHandler) WithReadiness(sf StatusFunc) *HealthHandler {
	h.r.Get("/readiness", func(app *fiber.Ctx) error {
		es := sf()
		for _, v := range es.Components {
			if v.Status == DOWN {
				es.Status = DOWN
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
			if v.Status == DOWN {
				es.Status = DOWN
				break
			}
		}
		app.JSON(es)
		return nil
	})
	return h
}
