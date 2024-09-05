package userOrderEvent

import (
	"backend/internal/billing_app/models/order"
	"backend/pkg/events"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

//type EventUserOrder struct {
//	// Данные события
//	UserOrder order.UserOrderServices
//}

// Ваш токен бота
const telegramBotToken = "6605794583:AAFmhnarpaJrrdRLQh9YPSIilBlM76KT8Ng"

// ID чата админа, куда будет отправляться сообщение
const adminChatID = 1248134771

func telegramSendMsgIds(adminsIds []int64, combinedMsgString string, keyboardRow []tgbotapi.InlineKeyboardButton) {
	for _, ChatId := range adminsIds {
		//fmt.Println(ChatId, combinedMsgString)

		// Отправляем сообщение в Telegram
		//if userTgId != 0 {
		msg := tgbotapi.NewMessage(int64(ChatId), combinedMsgString)
		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboardRow)
		//

		bot, _ := tgbotapi.NewBotAPI(telegramBotToken)
		// Отправьте сообщение
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("err: ", err)
		}
	}
}

func sendNewOrderByTelegram(adminsIds []int64, service *order.UserOrderServices) {
	var arrMsg []string
	//var paramTagsText string
	var deviceName string

	if len(service.Services) > 0 {
		//fmt.Println("SERVDF===>", service.Services[0].Service)
		deviceName = service.Services[0].Service.FullName
	}

	// Добавляем строки в срез
	arrMsg = append(arrMsg,
		"💸🔥 ——— НОВЫЙ ЗАКАЗ ——— 💸🔥\n",
		"----",
		"Оборудование -> "+deviceName,
		//service.,
		"----\n",
	)

	if len(service.OrderParams) > 0 {
		//if paramTagsText != "" {
		//	tags = strings.Split(paramTagsText, ",")
		//}

		// Используем map[string]interface{} для динамического хранения данных
		var data map[string]map[string]interface{}

		if err := json.Unmarshal([]byte(service.OrderParams), &data); err != nil {
			fmt.Println("Ошибка при разборе JSON:", err)
		}

		arrMsg = append(arrMsg, fmt.Sprintf("☕️🤓 Услуги: \n"))

		// Перебираем все компоненты и выводим данные
		for componentName, componentData := range data {
			arrMsg = append(arrMsg, fmt.Sprintf("——> %s: | Цена: %v | Количество %v \n", componentName, componentData["price"], componentData["value"]))
			//fmt.Printf("%s:\n", componentName)
			//fmt.Printf("price: %v\ncount: %v\n", componentData["price"], componentData["value"])
			//fmt.Println()
		}

	}

	// Объединяем все строки в одну строку с переносами строк
	combinedMsgString := strings.Join(arrMsg, "\n")

	//
	var keyboardRow = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL(
			fmt.Sprintf("💰Принять заказ — сум. %d Руб.", int64(service.TotalAmount)),
			"https://retryhost.ru/"),
	)

	// Send msg
	telegramSendMsgIds(adminsIds, combinedMsgString, keyboardRow)
}

func Handle(event events.Event) {
	//
	userOrderService := event.Data.(*order.UserOrderServices)

	sendNewOrderByTelegram([]int64{
		1248134771,
		2015896349,
	}, userOrderService)

	//fmt.Println(event)

	// Обработка события NewOrderEvent
	//fmt.Printf("Received a new order event: OrderID=%d, OrderName=%s\n", userOrderService.ID, userOrderService.OrderParams)
	// Вы можете выполнять здесь необходимую логику, например, сохранение в базу данных или отправку уведомлений.
}
