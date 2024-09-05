package middleware

import (
	"backend/pkg/render/inertia-fiber"
	"github.com/gofiber/fiber/v2"
)

func NewRenderLocMidl(engine *inertia.Engine) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("renderEngine", engine)
		return c.Next()
	}
}
