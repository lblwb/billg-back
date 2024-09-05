package response

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	TgId     string `json:"tgId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Balance  int64  `json:"balance"`
}

type Tariffs struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Services struct {
	Id         int    `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	DeviceName string `json:"device_name"`
	DeviceSlug string `json:"device_slug"`
	//Tariffs    []Tariffs `json:"tariffs"`
}

//var ServicesTariffs = []Tariffs{{Id: 1, Name: "Light"}, {Id: 2, Name: "Premium"}}

//var Services = []Service{
//	{Id: 1, Slug: "abuse", FullName: "Абузоустойчивые сервера", Name: "Абузоустойчивый", DeviceName: "Сервер", DeviceSlug: "vps", Tariffs: ServicesTariffs},
//	{Id: 2, Slug: "vps", FullName: "Виртуальные сервера", Name: "Виртуальный", DeviceName: "Сервер", DeviceSlug: "vps", Tariffs: ServicesTariffs},
//	{Id: 3, Slug: "web", FullName: "Виртуальный хостинг", Name: "Виртуальный", DeviceName: "Сервер", DeviceSlug: "web", Tariffs: ServicesTariffs},
//	{Id: 4, Slug: "dedik", FullName: "Выделенные сервера", Name: "Выделенный", DeviceName: "Сервер", DeviceSlug: "dedik", Tariffs: ServicesTariffs},
//	{Id: 5, Slug: "netguard", FullName: "Защита сети", Name: "Защита", DeviceName: "Сеть", DeviceSlug: "guard", Tariffs: ServicesTariffs},
//}

func ResponseTemp(ctx *fiber.Ctx, viewPath string, layoutPath string, data fiber.Map) error {
	// Пример модели авторизованного пользователя
	//user := User{Id: "r423432e", Username: "label.exe", TgId: "4444444", Email: "abuse@retry.host", Password: "$wrewwerw", Balance: 3000}

	//balance := fmt.Sprintf("%d", user.Balance)
	//balanceFormat := strings.Join([]string{balance[:1], balance[1:]}, ".")
	//
	//ctx.Locals("username", user.Username)
	//ctx.Locals("balance", balanceFormat)

	// Передаем список услуг в контекст
	//ctx.Locals("services", Services)

	data["c"] = ctx
	//
	return ctx.Render(viewPath, data, layoutPath)
}
