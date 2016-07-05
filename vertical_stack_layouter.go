package ui

import "github.com/veandco/go-sdl2/sdl"

type VerticalStackLayouter struct {
	y                        int
	top, left, right, bottom int
}

func (l *VerticalStackLayouter) Layout(parent *sdl.Rect, c Widget) {
	l.y += l.top
	dimensions := c.Dimensions()
	c.ScreenDimensions().Y = int32(l.y)
	c.ScreenDimensions().X = parent.X
	c.ScreenDimensions().W = dimensions.W
	c.ScreenDimensions().H = dimensions.H

	l.y += int(c.Dimensions().H)
	l.y += l.bottom
}

func (l *VerticalStackLayouter) Reset() {
	l.y = 0
}
