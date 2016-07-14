package ui

import "github.com/veandco/go-sdl2/sdl"

func EmptyBorder() *Border {
	return &Border{Margin: Margin{}}
}

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
	if b == nil || b.Empty() {
		return
	}

	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	if b.top == 1 && b.right == 1 && b.bottom == 1 && b.left == 1 {
		renderer.DrawRect(b.bounds)
	} else {
		var r []sdl.Rect

		r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y, W: b.bounds.W, H: b.top})
		r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.top, W: b.left, H: b.bounds.H - (b.top + b.bottom)})
		r = append(r, sdl.Rect{X: b.bounds.X + b.bounds.W - b.right, Y: b.bounds.Y + b.top, W: b.right, H: b.bounds.H - (b.top + b.bottom)})
		r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.bounds.H - b.bottom, W: b.bounds.W, H: b.bottom})
		renderer.FillRects(r)
	}

}
