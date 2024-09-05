package routes

import "github.com/gofiber/fiber/v2"

func PublicRoutes(app *fiber.App) *fiber.App {

	//Fallback Spa
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.Render("public/spa1/index", fiber.Map{
			"JsonData": "{'test':'1'}",
		}, "")
	})
	return app
}
