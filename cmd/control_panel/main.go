package main

import (
	"backend/internal/control_panel/bootstrap"
	"backend/internal/control_panel/middleware"
	"backend/internal/database"
	"backend/pkg"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"io/fs"
	"log"
	"os"
)

func renderEngine() *inertia.Engine {
	log.Println("exist:", fs.ValidPath("control_panel"))

	var rootPath = ""

	if fs.ValidPath("control_panel") {
		rootPath = "././internal/control_panel"
	}

	return inertia.New(inertia.Config{
		Root:         rootPath + "/resources/views",
		AssetsPath:   rootPath + "/resources/js",
		Template:     "app",
		ManifestRoot: rootPath,
	})
}

func dbConnect() *database.StorageDb {
	return database.
		NewStorageDb(os.Getenv("DB_DSN")).
		Connect()
}

func main() {
	pkg.LoadEnv()

	dbSql := dbConnect()
	if dbSql.Ping() == nil {
		log.Println("DB CONNECTED!")
	}

	viewEngine := renderEngine()

	app := fiber.New(
		fiber.Config{
			ServerHeader:  "RetryHost Control",
			AppName:       "RetryHost Control $Billing",
			StrictRouting: true,
			CaseSensitive: true,
			//
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
			//Show All Routes
			EnablePrintRoutes: false,
			//Prefork: true,
			Views: viewEngine,
		})

	app.Static("/assets", "./public/")
	//app.Static("resources", "./internal/control_panel/resources/")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.NewRenderLocMidl(viewEngine))

	instanceApp := bootstrap.NewAppInstance(app, dbSql)
	//
	app.Route("panel/bill", func(router fiber.Router) {
		instanceApp.Handlers.Orders.GroupOrders(router)
		instanceApp.Handlers.Services.GroupServices(router)
		instanceApp.Handlers.Users.GroupUsers(router)
	})
	//.
	//All("*", func(ctx *fiber.Ctx) error {
	//	return ctx.RedirectBack("")
	//})
	log.Fatal(app.Listen(":7722"))
}
