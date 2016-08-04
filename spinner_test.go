package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpinnerEmitsEventOnChange(t *testing.T) {
	s := NewSpinner(font)
	s.Step = 5.0
	s.SetDimensions(100, 20)

	called := false
	actual := 0.0

	s.AddEventHandler(SpinnerChanged, func(e *Event) bool {
		called = true

		data, _ := e.Data.(SpinnerChangedEvent)
		actual = data.NewValue

		return false
	})

	s.plusButton.OnEvent(&Event{
		Type: ButtonReleased,
		Data: ButtonClickEvent{},
	})

	assert.True(t, called)
	assert.Equal(t, 5.0, actual)
}
