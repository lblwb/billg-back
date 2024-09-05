package news

import (
	"backend/internal/billing_app/models/news/news_listing"
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type NwsBannerService struct {
	db          *database.StorageDb
	newsListing *news_listing.NewsListingEntity
}

func NewNewsBannerService(db *database.StorageDb) *NwsBannerService {
	return &NwsBannerService{db: db}
}

/**
Сервис для работы с баннерными новостями
*/

// GetBannerNewsListAll - Получение списка листинга новостей
func (nbs NwsBannerService) GetBannerNewsListAll(c *fiber.Ctx) error {
	allListingNews := nbs.newsListing.GetAllListingNews()
	if allListingNews != nil {
		return c.JSON(fiber.Map{
			"banner_nws": allListingNews,
		})
	} else {
		return c.JSON(fiber.Map{
			"banner_nws": []news_listing.NewsListing{},
		})
	}
}
