package bootstrap

import (
	"backend/internal/control_panel/handlers/panel/bill/orders"
	"backend/internal/control_panel/handlers/panel/bill/services"
	"backend/internal/control_panel/handlers/panel/bill/users"
	"backend/internal/database"
	"backend/pkg/control_panel/http/inertia"
	"github.com/gofiber/fiber/v2"
)

type AppInt struct {
	App      *fiber.App
	Db       *database.StorageDb
	Handlers *AppIntHandlers
	Resp     *inertia.ResponseInertia
}

// AppIntHandlers Handlers Connect
type AppIntHandlers struct {
	Orders   *orders.HandlerOrders
	Services *services.HandlerServices
	Users    *users.HandlerUsers
}

func NewAppInstance(app *fiber.App, db *database.StorageDb) *AppInt {
	appInstance := &AppInt{
		App: app,
		Db:  db,
	}
	resp := inertia.NewRespInertia()
	//Handlers
	appInstance.Handlers = &AppIntHandlers{
		Orders:   orders.NewHandlerOrders(db, resp),
		Services: services.NewHandlerServices(db, resp),
		Users:    users.NewHandlerUsers(db, resp),
	}

	return appInstance
}
