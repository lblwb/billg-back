package auth

import (
	userModel "backend/internal/billing_app/models/user"
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserSignIn struct {
	User     string `json:"login_email"`
	Password string `json:"password"`
}

type SignInService struct {
	db         *database.StorageDb
	jwtAuth    *jwt_auth.JwtAuths
	userEntity *userModel.UsersEntity
	data       *UserSignIn
}

func NewSingInService(db *database.StorageDb) *SignInService {
	return &SignInService{
		db:         db,
		jwtAuth:    jwt_auth.NewJwtAuths(db),
		userEntity: userModel.NewUsersEntity(db),
		data:       &UserSignIn{},
	}
}

//createUser, err := userModel.CreateUser(userReqs.User, userReqs.Password)
//if err != nil {
//	return err
//}

// SignIn @Router /api/v1/auth/signin
func (sis SignInService) SignIn(c *fiber.Ctx) error {
	log.Println(sis.data)

	//Тело-запроса
	if err := c.BodyParser(sis.data); err != nil {
		fmt.Println("body parser error ->", err)
		fmt.Println(fiber.Map{"error": true, "err_msg": "not auth!"})
		//
		return c.Status(403).JSON(fiber.Map{"error": true, "err_msg": "not auth!"})
	}

	fmt.Println(sis.jwtAuth)

	// Получение пользователя
	var user userModel.Users
	user, err := sis.userEntity.GetUserByLogin(userModel.AuthUser{Username: sis.data.User})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials #592",
		})
	}

	// Сверка шифра пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sis.data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials #842",
		})
	}

	// Выдача подписанного токена
	t, claims, err := sis.jwtAuth.GenerateJwtTokenWithClaims(user.Username, []byte(sis.data.Password))
	if err != nil {
		fmt.Println("Error Jwt Auth | User {}", user)
	}

	user, err = sis.userEntity.GetUserByLogin(userModel.AuthUser{Username: user.Username})
	//userAndUsd := user.CalcExchangesRates("USD")

	if err != nil {
		return c.JSON(fiber.Map{"info": claims, "user": fiber.Map{}, "auth": "Bearer", "act": t})
	} else {
		return c.JSON(fiber.Map{"info": claims, "user": user, "auth": "Bearer", "act": t})
	}
}

//func SignInView(c *fiber.Ctx) error {
//	return c.Render(
//		"auth/signin", fiber.Map{
//			"Title": "Авторизация",
//		},
//		"layout/main")
//}
