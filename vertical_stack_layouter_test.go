package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestVerticalStackLayouter(t *testing.T) {
	layouter := &VerticalStackLayouter{}

	c := NewContainer()
	c.SetBounds(0, 0, 200, 300)
	s1 := c.Add(NewSpacer(100, 100))
	s2 := c.Add(NewSpacer(0, 100))

	layouter.Layout(c)

	assert.Equal(t, &sdl.Rect{X: 0, Y: 0, W: 100, H: 100}, s1.Bounds())
	assert.Equal(t, &sdl.Rect{X: 0, Y: 100, W: 200, H: 100}, s2.Bounds())
}
