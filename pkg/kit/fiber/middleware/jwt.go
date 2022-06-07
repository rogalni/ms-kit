package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/rogalni/ms-kit/pkg/kit/config"
	"github.com/rogalni/ms-kit/pkg/kit/log"
)

var jwt fiber.Handler

func Authorized() fiber.Handler {
	if jwt == nil {
		setupJwt()
	}
	return jwt
}

func setupJwt() {
	ksu := config.Kit.JwkSetUri
	if len(ksu) > 0 {
		jwt = jwtware.New(jwtware.Config{
			KeySetURL: ksu,
		})
	} else {
		log.Ctx(context.Background()).Warn("No JWK-URI configured!")
		jwt = func(c *fiber.Ctx) error {
			c.Next()
			return nil
		}
	}
}
