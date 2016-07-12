package ui

import (
	"fmt"
	"log"
)

type Event struct {
	Timestamp uint32
	Type      string
	Data      interface{}
}

func (e Event) String() string {
	return fmt.Sprintf("[%dms] %s", e.Timestamp, e.Type)
}

type EventHandlerFunc func(event Event)

type EventHandlers struct {
	eventHandlers map[string][]EventHandlerFunc
}

func (h *EventHandlers) AddEventHandler(t string, handler EventHandlerFunc) {
	if h.eventHandlers == nil {
		h.eventHandlers = make(map[string][]EventHandlerFunc)
	}
	h.eventHandlers[t] = append(h.eventHandlers[t], handler)
}

func (h *EventHandlers) OnEvent(event Event) {
	log.Printf("EventHandlers.OnEvent(%s)\n", event)
	if h.eventHandlers == nil || len(h.eventHandlers) == 0 {
		return
	}

	handlers := h.eventHandlers[event.Type]
	if len(handlers) == 0 {
		return
	}

	for i := range handlers {
		handlers[i](event)
	}
}
