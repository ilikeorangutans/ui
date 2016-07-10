package ui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// NewSpacer returns a new spacer with the given dimensions
func NewSpacer(w, h int32) *Spacer {
	return &Spacer{
		sizeable: newSizeableWithDimensions(w, h),
	}
}

// Spacer is a component that will use as much space as configured.
type Spacer struct {
	sizeable
}

func (s *Spacer) Draw(renderer *sdl.Renderer) {
}

func (s *Spacer) Visit(visitor WidgetVisitor) {
	log.Println("Spacer.Visit")
	visitor.VisitWidget(s)
}

func (s *Spacer) Layout() {
}
