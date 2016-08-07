package ui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestLocatingVisitorReturnsContainer(t *testing.T) {
	f := &LocatingFinder{
		X: 20,
		Y: 20,
	}
	c := NewHorizontalContainer()
	c.SetBounds(0, 0, 100, 100)

	c.Visit(f)
	assert.Len(t, f.Stack, 1)
}

func TestLocatingVisitorReturnsContainerChildren(t *testing.T) {
	f := &LocatingFinder{
		X: 20,
		Y: 20,
	}
	c := NewHorizontalContainer()
	c.SetBounds(0, 0, 100, 100)
	c.Add(NewSpacer(50, 100))
	c.Add(NewSpacer(50, 100))
	c.Layout()

	s1 := c.Children()[0]

	c.Visit(f)
	assert.Len(t, f.Stack, 2)
	assert.Equal(t, c, f.Stack[0])
	assert.Equal(t, s1, f.Stack[1])
}

func TestPointInRect(t *testing.T) {
	data := []struct {
		x, y   int32
		rect   *sdl.Rect
		inside bool
	}{
		{0, 0, &sdl.Rect{0, 0, 10, 10}, true},
		{10, 10, &sdl.Rect{0, 0, 10, 10}, false},
		{-1, -1, &sdl.Rect{0, 0, 10, 10}, false},
		{9, 9, &sdl.Rect{0, 0, 10, 10}, true},
	}
	for _, d := range data {
		var message string
		if d.inside {
			message = fmt.Sprintf("Point %d/%d should be inside %v", d.x, d.y, d.rect)
		} else {
			message = fmt.Sprintf("Point %d/%d should not be inside %v", d.x, d.y, d.rect)
		}
		assert.Equal(t, d.inside, PointInRect(d.x, d.y, d.rect), message)
	}
}
