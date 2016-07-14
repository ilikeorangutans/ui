package ui

import "github.com/veandco/go-sdl2/sdl"

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
