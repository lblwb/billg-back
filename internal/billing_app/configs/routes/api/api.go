package api

import (
	v1 "backend/internal/billing_app/configs/routes/api/v1"
	v2 "backend/internal/billing_app/configs/routes/api/v2"
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

func MainApiRoutes(app *fiber.App, db *database.StorageDb) *fiber.App {
	//c.Set("Version", "v1")
	apiGroup := app.Group("api")
	apiV1Routes := v1.NewApiV1Routes(db, app)
	apiV1Routes.ApiV1Routes(apiGroup)

	//Initialize v1
	v2.ApiV2Routes(app, apiGroup)

	//Not Found Routes
	apiGroup.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(struct {
			Error   string `json:"error"`
			Success bool   `json:"success"`
		}{
			"Not Found",
			false,
		})
	})

	return app
}
