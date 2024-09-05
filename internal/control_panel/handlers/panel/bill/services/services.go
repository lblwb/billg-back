package services

import (
	"backend/internal/billing_app/models/service"
	"backend/internal/database"
	"backend/pkg/control_panel/http/inertia"
	"github.com/gofiber/fiber/v2"
)

type HandlerServices struct {
	////db *database.StorageDb
	////resp *inertia.ResponseInertia
	//appInt *bootstrap.AppInt

	db             *database.StorageDb
	resp           *inertia.ResponseInertia
	servicesEntity *service.ServicesEntity
}

func (hs HandlerServices) GroupServices(router fiber.Router) {
	router.Group("services").
		Get("all", hs.GetAllServices).
		//
		Get(":slug/info", hs.GetInfoService).
		Get(":slug/update", hs.GetUpdateService).
		//Update PostData Info
		Post(":slug/update", hs.updateServiceInfo)
}

func NewHandlerServices(db *database.StorageDb, resp *inertia.ResponseInertia) *HandlerServices {
	return &HandlerServices{
		//appInt: appInt,
		db:             db,
		resp:           resp,
		servicesEntity: service.NewServicesEntity(db),
	}
}

func (hs HandlerServices) GetAllServices(ctx *fiber.Ctx) error {
	return hs.resp.RenderEngine(ctx).
		View("Bill/Services/Index",
			fiber.Map{
				"services": hs.servicesEntity.GetAllServices(),
			},
			ctx)
}

func (hs HandlerServices) GetInfoService(ctx *fiber.Ctx) error {
	serviceSlug := ctx.Params("slug")
	serviceBySlug, err := hs.servicesEntity.GetServiceBySlug(serviceSlug)
	if err != nil {
		return err
	}
	return hs.resp.RenderEngine(ctx).
		View("Bill/Services/Action/Info",
			fiber.Map{
				"serviceData": serviceBySlug,
				"serviceId":   serviceSlug,
			},
			ctx)
}

func (hs HandlerServices) GetUpdateService(ctx *fiber.Ctx) error {
	serviceSlug := ctx.Params("slug")
	serviceBySlug, err := hs.servicesEntity.GetServiceBySlug(serviceSlug)
	if err != nil {
		return err
	}
	return hs.resp.RenderEngine(ctx).
		View("Bill/Services/Action/Edit",
			fiber.Map{
				"serviceData": serviceBySlug,
				"serviceId":   serviceSlug,
			},
			ctx)
}

func (hs HandlerServices) updateServiceInfo(ctx *fiber.Ctx) error {

	//TODO: Обновление данных

	return nil

	//serviceSlug := ctx.Params("slug")
	//serviceBySlug, err := hs.servicesEntity.GetServiceBySlug(serviceSlug)
	//if err != nil {
	//	return err
	//}
	//return hs.resp.RenderEngine(ctx).
	//	View("Bill/Services/Index",
	//		fiber.Map{
	//			"serviceInfo": serviceBySlug,
	//			"serviceId":   serviceSlug,
	//		},
	//		ctx)
}
