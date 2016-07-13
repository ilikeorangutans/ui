package ui

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
	Bounds() *sdl.Rect
	Dimensions() *sdl.Rect
	Draw(renderer *sdl.Renderer)
	Layout()
	OnEvent(event Event)
	SetBounds(x, y, w, h int32)
	Visit(WidgetVisitor)
	SetBorder(border *Border)
}
