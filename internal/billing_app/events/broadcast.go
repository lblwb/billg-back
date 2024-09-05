package events

import (
	"backend/internal/billing_app/events/userOrderEvent"
	"backend/pkg/events"
	"log"
)

func Subscribers(event events.Event) {
	switch event.Name {

	case "newUserOrder":
		//Новые заказы
		userOrderEvent.Handle(event)

	case "newRegisterUser":
		log.Println("Новая регистрация!")
		//Новые регистрации
		//userOrderEvent.Handle(event)
	}
}
