package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"net/http"
	"time"
)

func NewFbSecurityService(app *fiber.App) *fiber.App {

	//CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Разрешенные источники
		//AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
		AllowCredentials: true,
	}))

	// Safe App to XSS and security http
	app.Use(helmet.New(helmet.Config{
		//XSSProtection: "1",
		XSSProtection: "0",
		CSPReportOnly: true,
	}))

	// Safe to network duplicate.request problem
	app.Use(idempotency.New(idempotency.Config{
		Lifetime:  2 * time.Minute,
		KeyHeader: "X-Idp-4j4bv1b3b1",
		// ...
	}))

	// Safe to Many Requests
	app.Use(limiter.New(limiter.Config{
		Max: 64, // max count of connections
		//LimiterMiddleware: {},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error": "many request",
			})
		},
	}))

	return app
}
