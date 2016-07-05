package ui

import "github.com/veandco/go-sdl2/sdl"

type Layouter interface {
	Layout(parent *sdl.Rect, w Widget)
	Reset(parent *sdl.Rect)
}
