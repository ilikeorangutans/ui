package ui

import "github.com/veandco/go-sdl2/sdl"

type HorizontalAlignment uint8

const (
	FillHorizontal HorizontalAlignment = iota
	Left
	Center
	Right
)

type VerticalAlignment uint8

const (
	FillVertical VerticalAlignment = iota
	Top
	Middle
	Bottom
)

// Alignment defines how content is laid out within its parent.
type Alignment struct {
	Horizontal HorizontalAlignment
	Vertical   VerticalAlignment
}

func (a Alignment) Fill(available *sdl.Rect, fill *sdl.Rect) *sdl.Rect {
	if a.Vertical == FillVertical && a.Horizontal == FillHorizontal {
		return available
	}

	var x, y, w, h int32
	switch a.Horizontal {
	case FillHorizontal:
		x = available.X
		w = available.W
	case Left:
		x = available.X
		w = fill.W
	case Center:
		x = available.X + ((available.W - fill.W) / 2)
		w = fill.W
	case Right:
		x = available.X + available.W - fill.W
		w = fill.W
	}

	switch a.Vertical {
	case FillVertical:
		y = available.Y
		h = available.H
	case Top:
		y = available.Y
		h = fill.H
	case Middle:
		y = available.Y + ((available.H - fill.H) / 2)
		h = fill.H
	case Bottom:
		y = available.Y + available.H - fill.H
		h = fill.H
	}

	result := &sdl.Rect{x, y, w, h}
	return result
}
