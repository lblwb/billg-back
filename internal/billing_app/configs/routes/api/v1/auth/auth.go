package auth

import (
	"backend/internal/billing_app/service/api/auth"
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/database"
	fiber "github.com/gofiber/fiber/v2"
)

type AuthsService struct {
	singInService *auth.SignInService
	singUpService *auth.SignUpService
}

type RoutesApiAuth struct {
	db          *database.StorageDb
	authService *AuthsService
	userService *auth.UserService
}

func NewRoutesApiAuth(db *database.StorageDb) *RoutesApiAuth {
	return &RoutesApiAuth{
		db:          db,
		userService: auth.NewUserService(db),
		authService: &AuthsService{
			singInService: auth.NewSingInService(db),
			singUpService: auth.NewSingUpService(db),
		},
	}
}

func (raa RoutesApiAuth) ApiAuthRoutes(app *fiber.App, group fiber.Router) *fiber.App {
	apiAuthGroup := group.Group("auth")
	apiAuthGroup.Post("/signin", raa.authService.singInService.SignIn)
	apiAuthGroup.Post("/signup", raa.authService.singUpService.SignUp)
	//
	user := apiAuthGroup.Group("user")
	jwt_auth.MiddlewareJwtAuthRoute(app, user, raa.db)
	user.Get("", raa.userService.GetProfile)

	return app
}
