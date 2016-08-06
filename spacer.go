package ui

import "github.com/veandco/go-sdl2/sdl"

// NewSpacer returns a new spacer with the given dimensions
func NewSpacer(w, h int32) *Spacer {
	return &Spacer{
		BoxModel: newBoxModelWithDimensions(w, h),
	}
}

// Spacer is a component that will use as much space as configured.
type Spacer struct {
	BoxModel
	EventHandlers
}

func (s *Spacer) Draw(renderer *sdl.Renderer) {
	s.border.Draw(renderer)
}

func (s *Spacer) Visit(visitor WidgetVisitor) {
	visitor.VisitWidget(s)
}

func (s *Spacer) Layout() {
}
