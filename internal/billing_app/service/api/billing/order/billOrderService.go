package order

/**
Сервис для работы
с заказами
*/

import (
	"backend/internal/billing_app/models/order"
	"backend/internal/billing_app/models/service"
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type BillOrderService struct {
	db             *database.StorageDb
	svcOrdEntity   *order.ServiceOrdersEntity
	servicesEntity *service.ServicesEntity
	jwtAuth        *jwt_auth.JwtAuths
}

func NewBillOrderService(db *database.StorageDb) *BillOrderService {
	return &BillOrderService{
		//orderEntity:
		db:             db,
		svcOrdEntity:   order.NewServiceOrdersEntity(db),
		servicesEntity: service.NewServicesEntity(db),
		jwtAuth:        jwt_auth.NewJwtAuths(db),
	}
}

// GetPreviousServicesByUser - Список сделанных заказов по данной услуге
func (bos BillOrderService) GetPreviousServicesByUser(c *fiber.Ctx) error {
	paramsSlug := c.Params("slug")

	modelService, err := bos.servicesEntity.GetServiceBySlug(paramsSlug)
	fmt.Println(modelService)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {

		return c.JSON(fiber.Map{
			"service":  modelService,
			"services": []service.Services{},
			"count":    0,
			"success":  true,
		})
	}

}

func (bos BillOrderService) GetOrdersByUserId(c *fiber.Ctx) error {
	user, err := bos.jwtAuth.GetUserDataByClaim(c)
	//
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {
		//
		modelServices := bos.svcOrdEntity.GetServiceOrdersByUserId(user.Id)

		return c.JSON(fiber.Map{
			"orders":  modelServices,
			"count":   len(modelServices),
			"success": true,
		})
	}
}

func (bos BillOrderService) GetOrdersServicesByUserId(c *fiber.Ctx) error {
	user, err := bos.jwtAuth.GetUserDataByClaim(c)
	//
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {
		//
		modelServices := bos.svcOrdEntity.GetOrdersByUserId(user.Id)

		return c.JSON(fiber.Map{
			"services": modelServices,
			"count":    len(modelServices),
			"success":  true,
		})
	}
}
