package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterEventHandlers(t *testing.T) {
	c := NewContainer()

	called := false
	c.AddEventHandler("foo", func(e Event) { called = true })

	assert.Equal(t, 1, len(c.eventHandlers))

	event := NewMouseClickEvent(123)
	c.OnEvent(event)
}
