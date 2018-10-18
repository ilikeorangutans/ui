package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// HorizontalAlignment describes how a widget will be laid out horizontally within its assigned bounds.
type HorizontalAlignment uint8

func (a HorizontalAlignment) String() string {
	switch a {
	case FillHorizontal:
		return "Fill"
	case Left:
		return "Left"
	case Center:
		return "Center"
	case Right:
		return "Right"
	default:
		panic("Unknown alignment")
	}
}

const (
	FillHorizontal HorizontalAlignment = iota
	Left
	Center
	Right
)

// VerticalAlignment describes how a widget will be laid out vertically within its assigned bounds.
type VerticalAlignment uint8

func (a VerticalAlignment) String() string {
	switch a {
	case FillVertical:
		return "Fill"
	case Top:
		return "Top"
	case Middle:
		return "Middle"
	case Bottom:
		return "Bottom"
	default:
		panic("Unknown alignment")
	}
}

const (
	FillVertical VerticalAlignment = iota
	Top
	Middle
	Bottom
)

// Alignment defines how content is laid out within its parent.
type Alignment struct {
	Vertical   VerticalAlignment
	Horizontal HorizontalAlignment
}

// Fill fits the given Rect into the available Rect and returns a Rect according to its settings.
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

	return &sdl.Rect{x, y, w, h}
}

func (a Alignment) String() string {
	return fmt.Sprintf("Alignment{%s|%s}", a.Vertical, a.Horizontal)
}
