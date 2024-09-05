package auth

import (
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	jwtAuth *jwt_auth.JwtAuths
	db      *database.StorageDb
}

func NewUserService(db *database.StorageDb) *UserService {
	return &UserService{db: db, jwtAuth: jwt_auth.NewJwtAuths(db)}
}

// GetProfile @Router /api/v1/auth/user
func (us UserService) GetProfile(c *fiber.Ctx) error {
	user, err := us.jwtAuth.GetUserDataByClaim(c)
	//userAndExc := user.CalcExchangesRates("USD")
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"user":    nil,
		})
	} else {
		return c.JSON(fiber.Map{
			"success": true,
			"user":    user,
			"balance": 0,
		})
	}
}
