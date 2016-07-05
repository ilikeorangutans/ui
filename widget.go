package ui

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
	Draw(renderer *sdl.Renderer)
	SetBounds(x, y, w, h int32)
	Bounds() *sdl.Rect
	Dimensions() *sdl.Rect
}
