package billing

import (
	"backend/internal/billing_app/response"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {
	dataResp := fiber.Map{
		"Title": "Личный кабинет",
	}
	//
	return response.ResponseTemp(c,
		"billing/dashboard",
		"billing/layout/dashboard",
		dataResp,
	)
}

func Cart(c *fiber.Ctx) error {
	return c.Render("billing/dashboard", fiber.Map{
		"Title": "Конфигурирование услуги — " + c.Params("name"),
	}, "billing/layout/dashboard")
}
