package eventHook

import (
	events_broadcast "backend/internal/billing_app/events"
	"backend/pkg/events"
	"github.com/gofiber/fiber/v2"
	"log"
)

func NewBootFbBroadcastEvents(app *fiber.App) *fiber.App {
	go broadcastEvents()
	//
	return app
}

// Worker
func broadcastEvents() {
	for {
		// Broadcast the event to all subscribers
		// You can implement this logic as needed
		// For example, you can maintain a list of subscribers and send the event to each of them.

		event := <-events.EventChannel
		log.Println("new Event -> "+event.Name+" -> by Broadcast Channel: \n", event.Data)
		events_broadcast.Subscribers(event)
	}
}
