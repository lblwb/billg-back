package news

import (
	"backend/internal/billing_app/models/news/news_listing"
	"backend/internal/billing_app/models/service"
	"backend/internal/database"
	CacheManagerSync "backend/pkg/cache"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

/**
Сервис для работы с последними новостями
*/

type NwsListingService struct {
	newsListingEntity *news_listing.NewsListingEntity
}

func NewNwsListingService(db *database.StorageDb) *NwsListingService {
	return &NwsListingService{
		newsListingEntity: news_listing.NewNewsListEntity(db),
	}
}

func (nls NwsListingService) GetAlertNewsListAll(c *fiber.Ctx) error {
	cacheManager := CacheManagerSync.GetCacheByLocals(c)
	if cacheManager != nil {
		//Cache
		allServices := cacheManager.Remember("news_listing_list_all", 2*time.Minute, func() interface{} {
			allServices := nls.newsListingEntity.GetAllListingNews()
			return allServices
		})
		if allServices != nil {
			return c.JSON(fiber.Map{
				"list_news": allServices,
			})
		}
	}
	return c.JSON(fiber.Map{
		"list_news": []service.Services{},
	})
}

func (nls NwsListingService) GetListingNewsBySlug(c *fiber.Ctx) error {
	paramsSlug := c.Params("slug")
	cacheManager := CacheManagerSync.GetCacheByLocals(c)
	modelListingNews := cacheManager.Remember(fmt.Sprintf("news_by_slug_%s", paramsSlug), 2*time.Minute, func() interface{} {
		modelListingNews, err := nls.newsListingEntity.GetBySlugListingNews(paramsSlug)
		if err != nil {
			return modelListingNews
		}
		return modelListingNews
	})

	//fmt.Println(modelListingNews)
	if modelListingNews == nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {
		return c.JSON(fiber.Map{
			"news":    modelListingNews,
			"success": true,
		})
	}

}
