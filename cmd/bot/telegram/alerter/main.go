package alerter

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
	"time"
)

// TelegramAlerts Я представляю это как такую вот структуру
type TelegramAlerts struct {
	TgId                 string    `json:"tg_id"`                  //TgId в тг на его чат (Нужно понять как определять его id)
	UserId               uint      `json:"user_id"`                //Id пользователя вкл_ оповещение в тг
	UserServicesDeadline time.Time `json:"user_services_deadline"` //Сервисы которые скоро будут завершены в течение 2 дней
}

type OrderedServices struct {
	ID           int
	DeadlineAt   time.Time
	ServicePrice float64
	Service      Service
	Order        UserOrderServices
	StatusSend   bool
}

type Service struct {
	Id         int
	Slug       string
	Name       string
	FullName   string
	DeviceName string
	DeviceSlug string
	BannerDesc string
	Tariffs    []string
}

type UserOrderServices struct {
	UserID      int
	TotalAmount float64
	PromoCode   string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Services    []string
}

func TgWorker() {

	//
	servicesExp := []OrderedServices{
		{
			ID:           1,
			DeadlineAt:   time.Now().Add(time.Second * 25),
			ServicePrice: 14620,
			Service:      Service{Id: 2, FullName: "Виртуальный сервер 2345"},
			Order:        UserOrderServices{UserID: 1, TotalAmount: 1000.00, Status: "completed"},
			StatusSend:   false,
		},
		{
			ID:           2,
			DeadlineAt:   time.Now().Add(time.Second * 50),
			ServicePrice: 12620,
			Service:      Service{Id: 2, FullName: "Виртуальный сервер 32"},
			Order:        UserOrderServices{UserID: 2, TotalAmount: 1000.00, Status: "completed"},
			StatusSend:   false,
		},
		{
			ID:           3,
			DeadlineAt:   time.Now().Add(time.Second * 120),
			ServicePrice: 14350,
			Service:      Service{Id: 2, FullName: "Виртуальный сервер 43432"},
			Order:        UserOrderServices{UserID: 1, TotalAmount: 1000.00, Status: "completed"},
			StatusSend:   false,
		},
	}

	bot, err := tgbotapi.NewBotAPI("6899671515:AAFBMKrisGxrZxqmwlTvMHmNwDzYoHQhcmM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 13

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			notifier := NewNotifier()
			notifier.ProcessServices(bot, servicesExp)

		}
	}
}

func Run() {
	//TgWorker()
}

type Notifier struct {
	wg               sync.WaitGroup
	notificationChan chan OrderedServices
}

func NewNotifier() *Notifier {
	return &Notifier{
		notificationChan: make(chan OrderedServices),
	}
}

func (n *Notifier) ProcessServices(bot *tgbotapi.BotAPI, servicesExp []OrderedServices) {
	for i, service := range servicesExp {
		n.wg.Add(1)
		go func(i int, service OrderedServices) {
			defer n.wg.Done()

			// Определяем, сколько времени осталось до DeadlineAt
			timeUntilDeadline := time.Until(service.DeadlineAt)
			if timeUntilDeadline > 0 {
				// Ожидаем до наступления DeadlineAt Как-только время наступает отправляеться сообщение)
				time.Sleep(timeUntilDeadline)
			}

			// Здесь можно добавить логику отправки уведомления на service.Order.UserID

			// Завершаем StatusSend
			servicesExp[i].StatusSend = true
			n.notificationChan <- service
		}(i, service)
	}

	go func() {
		n.wg.Wait()
		close(n.notificationChan)
	}()

	n.ProcessNotifications(bot)
}

func (n *Notifier) ProcessNotifications(bot *tgbotapi.BotAPI) {
	for serviceItemChan := range n.notificationChan {

		//
		fmt.Printf("Уведомление отправлено пользователю с ID %d\n", serviceItemChan.Order.UserID)
		messageFormat := fmt.Sprintf("🔥[%d] Скоро срок оплаты! Услуги: %s / #%d / \n К оплате: %f", serviceItemChan.Order.UserID, serviceItemChan.Service.FullName, serviceItemChan.Service.Id, serviceItemChan.ServicePrice)

		keyboardText := "Перейти к оплате"
		//
		sendMessage(bot, 1248134771, messageFormat, true, keyboardText, fmt.Sprintf("https://rtyhost.com/bill/service/%d/pay", serviceItemChan.Order.UserID))
		tgbotapi.NewInlineKeyboardButtonURL("Перейти", "http://pay.ru/")
		//fmt.Println("Отправил оповещение об оплате", userID)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatId int64, textMsg string, keyboardShow bool, keyboardText, keyboardUrl string) int {
	msg := tgbotapi.NewMessage(chatId, textMsg)

	if keyboardShow {
		var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL(keyboardText, keyboardUrl),
			),
		)

		msg.ReplyMarkup = numericKeyboard
		//msg.ReplyToMessageID = update.Message.MessageID
	}

	send, err := bot.Send(msg)
	if err != nil {
		_ = fmt.Errorf("%d", err)
	}
	return send.MessageID
}

func handleChatMessage() {
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID
}
