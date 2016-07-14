package ui

import "fmt"

// Event is a single event the ui framework is interested in. An Event has a timestamp (relative to app startup), a string Type, and a Data field with event specific payload.
type Event struct {
	// When the event occured
	Timestamp uint32
	// Type of the event
	Type string
	// Widget that emitted this event. Might be nil.
	Emitter Widget
	// Event payload. Might be nil.
	Data interface{}
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
