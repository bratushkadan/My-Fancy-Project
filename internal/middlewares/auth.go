package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// https://github.com/gofiber/jwt
// https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
func Auth() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Get()

		return nil
	}
}
