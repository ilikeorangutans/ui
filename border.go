package ui

import "github.com/veandco/go-sdl2/sdl"

func EmptyBorder() *Border {
	return &Border{Margin: Margin{}}
}

func NewBorder(width int32, color sdl.Color) *Border {
	return &Border{
		Margin: Margin{width, width, width, width},
		color:  color,
		style:  FlatBorderStyle{},
	}
}

type Border struct {
	Margin
	bounds *sdl.Rect
	color  sdl.Color
	style  interface {
		Draw(*sdl.Renderer, *Border)
	}
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

	b.style.Draw(renderer, b)
}

// FlatBorderStyle draws a flat colored border
type FlatBorderStyle struct{}

func (s FlatBorderStyle) Draw(renderer *sdl.Renderer, b *Border) {
	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	if b.Top == 1 && b.Right == 1 && b.Bottom == 1 && b.Left == 1 {
		renderer.DrawRect(b.bounds)
	} else {
		var r []sdl.Rect
		r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y, W: b.bounds.W, H: b.Top})
		r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.Top, W: b.Left, H: b.bounds.H - (b.Top + b.Bottom)})
		r = append(r, sdl.Rect{X: b.bounds.X + b.bounds.W - b.Right, Y: b.bounds.Y + b.Top, W: b.Right, H: b.bounds.H - (b.Top + b.Bottom)})
		r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.bounds.H - b.Bottom, W: b.bounds.W, H: b.Bottom})
		renderer.FillRects(r)
	}
}

type LoweredBorderStyle struct{}

func (s LoweredBorderStyle) Draw(renderer *sdl.Renderer, b *Border) {
	var r []sdl.Rect

	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y, W: b.bounds.W, H: b.Top})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.Top, W: b.Left, H: b.bounds.H - (b.Top + b.Bottom)})
	renderer.SetDrawColor(108, 122, 137, 255)
	renderer.FillRects(r)

	r = nil
	r = append(r, sdl.Rect{X: b.bounds.X + b.bounds.W - b.Right, Y: b.bounds.Y + b.Top, W: b.Right, H: b.bounds.H - (b.Top + b.Bottom)})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.bounds.H - b.Bottom, W: b.bounds.W, H: b.Bottom})
	renderer.SetDrawColor(236, 240, 241, 255)
	renderer.FillRects(r)

	renderer.SetDrawColor(171, 183, 183, 255)
	renderer.FillRect(&sdl.Rect{X: b.bounds.X + b.Left, Y: b.bounds.Y + b.Top, W: b.bounds.W - b.Left - b.Right, H: b.bounds.H - b.Top - b.Bottom})
}

type RaisedBorderStyle struct{}

func (s RaisedBorderStyle) Draw(renderer *sdl.Renderer, b *Border) {
	var r []sdl.Rect

	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y, W: b.bounds.W, H: b.Top})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.Top, W: b.Left, H: b.bounds.H - (b.Top + b.Bottom)})
	renderer.SetDrawColor(236, 240, 241, 255)
	renderer.FillRects(r)

	r = nil
	r = append(r, sdl.Rect{X: b.bounds.X + b.bounds.W - b.Right, Y: b.bounds.Y + b.Top, W: b.Right, H: b.bounds.H - (b.Top + b.Bottom)})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.bounds.H - b.Bottom, W: b.bounds.W, H: b.Bottom})
	renderer.SetDrawColor(108, 122, 137, 255)
	renderer.FillRects(r)

	renderer.SetDrawColor(189, 195, 199, 255)
	renderer.FillRect(&sdl.Rect{X: b.bounds.X + b.Left, Y: b.bounds.Y + b.Top, W: b.bounds.W - b.Left - b.Right, H: b.bounds.H - b.Top - b.Bottom})
}
