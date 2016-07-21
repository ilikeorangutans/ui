package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type MouseButton uint8

const (
	_ MouseButton = iota
	LMB
	MMB
	RMB
)

func (s MouseButton) String() string {
	switch s {
	case LMB:
		return "LMB"
	case MMB:
		return "MMB"
	case RMB:
		return "RMB"
	default:
		return fmt.Sprintf("%dMB", s)
	}
}

type MouseButtonState uint8

const (
	ButtonUp MouseButtonState = iota
	ButtonDown
)

func NewMouseClickEventFromSdlEvent(e *sdl.MouseButtonEvent) *Event {
	return NewMouseClickEvent(e.Timestamp, MouseButton(e.Button), MouseButtonState(e.State), e.X, e.Y)
}

func NewMouseClickEvent(timestamp uint32, button MouseButton, state MouseButtonState, X, Y int32) *Event {
	return &Event{
		Timestamp: timestamp,
		Type:      "MouseClickEvent",
		Data: MouseClickEvent{
			Button: button,
			State:  state,
			Point: sdl.Point{
				X: X,
				Y: Y,
			},
		},
	}
}

type MouseOverEvent struct {
	sdl.Point
}

type MouseClickEvent struct {
	sdl.Point
	Button MouseButton
	State  MouseButtonState
}

func (m MouseClickEvent) String() string {
	return fmt.Sprintf("MouseClickEvent{X: %d, Y: %d, Button: %s, State: %d}", m.X, m.Y, m.Button, m.State)
}
