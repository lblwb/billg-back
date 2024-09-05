package auth

import (
	userModel "backend/internal/billing_app/models/user"
	"backend/internal/database"
	"backend/pkg/events"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SignUpService struct {
	db         *database.StorageDb
	userEntity *userModel.UsersEntity
}

func NewSingUpService(db *database.StorageDb) *SignUpService {
	userEntity := userModel.NewUsersEntity(db)
	return &SignUpService{
		db:         db,
		userEntity: userEntity,
	}
}

type UserSignUp struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (usu UserSignUp) AfterCreate(tx *gorm.DB) (err error) {
	events.NewPushEventToBroadcast("newRegisterUser", tx)
	return
}

// TODO: Register!

// SignUp @Router /api/v1/auth/signup
func (sus SignUpService) SignUp(c *fiber.Ctx) error {
	user := new(UserSignUp)
	if err := c.BodyParser(user); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	//Тело-запроса
	userReqs := new(UserSignIn)
	if err := c.BodyParser(userReqs); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	createUser, err := sus.userEntity.CreateUser(userReqs.User, userReqs.Password)
	if err != nil {
		return err
	}

	fmt.Println(createUser)

	//var users []UserSignUp
	//
	//for i := 0; i < 900000; i++ {
	//	users = append(users, UserSignUp{
	//		User:     fmt.Sprintf("user%dname", i),
	//		Password: fmt.Sprintf("2432%drwerew", i),
	//	})
	//}

	return c.JSON(fiber.Map{"success": true})
}

//func SignUpView(c *fiber.Ctx) error {
//	return c.Render("auth/signup", fiber.Map{
//		"Title": "Регистрация",
//	}, "auth/layout/auth")
//}
