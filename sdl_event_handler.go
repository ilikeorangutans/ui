package ui

import "github.com/veandco/go-sdl2/sdl"

type SDLEventHandler struct {
	Root              Widget
	lastMousePosition sdl.Point
	lastHoveredWidget Widget
}

func (s *SDLEventHandler) Handle(sdlEvent sdl.Event) {
	switch t := sdlEvent.(type) {
	case *sdl.MouseButtonEvent:
		event := NewMouseClickEventFromSdlEvent(t)
		w := findComponentsUnder(s.Root, t.X, t.Y)
		w.Notify(event)
	case *sdl.MouseWheelEvent:

	case *sdl.MouseMotionEvent:
		s.handleMouseMotion(t)
	case *sdl.WindowEvent:
		switch t.Event {
		case sdl.WINDOWEVENT_RESIZED:
			s.Root.SetBounds(0, 0, t.Data1, t.Data2)
		case sdl.WINDOWEVENT_SIZE_CHANGED:
			s.Root.SetBounds(0, 0, t.Data1, t.Data2)
		}
	}
}

func (s *SDLEventHandler) handleMouseMotion(t *sdl.MouseMotionEvent) {
	stack := findComponentsUnder(s.Root, t.X, t.Y)
	defer func() {
		s.lastMousePosition.X = t.X
		s.lastMousePosition.Y = t.Y
	}()

	if len(stack) == 0 {
		return
	}

	w := stack.Top()
	if s.lastHoveredWidget != nil && w != s.lastHoveredWidget {
		event := &Event{
			Timestamp: t.Timestamp,
			Type:      "MouseOut",
		}
		s.lastHoveredWidget.OnEvent(event)
		s.lastHoveredWidget = nil
	}

	wasOutside := !pointInsideOfRect(s.lastMousePosition, w.Bounds())
	if wasOutside {
		event := &Event{
			Timestamp: t.Timestamp,
			Type:      "MouseOver",
			Data: MouseOverEvent{
				Point: sdl.Point{t.X, t.Y},
			},
		}
		s.lastHoveredWidget = w
		w.OnEvent(event)
	}
}

func findComponentsUnder(root Widget, x, y int32) WidgetStack {
	// TODO: this could probably be cached?
	f := &LocatingFinder{
		X: x,
		Y: y,
	}
	root.Visit(f)
	return f.Stack
}

func pointInsideOfRect(p sdl.Point, r *sdl.Rect) bool {
	return r.X <= p.X && p.X < r.X+r.W && r.Y <= p.Y && p.Y < r.Y+r.H
}
