package middlewares

import (
	"github.com/eduardoths/finance-api/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func WithContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		conn := db.GetFromContext(ctx)
		ctx = db.SetToContext(conn, ctx)

		c.SetUserContext(ctx)
		return c.Next()
	}
}
