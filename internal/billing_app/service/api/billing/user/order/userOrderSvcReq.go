package order

import (
	"backend/internal/billing_app/models/order"
	"backend/internal/billing_app/models/tariff"
	"backend/internal/billing_app/models/user"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"time"
)

type Params struct {
	ServicesParams string `json:"pe"`
	TariffID       string `json:"td"`
}

type OrderParamsKey struct {
	Key   string
	Value int
}

//func decodeDataServiceParams(arrayMap map[int]int) {
//	servicesParamsStrJson := decodeBytesIntData(arrayMap)
//}

func (uos BillUsrOrderService) encodeToString(max int) string {
	var table = [...]byte{'1', '2', 'b', '4', '5', 'v', '7', '8', 'e', 'j'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (uos BillUsrOrderService) calculateTotalPrice(jsonData []byte) int {
	log.Println(jsonData)
	// Структура для разбора JSON
	var data map[string]struct {
		Price      int         `json:"price"`
		Value      interface{} `json:"value"`
		OptionSlug string      `json:"option_slug"`
	}

	// Разбор JSON-строки
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return 0
	}

	// Инициализация суммы цен
	totalPrice := 0

	// Перебор элементов в JSON-объекте и подсчёт суммы цен
	for _, item := range data {
		fmt.Println(item)
		switch v := item.Value.(type) {
		case string:
			totalPrice += item.Price
			fmt.Println("total Price -> string: total", totalPrice, "price", item.Price)
		case float64:
			totalPrice += item.Price
			fmt.Println("total Price -> float: total", totalPrice, "price", item.Price)
		default:
			fmt.Printf("Неизвестный тип данных: %T\n", v)
		}

		//totalPrice += value.Price
	}

	// Возвращение общей суммы цен
	return totalPrice
}

// checkoutServicesOrderParams - Проверка доступности суммы заказа на балансе
func (uos BillUsrOrderService) checkoutAvailBalanceByOrderParams(user user.Users, availBalanceByParams float64) (string, error) {
	//Сумма заказа
	//var availBalanceByParams float64
	if availBalanceByParams <= 0 {
		return "not_amount", errors.New("сумма услуг ровна нулю")
	}

	if user.Id == 0 {
		return "not_found", errors.New("not found user")
	} else if user.Balance.Amount >= availBalanceByParams {
		return "withdraw_ok", nil
	} else {
		return "not_avail", errors.New("avail Balance not amount sum")
	}

	//status, _, errorAvail := buo.usersBalanceEntity.GetAvailBalance(UserId, availBalanceByParams)
	//if errorAvail != nil {
	//	fmt.Println(errorAvail)
	//	return "not_get_balance", errors.New("проблема! Получение баланса")
	//} else if !status {
	//	return "not_avail", errorAvail
	//} else {
	//	// Снимаем нужную нам сумму для оплаты услуг
	//	return "withdraw_ok", nil
	//}
}

// checkoutServicesOrderParams - Расчет заказанных пользователем услуг
func (uos BillUsrOrderService) checkoutServicesOrderParams(jsonData []byte) []OrderParamsKey {
	// Структура для разбора JSON
	var data map[string]struct {
		Price int `json:"price"`
		Value int `json:"value"`
	}

	// Разбор JSON-строки
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return []OrderParamsKey{}
	}

	// Инициализация массива с заказанными параметрами
	var orderedParams []OrderParamsKey

	// Перебор элементов в JSON-объекте и сохранение данных о заказанных параметрах
	for key, value := range data {
		orderedParams = append(orderedParams, struct {
			Key   string
			Value int
		}{key, value.Value})
	}

	// Возвращение массива с заказанными параметрами
	return orderedParams
}

//func (buo BillUsrOrderService) decodeBytesIntData(encodedArray map[int]int) string {
//	// Convert the encoded map into a string
//	var decodedString string
//	for i := 0; i < len(encodedArray); i++ {
//		if charCode, found := encodedArray[i]; found {
//			decodedString += string(charCode)
//		}
//	}
//
//	return decodedString
//}

func (uos BillUsrOrderService) insertUserServiceOrder(c *fiber.Ctx,
	userClaim user.Users,
	servicesParamsStr []byte,
	totalAmountByParams float64,
	tariffData tariff.TariffsServices,
) error {

	//var servicesParams map[string]struct {
	//	Price int `json:"price"`
	//	Value int `json:"value"`
	//}

	//err := json.Unmarshal(&servicesParamsStr, &servicesParams)
	//if err != nil {
	//	return err
	//}

	orderId := uos.encodeToString(10)
	userOrder := order.UserOrderServices{
		//ID:
		Slug:        orderId,
		UserID:      userClaim.Id,
		TotalAmount: totalAmountByParams,
		PromoCode:   "",
		Status:      "pending", // Начальный статус заявки
		UpdatedAt:   time.Now(),
		OrderParams: string(servicesParamsStr),
		Services: []order.OrderedServices{{
			TariffID:            tariffData.Id,
			ServiceID:           tariffData.ServiceID,
			ServiceInstructions: "",
			OrderStatus:         "pending",
		}},
	}

	//Check exist domain
	//uos.createNewDnsOrder()

	//{"domain name": {"slug": "", "price": 240, "value": "2311forget.xyz"}}

	var serviceType = ""

	//var (
	//	domain_resolve_type = false
	//	domain_guard_type   = false
	//	server_type         = false
	//)
	paramInfoService := make(map[string]interface{})

	var data map[string]struct {
		Price      int         `json:"price"`
		Value      interface{} `json:"value"`
		OptionSlug string      `json:"option_slug"`
	}

	err := json.Unmarshal(servicesParamsStr, &data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"code":    469,
			"err":     "Упс! Заказ не создан!",
		})
	}

	if len(data) > 0 {
		for s, s2 := range data {
			log.Println("types service", s, s2)

			//if s == "cpu" {
			//	log.Println("Клиент оформил сервер, 0")
			//}

			//
			switch s {
			case "cpu":
				log.Println("Клиент оформил сервер")
				paramInfoService["server"] = s2.Value
				serviceType = "server"
				break
			case "domain sgd name":
				log.Println("Клиент оформил что-то на доменное имя!")
				paramInfoService["domain"] = s2.Value
				serviceType = "domain_guard"
				break
			case "":
				continue
			}

		}
	}

	log.Println(paramInfoService)

	//svcOrdEntity
	//Create Order

	// Записываем в ресурс инфо о текущем ресурсе
	if serviceType == "domain_guard" || serviceType == "domain_resolve" {
		userOrder.Services[0].Resource = paramInfoService["domain"].(string)
		userOrder.Services[0].Type = 4
		userOrder.Status = "accept"
	}

	err = uos.userOrderSvcEntity.InsertUserOrderServices(&userOrder)
	if err != nil {
		fmt.Println("User create new Order Request Err:", err)
		//
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"code":    472,
			"err":     "Упс! Заказ не создан!",
		})
	} else {
		// You can now process the decoded data as needed

		//HandleEvent
		//userOrderEvent.Handle(userOrderEvent.EventUserOrder{
		//	UserOrder: userOrder,
		//})

		// Create an event and send it to the event channel
		//events.NewPushEventToBroadcast("newUserOrder", userOrder)

		switch serviceType {
		case "domain_guard":
			log.Println("Запрос на оформление защиты домена")
			//Создаем на нашем сервисе защиты запрос на новую услугу
			dnsOrder, err := uos.createNewDnsOrder(userOrder.Services[0].Resource)
			if err != nil {
				return err
			}
			log.Println(dnsOrder)

		case "domain_resolve":
			log.Println("Запрос на оформление нового домена")
			//break
		case "server":
			log.Println("Запрос на оформление нового сервера")
			//break
		default:
			log.Println("Тип услуги не определен!")
			//return errors.New("service type not found")

			return c.Status(400).
				JSON(fiber.Map{
					"success": false,
					"error":   "invalid_service",
				})
		}

		return nil
	}

}

// CreateOrderServicesByUser - Создает заказа пользователем на услугу
func (uos BillUsrOrderService) CreateOrderServicesByUser(c *fiber.Ctx) error {
	//Check user
	userClaim, err := uos.jwtAuth.GetUserDataByClaim(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
		})
	} else {

		var params Params
		err := c.BodyParser(&params)
		if err != nil {
			return c.Status(400).
				JSON(fiber.Map{
					"success": false,
					"error":   "invalid",
				})
		}

		//var serviceParams string
		//json.Unmarshal(params.ServicesParams, &serviceParams)
		servicesParamsStr := []byte(params.ServicesParams)

		tariffIdStr := params.TariffID
		//
		tariffData, err := uos.tariffServiceEntity.GetTariffBySlug(tariffIdStr)
		if err != nil {
			fmt.Println("Not Traiff")
		}

		fmt.Println("Service-TariffID:")
		fmt.Println(tariffIdStr)
		fmt.Println(tariffData)

		// Access the decoded data
		fmt.Println("ServicesParams:")
		//
		//fmt.Println("Params Ordered:", buo.checkoutServicesOrderParams(servicesParamsStr))
		//
		fmt.Println("Calc-Total-Price:: ", uos.calculateTotalPrice(servicesParamsStr))

		//
		totalAmountByParams := float64(uos.calculateTotalPrice(servicesParamsStr))

		log.Println("userClaim", userClaim)

		//userTotalAmount := int64(userClaim.Balance.Amount)

		//Проверяем баланс пользователя на наличие суммы заказа
		userBalanceAvail, err := uos.checkoutAvailBalanceByOrderParams(userClaim, totalAmountByParams)
		if err != nil {
			fmt.Println(err)
		}

		switch userBalanceAvail {
		case "not_get_balance":
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"code":    429,
				"err":     "Пополните баланс для заказа услуги!",
			})
		case "not_amount":
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"code":    429,
				"err":     "Пополните баланс для заказа услуги!",
			})
		case "withdraw_ok":
			//Снимаем нужную суммы за заказ с профиля пользователя
			err := uos.insertUserServiceOrder(c, userClaim, servicesParamsStr, totalAmountByParams, tariffData)
			if err == nil {
				success, msg, err := uos.usersBalanceEntity.WithdrawUserBalance(userClaim.Id, totalAmountByParams)
				if success {
					//
					fmt.Println("Услуга успешно создана, баланс списан!")

					return c.Status(http.StatusCreated).JSON(fiber.Map{
						"success": true,
						"code":    463,
						"message": "Заказ успешно создан!",
					})
				} else {
					fmt.Println("Баланс пользователя не был списан!", err)
					fmt.Println("--", msg)
				}

			} else {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"success": false,
					"code":    488,
					"err":     "Усп! Что-то пошло не так! Уже разбираемся",
				})
			}

		default:
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"code":    491,
				"err":     "Усп! Что-то пошло не так! Уже разбираемся",
			})
		}
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"code":    486,
		"err":     "Усп! Что-то пошло не так! Уже разбираемся",
	})
}
