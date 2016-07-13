package ui

import "github.com/veandco/go-sdl2/sdl"

func newSizeable() sizeable {
	return sizeable{
		dimensions: &sdl.Rect{},
		bounds:     &sdl.Rect{},
		border:     Margin{},
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
		border:   Margin{},
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

func (s *sizeable) SetBorder(m Margin) {
	s.border = m
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
	border     Margin
	bounds     *sdl.Rect
	dimensions *sdl.Rect
	drawArea   *sdl.Rect
}

type Layouter interface {
	Layout(parent *Container)
}

// Margin is a set of margins: top, right, bottom, and left
type Margin struct {
	top, right, bottom, left int32
}

// Empty returns true if all Margins are set to zero.
func (m Margin) Empty() bool {
	return m.top == 0 || m.right == 0 || m.bottom == 0 || m.left == 0
}

// Reduce takes a rect and returns a rect with the Margins applied.
func (m Margin) Reduce(input *sdl.Rect) *sdl.Rect {
	if m.Empty() {
		return input
	}

	return &sdl.Rect{
		X: input.X + m.left,
		Y: input.Y + m.top,
		W: input.W - (m.left + m.right),
		H: input.H - (m.top + m.bottom),
	}
}
