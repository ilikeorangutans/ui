package ui

import "github.com/veandco/go-sdl2/sdl"

func NewBorder(width int32, color sdl.Color) *Border {
	return &Border{
		Margin: Margin{width, width, width, width},
		color:  color,
	}
}

type Border struct {
	Margin
	bounds *sdl.Rect
	color  sdl.Color
}

func (b *Border) SetBounds(bounds *sdl.Rect) {
	if b == nil {
		return
	}
	b.bounds = bounds
}

func (b *Border) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	renderer.DrawRect(b.bounds)
}
