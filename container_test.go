package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestRegisterEventHandlers(t *testing.T) {
	c := NewHorizontalContainer()

	called := false
	c.AddEventHandler("MouseClickEvent", func(e *Event) bool {
		called = true
		return true
	})

	assert.Equal(t, 1, len(c.eventHandlers))

	event := NewMouseClickEvent(123, LMB, ButtonDown, 10, 10)
	c.OnEvent(event)

	assert.True(t, called)
}

func TestNestedContainerLayoutChildren(t *testing.T) {
	parent := NewHorizontalContainer()
	parent.SetDimensions(200, 100)
	//parent.SetBounds(0, 0, 200, 100)
	// TODO BUG it's not asigning bounds when only assigning dimensions!
	child := NewHorizontalContainer()
	child.SetDimensions(0, 0)
	parent.Add(child)

	parent.Layout()

	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, child.Bounds())
}

func TestContainerAppliesMarginAndPadding(t *testing.T) {
	c := NewHorizontalContainer()
	c.SetDimensions(400, 200)
	c.SetMargin(Margin{3, 3, 3, 3})
	c.SetPadding(Margin{5, 5, 5, 5})

	c.Layout()

	assert.Equal(t, &sdl.Rect{3, 3, 394, 194}, c.BorderArea())
	assert.Equal(t, &sdl.Rect{8, 8, 384, 184}, c.WidgetArea())
}

func TestContainerClearRemovesAllChildren(t *testing.T) {
	c := NewHorizontalContainer()
	c.Add(NewSpacer(0, 0))
	c.Add(NewSpacer(0, 0))

	c.Clear()
	assert.Len(t, c.Children(), 0)
}
