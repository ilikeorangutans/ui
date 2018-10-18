package ui

import "github.com/veandco/go-sdl2/sdl"

func EmptyBorderStyle(renderer *sdl.Renderer, b *Border) {
	// Do nothing
}

func FlatBorderStyle(renderer *sdl.Renderer, b *Border) {
	setDrawColor(renderer, b.color)
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

func LoweredBorderStyle(renderer *sdl.Renderer, b *Border) {
	var r []sdl.Rect

	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y, W: b.bounds.W, H: b.Top})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.Top, W: b.Left, H: b.bounds.H - (b.Top + b.Bottom)})
	setDrawColor(renderer, borderDarkColor)
	renderer.FillRects(r)

	r = nil
	r = append(r, sdl.Rect{X: b.bounds.X + b.bounds.W - b.Right, Y: b.bounds.Y + b.Top, W: b.Right, H: b.bounds.H - (b.Top + b.Bottom)})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.bounds.H - b.Bottom, W: b.bounds.W, H: b.Bottom})
	setDrawColor(renderer, borderLightColor)
	renderer.FillRects(r)

	setDrawColor(renderer, borderLoweredFill)
	renderer.FillRect(&sdl.Rect{X: b.bounds.X + b.Left, Y: b.bounds.Y + b.Top, W: b.bounds.W - b.Left - b.Right, H: b.bounds.H - b.Top - b.Bottom})
}

func RaisedBorderStyle(renderer *sdl.Renderer, b *Border) {
	var r []sdl.Rect

	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y, W: b.bounds.W, H: b.Top})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.Top, W: b.Left, H: b.bounds.H - (b.Top + b.Bottom)})
	setDrawColor(renderer, borderLightColor)
	renderer.FillRects(r)

	r = nil
	r = append(r, sdl.Rect{X: b.bounds.X + b.bounds.W - b.Right, Y: b.bounds.Y + b.Top, W: b.Right, H: b.bounds.H - (b.Top + b.Bottom)})
	r = append(r, sdl.Rect{X: b.bounds.X, Y: b.bounds.Y + b.bounds.H - b.Bottom, W: b.bounds.W, H: b.Bottom})
	setDrawColor(renderer, borderDarkColor)
	renderer.FillRects(r)

	setDrawColor(renderer, borderRaisedFill)
	renderer.FillRect(&sdl.Rect{X: b.bounds.X + b.Left, Y: b.bounds.Y + b.Top, W: b.bounds.W - b.Left - b.Right, H: b.bounds.H - b.Top - b.Bottom})
}
