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

func (s *sizeable) layoutChanged() {
	s.layoutHasChanged = true
}

// Dimensions returns the desired size.
func (s *sizeable) Dimensions() *sdl.Rect {
	return s.dimensions
}

func (s *sizeable) SetDimensions(w, h int32) {
	s.dimensions.W = w
	s.dimensions.H = h
	s.layoutChanged()
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
	s.layoutChanged()
}

func (s *sizeable) SetPadding(m Margin) {
	s.padding = m
	s.layoutChanged()
}

func (s *sizeable) PaddingArea() *sdl.Rect {
	return s.paddingArea
}

func (s *sizeable) SetBorder(b *Border) {
	s.border = b
	s.layoutChanged()
}

func (s *sizeable) BorderArea() *sdl.Rect {
	return s.borderArea
}

func (s *sizeable) SetBounds(x, y, w, h int32) {
	s.bounds.X = x
	s.bounds.Y = y
	s.bounds.W = w
	s.bounds.H = h
	s.layoutChanged()
}

func (s *sizeable) Layout() {
	if !s.layoutHasChanged {
		return
	}

	hasOnlyDimensions := !s.dimensions.Empty() && s.bounds.Empty()
	if hasOnlyDimensions {
		s.bounds.W = s.dimensions.W
		s.bounds.H = s.dimensions.H
	}

	s.borderArea = s.margin.Reduce(s.bounds)
	s.border.SetBounds(s.borderArea)
	s.paddingArea = s.border.Reduce(s.borderArea)
	s.widgetArea = s.alignment.Fill(s.padding.Reduce(s.paddingArea), s.dimensions)

	s.layoutHasChanged = false
}

func (s *sizeable) SetAlignment(a Alignment) {
	s.alignment = a
	s.layoutChanged()
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
	alignment        Alignment
	border           *Border
	borderArea       *sdl.Rect
	padding          Margin
	paddingArea      *sdl.Rect
	margin           Margin
	bounds           *sdl.Rect
	dimensions       *sdl.Rect
	widgetArea       *sdl.Rect
	layoutHasChanged bool
}

type Layouter interface {
	Layout(parent *Container)
}
