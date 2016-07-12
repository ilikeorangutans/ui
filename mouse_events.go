package ui

type MouseButton int

const (
	LMB = iota
	MMB
	RMB
)

func NewMouseClickEvent(timestamp uint32) Event {
	return Event{
		Timestamp: timestamp,
		Type:      "MouseClickEvent",
		Data:      MouseClickEvent{},
	}
}

type MouseClickEvent struct {
	Event
	Button MouseButton
}

type MouseClickHandler interface {
	OnMouseClick(e MouseClickEvent)
}

func (e MouseClickEvent) Notify(w Widget) {
	t, ok := w.(MouseClickHandler)
	if !ok {
		return
	}

	t.OnMouseClick(e)
}
