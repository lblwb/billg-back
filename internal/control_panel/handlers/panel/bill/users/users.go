package users

import (
	"backend/internal/billing_app/models/user"
	"backend/internal/database"
	"backend/pkg/control_panel/http/inertia"
	"github.com/gofiber/fiber/v2"
)

type HandlerUsers struct {
	////db *database.StorageDb
	////resp *inertia.ResponseInertia
	//appInt *bootstrap.AppInt
	resp       *inertia.ResponseInertia
	db         *database.StorageDb
	userEntity *user.UsersEntity
}

func (ho HandlerUsers) GroupUsers(router fiber.Router) {

	router.Group("clients").
		Get("all", func(ctx *fiber.Ctx) error {

			allUsers, err := ho.userEntity.GetAllUsers()
			if err != nil {
				return err
			}

			return ho.resp.RenderEngine(ctx).
				View("Bill/Users/Index",
					fiber.Map{
						"users": allUsers,
					},
					ctx)
		})
}

func NewHandlerUsers(db *database.StorageDb, resp *inertia.ResponseInertia) *HandlerUsers {
	return &HandlerUsers{
		//appInt: appInt,
		db:         db,
		resp:       resp,
		userEntity: user.NewUsersEntity(db),
	}
}
