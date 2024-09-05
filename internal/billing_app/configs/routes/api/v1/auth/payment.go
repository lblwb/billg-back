package auth

import (
	"backend/internal/database"
)

type ApiPaymentRoutes struct {
	//app *fiber.App
	db *database.StorageDb
}

//func NewApiPaymentRoutes(db *database.StorageDb) *ApiPaymentRoutes {
//	return &ApiPaymentRoutes{
//		//app: app,
//		db: db,
//	}
//}

//func (upr ApiPaymentRoutes) PaymentRoutes(app *fiber.App) *fiber.App {
//
//	//upr.app.Get("/", func(c *fiber.Ctx) error {
//	//
//	//	// Render index
//	//	return c.Render("index", fiber.Map{
//	//		"Title": "Биллинг",
//	//	}, "layout/main")
//	//
//	//})
//
//	upr.app.Post("p")
//
//	return app
//}

//
//func PrivateRoute() {
//
//}
