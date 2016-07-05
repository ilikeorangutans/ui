package ui

import "github.com/veandco/go-sdl2/sdl"

// NewSpacer returns a new spacer with the given dimensions
func NewSpacer(w, h int32) *Spacer {
	return &Spacer{
		bounds:     &sdl.Rect{},
		dimensions: &sdl.Rect{W: w, H: h},
	}
}

// Spacer is a component that will use as much space as configured.
type Spacer struct {
	bounds     *sdl.Rect
	dimensions *sdl.Rect
}

func (s *Spacer) SetBounds(x, y, w, h int32) {
	s.bounds.X = x
	s.bounds.Y = y
	s.bounds.W = w
	s.bounds.H = h
}

func (s *Spacer) Draw(renderer *sdl.Renderer) {
}

func (s *Spacer) Bounds() *sdl.Rect {
	return s.bounds
}

func (s *Spacer) Dimensions() *sdl.Rect {
	return s.dimensions
}
