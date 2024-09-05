package v2

import "github.com/gofiber/fiber/v2"

func ApiV2Routes(app *fiber.App, group fiber.Router) *fiber.App {
	//c.Set("Version", "v1")
	apiV2Group := group.Group("/v2", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v2")
		return c.Next()
	})

	apiV2Group.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello world!")
	})

	return app
}
