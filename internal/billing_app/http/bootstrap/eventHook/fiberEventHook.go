package eventHook

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func NewBootFbEventHkService(app *fiber.App) *fiber.App {
	app.Hooks().OnShutdown(func() error {
		fmt.Println("Завершение работы!")
		return nil
	})

	return app
}
