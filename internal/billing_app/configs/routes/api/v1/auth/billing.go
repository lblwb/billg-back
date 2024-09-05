package auth

import (
	"backend/internal/billing_app/configs/routes/api/v1/payment"
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/billing_app/service/api/billing/order"
	"backend/internal/billing_app/service/api/billing/service"
	billUserOrder "backend/internal/billing_app/service/api/billing/user/order"
	"backend/internal/database"
	fiber "github.com/gofiber/fiber/v2"
	//billUserOrder "backend/app/service/api/billing/user/order"
)

type BillRoutes struct {
	db                  *database.StorageDb
	billOrderService    *order.BillOrderService
	billServices        *service.BillServices
	billUsrOrderService *billUserOrder.BillUsrOrderService
	billPayGwService    *payment.BillPayGatewayService
}

func NewBillRoutes(db *database.StorageDb) *BillRoutes {
	return &BillRoutes{
		db:                  db,
		billOrderService:    order.NewBillOrderService(db),
		billServices:        service.NewBillServices(db),
		billUsrOrderService: billUserOrder.NewBillUsrOrderService(db),
		billPayGwService:    payment.NewBillPayGatewayService(db),
	}
}

func (br BillRoutes) billOrderGroup(billingGroup fiber.Router) {
	//billingGroup.Get("/dash", billing.Dashboard)

	//Работа с заказами
	ordersBillingGroup := billingGroup.Group("orders")
	ordersBillingGroup.Get("", br.billOrderService.GetOrdersByUserId)
	ordersBillingGroup.Get("services/entire", br.billOrderService.GetOrdersServicesByUserId)
	//
	br.billSvcCfOrdGroup(billingGroup)
}

func (br BillRoutes) billSvcCfOrdGroup(ordersBillingGroup fiber.Router) {

	serviceOrdBillGroup := ordersBillingGroup.Group("services")
	serviceOrdBillGroup.Get("/:slug/info", br.billUsrOrderService.OrderServiceShow)

	// Конфигурирование-услуг
	servicesConfigureOrder := serviceOrdBillGroup.Group("configure")
	// Создание заявки на услугу
	servicesConfigureOrder.Post("request", br.billUsrOrderService.CreateOrderServicesByUser)
}

func (br BillRoutes) billServiceGroup(billingGroup fiber.Router) {
	//Работа с услугами
	serviceBillingGroup := billingGroup.Group("service")
	serviceBillingGroup.Get("/entire", br.billServices.GetEntireServices)
	//serviceBillingGroup.Get("/:name<string>?", service.GetByName)

	serviceBillingGroup.Get("/:slug/show", br.billServices.GetServiceBySlug)
	serviceBillingGroup.Get("/:slug/tariffs", br.billServices.GetTariffServiceBySlug)
	//
	//Проверяем заказывал ли ранее пользовать данную услугу
	serviceBillingGroup.Get("/:slug/previous", br.billOrderService.GetPreviousServicesByUser)
	//serviceBillingGroup.Post("/:slug/")
	//
}

func (br BillRoutes) billPaymentGroup(billingGroup fiber.Router) {
	//Работа с платежными системами
	paymentBillingGroup := billingGroup.Group("pay")
	payGatewayBillGroup := paymentBillingGroup.Group("gw")
	payGatewayBillGroup.Post(":gw_name", br.billPayGwService.CreateGateway)
	payGatewayBillGroup.Get("list", br.billPayGwService.GetAllGateway)
}

func (br BillRoutes) BillingRoutes(app *fiber.App, group fiber.Router) *fiber.App {
	billingGroup := group.Group("bill")
	//Auth middleware
	jwt_auth.MiddlewareJwtAuthRoute(app, billingGroup, br.db)
	// -------- ORDER ------
	br.billOrderGroup(billingGroup)
	// -------- SERVICE ------
	br.billServiceGroup(billingGroup)

	br.billPaymentGroup(billingGroup)

	return app
}

// BillingPrivateRoutes - hidden services - Todo: Использовать для скрытых услуг
func BillingPrivateRoutes(app *fiber.App) *fiber.App {
	return app
}
