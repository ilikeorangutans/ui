package ui

import "github.com/veandco/go-sdl2/sdl"

type Widget interface {
	Bounds() *sdl.Rect
	Dimensions() *sdl.Rect
	Draw(renderer *sdl.Renderer)
	Layout()
	OnEvent(event Event)
	SetBorder(border *Border)
	SetBounds(x, y, w, h int32)
	Visit(WidgetVisitor)
	Destroy()
}
