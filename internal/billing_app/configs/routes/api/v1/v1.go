package v1

import (
	"backend/internal/billing_app/configs/routes/api/v1/auth"
	"backend/internal/billing_app/configs/routes/api/v1/pub"
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type ApiV1Routes struct {
	db            *database.StorageDb
	app           *fiber.App
	apiAuthRoutes *auth.RoutesApiAuth
	apiBillRoutes *auth.BillRoutes
	apiNewsRoutes *pub.ApiNewsRoutes
}

func NewApiV1Routes(db *database.StorageDb, app *fiber.App) *ApiV1Routes {
	return &ApiV1Routes{
		db:            db,
		app:           app,
		apiAuthRoutes: auth.NewRoutesApiAuth(db),
		apiBillRoutes: auth.NewBillRoutes(db),
		apiNewsRoutes: pub.NewApiNewsRoutes(db),
		//apiPaymentRoutes: auth.NewApiPaymentRoutes(db),
	}
}

func (avr ApiV1Routes) ApiV1Routes(group fiber.Router) *fiber.App {
	apiV1Group := group.Group("/v1")

	//apiV1Group.Get("/news/test/translate", news_listing.CreateListingNews)

	avr.apiAuthRoutes.ApiAuthRoutes(avr.app, apiV1Group)
	//
	avr.apiBillRoutes.BillingRoutes(avr.app, apiV1Group)

	//Public
	//{
	//apiV1Group.Get("testContext", func(c *fiber.Ctx) error {
	//	cacheManager := CacheManagerSync.GetCacheByLocals(c)
	//	data, ok := cacheManager.Get("news_listing_list_all")
	//	if ok {
	//		return c.JSON(fiber.Map{
	//			"news_by_content": data.(*news_listing.NewsListing),
	//		})
	//	} else {
	//		return c.JSON(fiber.Map{
	//			"news_by_content": []news_listing.NewsListing{},
	//		})
	//	}
	//
	//})
	//	//News
	avr.apiNewsRoutes.ApiNewsRoutes(avr.app, apiV1Group)
	//}

	return avr.app
}
