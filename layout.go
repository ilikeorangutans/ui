package ui

import "github.com/veandco/go-sdl2/sdl"

func newSizeable() sizeable {
	return sizeable{
		dimensions: &sdl.Rect{},
		bounds:     &sdl.Rect{},
	}
}

func newSizeableWithDimensions(w, h int32) sizeable {
	return sizeable{
		dimensions: &sdl.Rect{
			W: w,
			H: h,
		},
		bounds: &sdl.Rect{},
	}
}
func (s *sizeable) Dimensions() *sdl.Rect {
	return s.dimensions
}

func (s *sizeable) Bounds() *sdl.Rect {
	return s.bounds
}

func (s *sizeable) SetBounds(x, y, w, h int32) {
	s.bounds.X = x
	s.bounds.Y = y
	s.bounds.W = w
	s.bounds.H = h
}

type sizeable struct {
	dimensions *sdl.Rect
	bounds     *sdl.Rect
}

type Layouter interface {
	Layout(parent *Container)
}
