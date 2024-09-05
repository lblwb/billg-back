package service

import (
	"backend/internal/billing_app/models/service"
	"backend/internal/database"
	CacheManagerSync "backend/pkg/cache"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

/**
Сервис для работы
с Услугами
*/

type BillServices struct {
	db        *database.StorageDb
	svcEntity *service.ServicesEntity
}

func NewBillServices(db *database.StorageDb) *BillServices {
	return &BillServices{
		db:        db,
		svcEntity: service.NewServicesEntity(db),
	}
}

func (bs BillServices) GetEntireServices(c *fiber.Ctx) error {
	//claims
	//user := c.Locals("user").(*jwt.Token)
	//log.Println("user", "locals", "claims", user)

	//TODO: Добавить проверку авторизации и сделать вывод
	// инфы на основе наличия ранее заказов у клиента

	allServices := bs.svcEntity.GetAllServices()

	if allServices != nil {
		return c.JSON(fiber.Map{
			"services": allServices,
		})
	} else {
		return c.JSON(fiber.Map{
			"services": []service.Services{},
		})
	}
}

func (bs BillServices) GetServiceBySlug(c *fiber.Ctx) error {
	paramsSlug := c.Params("slug")

	cacheManager := CacheManagerSync.GetCacheByLocals(c)
	modelCacheService := cacheManager.Remember(fmt.Sprintf("service_by_slug_%s", paramsSlug), 15*time.Second, func() interface{} {
		modelService, err := bs.svcEntity.GetServiceBySlug(paramsSlug)
		if err != nil {
			return modelService
		}
		return modelService
	})

	//fmt.Println(modelCacheService)
	if modelCacheService == nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {
		return c.JSON(fiber.Map{
			"service": modelCacheService,
			"success": true,
		})
	}
}

func (bs BillServices) GetTariffServiceBySlug(c *fiber.Ctx) error {
	paramsSlug := c.Params("slug")
	modelService, err := bs.svcEntity.GetServiceBySlug(paramsSlug)
	//fmt.Println(modelService)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {
		return c.JSON(fiber.Map{
			"tariffs": modelService.Tariffs,
			"success": true,
		})
	}
}

//func findServiceBySlug(services []map[string]interface{}, key string, slug string) map[string]interface{} {
//	for _, service := range services {
//		if service[key] == slug {
//			return service
//		}
//	}
//	return nil
//}

//
//func GetByName(c *fiber.Ctx) error {
//	name := c.Params("name")
//	fmt.Println("service by->", c.Params("name"))
//
//	var dbServices []map[string]interface{}
//	//data, _ := json.Marshal(response.Services)
//	//_ = json.Unmarshal(data, &dbServices)
//	//fmt.Println(dbServices)
//	//
//	searchObject := findServiceBySlug(dbServices, "slug", name)
//	serviceDeviceName, _ := searchObject["full_name"].(string)
//
//	if c.Params("name") != "" {
//		dataResp := fiber.Map{
//			"Title":         "Конфигурирование услуги — " + c.Params("name"),
//			"service":       searchObject,
//			"serviceDevice": strings.Title(serviceDeviceName),
//		}
//		//
//		return response.ResponseTemp(c,
//			"billing/service/service_card",
//			"billing/layout/dashboard",
//			dataResp,
//		)
//	} else {
//		return c.Redirect("/dash")
//	}
//}

//func GetUserServices(user_id int64) {
//
//	var serviceList []service.Services
//	// Fetch services
//	if database.
//		GetDB().
//		Select("*").Find(&serviceList).Error != nil {
//		return c.JSON(fiber.Map{
//			"services": []response.Services{},
//		})
//	} else {
//		//return err
//		return c.JSON(fiber.Map{
//			"services": serviceList,
//		})
//	}
//}

//func GetOld()  {
//
//return response.ResponseTemp(c,
//	"billing/service/entire",
//	"billing/layout/dashboard",
//	dataResp,
//)
//}
