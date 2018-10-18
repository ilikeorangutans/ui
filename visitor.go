package ui

import (
	"log"
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

// LocatingFinder visits containers until it has found a leaf of the widget graph that contains the given coordinates.
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
