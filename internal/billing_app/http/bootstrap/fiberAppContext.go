package bootstrap

import (
	CacheManagerSync "backend/pkg/cache"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
)

//func dbContext(c *fiber.Ctx) *fiber.Ctx {
//	db, _ := database.InitDB()
//	defer db.DB()
//	c.Context().SetUserValue("db_gorm", database.GetDB())
//	//
//	return c
//}

// localsContext - Подключение locals
func localsContext(c *fiber.Ctx) *fiber.Ctx {
	// Db locals
	//dbContext(c)
	cache := CacheManagerSync.NewCacheManagerSync()
	// Cache locals
	c.Locals("cache_manager", cache)

	//fmt.Println(CacheManagerSync.NewCacheManagerSync())
	//c.Locals("cache_manager", CacheManagerSync.NewCacheManagerSync())

	return c
}

// fiberLocalsContext - Передача чего-либо в контекст
func fiberLocalsContext(app *fiber.App) *fiber.App {
	//// Activate Use
	app.Use(func(c *fiber.Ctx) error {
		// Context
		localsContext(c)
		//return c.Next()
		return c.Next()
	})
	return app
}
func NewFbAppContextService(app *fiber.App) *fiber.App {
	//Worked Locals
	fiberLocalsContext(app)
	return app
}

// NewFbAppLoggerContext - Context Logger Service methods
func NewFbAppLoggerContext(app *fiber.App) *fiber.App {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	//detect fatal error
	//app.Use(recover2.New())

	return app
}
