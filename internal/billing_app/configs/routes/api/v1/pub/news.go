package pub

import (
	"backend/internal/billing_app/service/api/news"
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type ApiNewsRoutes struct {
	db                 *database.StorageDb
	NewsListingService *news.NwsListingService
	NewsBannerService  *news.NwsBannerService
}

func NewApiNewsRoutes(db *database.StorageDb) *ApiNewsRoutes {
	return &ApiNewsRoutes{
		db:                 db,
		NewsListingService: news.NewNwsListingService(db),
		NewsBannerService:  news.NewNewsBannerService(db),
	}
}

// ApiNewsRoutes - News Group Routes
func (anr ApiNewsRoutes) ApiNewsRoutes(app *fiber.App, group fiber.Router) *fiber.App {
	newsGroup := group.Group("pub/news")

	//Listing news
	newsGroupListing := newsGroup.Group("listing")
	newsGroupListing.Get(":slug/show", anr.NewsListingService.GetListingNewsBySlug)
	newsGroupListing.Get("all", anr.NewsListingService.GetAlertNewsListAll)

	//Banners news
	newsGroupBanner := newsGroup.Group("banner")
	newsGroupBanner.Get("list", anr.NewsBannerService.GetBannerNewsListAll)

	return app
}
