package ui

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
	Draw(renderer *sdl.Renderer)
	ScreenDimensions() *sdl.Rect
	Dimensions() *sdl.Rect
}
