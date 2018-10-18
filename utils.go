package ui

import "github.com/veandco/go-sdl2/sdl"

func setDrawColor(renderer *sdl.Renderer, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func PointInRect(x, y int32, rect *sdl.Rect) bool {
	return rect.X <= x && x < rect.X+rect.W && rect.Y <= y && y < rect.Y+rect.H
}

func pointInsideOfRect(p sdl.Point, r *sdl.Rect) bool {
	return r.X <= p.X && p.X < r.X+r.W && r.Y <= p.Y && p.Y < r.Y+r.H
}

func findComponentsUnder(root Widget, x, y int32) WidgetStack {
	// TODO: this could probably be cached?
	f := &LocatingFinder{
		X: x,
		Y: y,
	}
	root.Visit(f)
	return f.Stack
}
