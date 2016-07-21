package ui

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestRegisterEventHandlers(t *testing.T) {
	c := NewContainer()

	called := false
	c.AddEventHandler("MouseClickEvent", func(e Event) { called = true })

	assert.Equal(t, 1, len(c.eventHandlers))

	event := NewMouseClickEvent(123, LMB, ButtonDown, 10, 10)
	c.OnEvent(event)

	assert.True(t, called)
}

func TestNestedContainerLayoutChildren(t *testing.T) {
	log.Printf("XXX")
	parent := NewContainer()
	parent.SetDimensions(200, 100)
	//parent.SetBounds(0, 0, 200, 100)
	// TODO BUG it's not asigning bounds when only assigning dimensions!
	child := NewContainer()
	child.SetDimensions(0, 0)
	parent.Add(child)

	parent.Layout()

	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, child.Bounds())
}
