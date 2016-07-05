package ui

import "github.com/veandco/go-sdl2/sdl"

type HorizontalStackLayouter struct {
	x                        int32
	top, left, right, bottom int32
}

func (l *HorizontalStackLayouter) Layout(parent *sdl.Rect, c Widget) {
	l.x += l.left
	dimensions := c.Dimensions()
	c.ScreenDimensions().Y = parent.Y
	c.ScreenDimensions().X = parent.X + l.x
	c.ScreenDimensions().W = dimensions.W
	c.ScreenDimensions().H = dimensions.H

	l.x += c.Dimensions().W
	l.x += l.right
}

func (l *HorizontalStackLayouter) Reset(parent *sdl.Rect) {
	l.x = 0
}
