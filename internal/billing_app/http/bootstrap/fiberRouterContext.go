package bootstrap

import (
	"backend/internal/billing_app/configs/routes/api"
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

func mainRoutersMap(app *fiber.App, db *database.StorageDb) *fiber.App {
	//Initialize new Routers Context

	// Main Api Routes
	api.MainApiRoutes(app, db)
	return app
}

func NewFiberRouterContext(app *fiber.App, db *database.StorageDb) *fiber.App {
	//
	mainRoutersMap(app, db)
	//
	app.Static("/", "./public")

	//Notfound
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusFound).SendString("NotFound")
	})

	return app
}
