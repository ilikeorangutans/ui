package ui

import "fmt"

type EventType string

const (
	ValueChanged EventType = "ValueChanged"
)

// Event is a single event the ui framework is interested in. An Event has a timestamp (relative to app startup), a string Type, and a Data field with event specific payload.
type Event struct {
	// When the event occured
	Timestamp uint32
	// Type of the event
	Type EventType
	// Widget that emitted this event. Might be nil.
	Emitter Widget
	// Event payload. Might be nil.
	Data interface{}
	// Flag if the event has been handled yet
	Handled bool
}

func (e Event) String() string {
	return fmt.Sprintf("[%dms] %s", e.Timestamp, e.Type)
}

// EventHandlerFunc is the interface for event handling functions. They take an event
// as parameter and return true if they handled the event.
type EventHandlerFunc func(event *Event) bool

type EventHandlers struct {
	eventHandlers map[EventType][]EventHandlerFunc
}

func (h *EventHandlers) AddEventHandler(t EventType, handler EventHandlerFunc) {
	if h.eventHandlers == nil {
		h.eventHandlers = make(map[EventType][]EventHandlerFunc)
	}
	h.eventHandlers[t] = append(h.eventHandlers[t], handler)
}

func (h *EventHandlers) OnEvent(event *Event) {
	if h.eventHandlers == nil || len(h.eventHandlers) == 0 {
		return
	}

	handlers := h.eventHandlers[event.Type]
	if len(handlers) == 0 {
		return
	}

	for i := range handlers {
		if handlers[i](event) {
			event.Handled = true
			break
		}
	}
}
