package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButtonEmitsButtonClickedEvent(t *testing.T) {
	b := NewButton("button", font)
	clicked := false
	b.AddEventHandler(ButtonClicked, func(e Event) {
		assert.Equal(t, ButtonClicked, e.Type)
		assert.Equal(t, b, e.Emitter)

		clicked = true
	})

	e := NewMouseClickEvent(123, LMB, ButtonUp, 10, 10)
	b.OnEvent(e)

	assert.True(t, clicked)
}
