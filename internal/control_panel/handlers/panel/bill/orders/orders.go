package orders

import (
	"backend/internal/billing_app/models/order"
	"backend/internal/database"
	"backend/pkg/control_panel/http/inertia"
	"github.com/gofiber/fiber/v2"
)

type HandlerOrders struct {
	//db     *database.StorageDb
	db            *database.StorageDb
	resp          *inertia.ResponseInertia
	userSvcEntity *order.UserOrderSvcEntity
}

func (ho HandlerOrders) GroupOrders(router fiber.Router) *fiber.Router {
	var routerGroups = router.Group("orders").
		//
		Get("all", func(ctx *fiber.Ctx) error {

			usersSvcOrders, err := ho.userSvcEntity.GetAllUserOrdersSvc()
			if err != nil {
				return err
			}

			var orderServices []order.OrderedServices

			for _, svcOrder := range usersSvcOrders {
				//log.Println(svcOrder.Services)
				orderServices = append(orderServices, svcOrder.Services...)
				//usersSvcOrders[i].Services = svcOrder.Services
				//usersSvcOrders[i].Services = usersSvcOrders[i].Services
			}

			return ho.resp.RenderEngine(ctx).
				View("Bill/Orders/Index", fiber.Map{
					"userSvcOrders": orderServices,
				}, ctx)
		}).
		Get("*", func(ctx *fiber.Ctx) error {
			return ho.resp.RenderEngine(ctx).
				View("Index", fiber.Map{}, ctx)
		})
	return &routerGroups
}

func NewHandlerOrders(db *database.StorageDb, resp *inertia.ResponseInertia) *HandlerOrders {
	return &HandlerOrders{
		resp:          resp,
		db:            db,
		userSvcEntity: order.NewUserOrderSvcEntity(db),
	}
}
