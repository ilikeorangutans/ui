package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButtonEmitsButtonReleasedEvent(t *testing.T) {
	b := NewClickButton("button", font)
	clicked := false
	b.AddEventHandler(ButtonReleased, func(e *Event) bool {
		assert.Equal(t, ButtonReleased, e.Type)
		assert.Equal(t, b, e.Emitter)

		clicked = true
		return true
	})

	b.OnEvent(&Event{Type: "MouseClickEvent", Data: MouseClickEvent{Button: LMB, State: ButtonDown}})
	e := NewMouseClickEvent(123, LMB, ButtonUp, 10, 10)
	b.OnEvent(e)

	assert.True(t, clicked)
}
