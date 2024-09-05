package events

import "time"

type Event struct {
	Data interface{}
	Name string
}

var EventChannel = make(chan Event)

// NewPushEventToBroadcast - Push Event the Broadcast channel by name
func NewPushEventToBroadcast(Name string, Data interface{}) {
	event := Event{Name: Name, Data: Data}
	EventChannel <- event
}

// NewPushArrEventToBroadcast - Push Arr Event the Broadcast channel by name
func NewPushArrEventToBroadcast(events []Event) {
	//event := []Event
	for _, event := range events {
		if event.Data != nil {
			EventChannel <- event
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
}
