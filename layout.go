package ui

import "github.com/veandco/go-sdl2/sdl"

func newSizeable() sizeable {
	return sizeable{
		dimensions: &sdl.Rect{},
		bounds:     &sdl.Rect{},
		border:     EmptyBorder(),
		drawArea:   &sdl.Rect{},
	}
}

func newSizeableWithDimensions(w, h int32) sizeable {
	return sizeable{
		dimensions: &sdl.Rect{
			W: w,
			H: h,
		},
		bounds: &sdl.Rect{
			W: w,
			H: h,
		},
		border:   EmptyBorder(),
		drawArea: &sdl.Rect{},
	}
}

// Dimensions returns the desired size.
func (s *sizeable) Dimensions() *sdl.Rect {
	return s.dimensions
}

// Bounds returns the size assigned by the layouter.
func (s *sizeable) Bounds() *sdl.Rect {
	return s.bounds
}

// DrawArea returns the area in which the widget is to be drawn, i.e.
// bounds, as assigned by the layouter, minus the border
func (s *sizeable) DrawArea() *sdl.Rect {
	return s.drawArea
}

func (s *sizeable) SetBorder(b *Border) {
	s.border = b
	s.Layout()
}

func (s *sizeable) Layout() {
	s.drawArea = s.border.Reduce(s.bounds)
}

func (s *sizeable) SetBounds(x, y, w, h int32) {
	s.bounds.X = x
	s.bounds.Y = y
	s.bounds.W = w
	s.bounds.H = h

	s.Layout()
}

type sizeable struct {
	border     *Border
	bounds     *sdl.Rect
	dimensions *sdl.Rect
	drawArea   *sdl.Rect
}

type Layouter interface {
	Layout(parent *Container)
}
