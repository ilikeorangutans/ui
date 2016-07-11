package ui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type WidgetVisitor interface {
	VisitContainer(*Container) []Widget
	VisitWidget(Widget)
}

type LoggingVisitor struct {
}

func (v *LoggingVisitor) VisitContainer(c *Container) []Widget {
	log.Printf("Visiting container %v\n", v)
	return c.Children()
}

func (v *LoggingVisitor) VisitWidget(w Widget) {
	log.Printf("Visiting widget %v\n", w)
}

type LocatingFinder struct {
	X, Y  int32
	Stack []Widget
}

func (v *LocatingFinder) VisitContainer(c *Container) []Widget {
	if PointInRect(v.X, v.Y, c.Bounds()) {
		v.Stack = append(v.Stack, c)
		return c.Children()
	}
	return []Widget{}
}

func (v *LocatingFinder) VisitWidget(w Widget) {
	if PointInRect(v.X, v.Y, w.Bounds()) {
		v.Stack = append(v.Stack, w)
	}
}

func PointInRect(x, y int32, rect *sdl.Rect) bool {
	return rect.X <= x && x < rect.X+rect.W && rect.Y <= y && y < rect.Y+rect.H
}
