package bootstrap

import (
	"backend/internal/billing_app/http/bootstrap/eventHook"
	"backend/internal/database"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

//func initDb() *gorm.DB {
//	db, err := database.InitDB()
//	if err != nil {
//		fmt.Printf("Failed to connect to the database: %v\n", err)
//		return nil
//	}
//	defer db.DB()
//	// Глобально вызывайте базу данных, когда это необходимо
//	return database.GetDB()
//}

func NewFiberBootstrap(fiberApp *fiber.App, db *database.StorageDb) {
	NewFbAppLoggerContext(fiberApp)
	// FIBER BOOTSTRAP -> Service Security Config
	NewFbSecurityService(fiberApp)
	//Service Security Config
	NewFbAppContextService(fiberApp)
	//!IMPORTANT POSITION METHOD!
	// FIBER BOOTSTRAP -> Routes
	NewFiberRouterContext(fiberApp, db)
}

func NewHooksServices(fiberApp *fiber.App) {
	//Event BOOTSTRAP HOOKS Service
	eventHook.NewBootFbEventHkService(fiberApp)
	//Event BOOTSTRAP Broadcast Event Service
	eventHook.NewBootFbBroadcastEvents(fiberApp)
}

func NewFiberAppBoot(db *database.StorageDb) *fiber.App {
	////Initialize db
	//initDb()
	//
	fiberApp := fiber.New(
		fiber.Config{
			ServerHeader:  "RetryHost",
			AppName:       "RetryHost $Billing",
			StrictRouting: true,
			CaseSensitive: true,
			//
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
			//Show All Routes
			//EnablePrintRoutes: true,
			//Prefork: true,
			//Views:   engineView,
		})

	//Fiber Bootstrap
	NewFiberBootstrap(fiberApp, db)
	NewHooksServices(fiberApp)

	return fiberApp
}
