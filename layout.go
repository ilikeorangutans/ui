package ui

import "github.com/veandco/go-sdl2/sdl"

func newSizeable() sizeable {
	return sizeable{
		dimensions: &sdl.Rect{},
		bounds:     &sdl.Rect{},
		border:     EmptyBorder(),
		widgetArea: &sdl.Rect{},
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
		border:     EmptyBorder(),
		widgetArea: &sdl.Rect{},
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
// bounds, as assigned by the layouter, minus all margins
func (s *sizeable) WidgetArea() *sdl.Rect {
	return s.widgetArea
}

func (s *sizeable) SetMargin(m Margin) {
	s.margin = m
	s.layoutChanged = true
}

func (s *sizeable) SetPadding(m Margin) {
	s.padding = m
	s.layoutChanged = true
}

func (s *sizeable) PaddingArea() *sdl.Rect {
	return s.paddingArea
}

func (s *sizeable) SetBorder(b *Border) {
	s.border = b
	s.layoutChanged = true
}

func (s *sizeable) BorderArea() *sdl.Rect {
	return s.borderArea
}

func (s *sizeable) SetBounds(x, y, w, h int32) {
	s.bounds.X = x
	s.bounds.Y = y
	s.bounds.W = w
	s.bounds.H = h
	s.layoutChanged = true
}

func (s *sizeable) Layout() {
	if !s.layoutChanged {
		return
	}
	s.layoutChanged = false

	s.borderArea = s.margin.Reduce(s.bounds)
	s.border.SetBounds(s.borderArea)
	s.paddingArea = s.border.Reduce(s.borderArea)
	s.widgetArea = s.alignment.Fill(s.padding.Reduce(s.paddingArea), s.dimensions)
}

func (s *sizeable) SetAlignment(a Alignment) {
	s.alignment = a
}

func (s *sizeable) Destroy() {
}

// sizeable is the base type for layouting widgets. A sizeable widget consists of
// several layered elements. From the outside to the inside:
// - border: rendererd outside of everything
// - padding: blank space inside of the border
// - widget: the actual widget to be rendered
// For each of these elements there's a corresponding area that sizeable will
// keep up to date based on the bounds set on the widget:
// - BorderArea
// - PaddingArea
// - WidgetArea
type sizeable struct {
	alignment     Alignment
	border        *Border
	borderArea    *sdl.Rect
	padding       Margin
	paddingArea   *sdl.Rect
	margin        Margin
	bounds        *sdl.Rect
	dimensions    *sdl.Rect
	widgetArea    *sdl.Rect
	layoutChanged bool
}

type Layouter interface {
	Layout(parent *Container)
}
