package ui

import "github.com/veandco/go-sdl2/sdl"

// Margin is a set of margins: top, right, bottom, and left
type Margin struct {
	Top, Right, Bottom, Left int32
}

// Empty returns true if all Margins are set to zero.
func (m Margin) Empty() bool {
	return m.Top == 0 || m.Right == 0 || m.Bottom == 0 || m.Left == 0
}

// Reduce takes a rect and returns a rect with the Margins applied.
func (m Margin) Reduce(input *sdl.Rect) *sdl.Rect {
	if m.Empty() {
		return input
	}

	return &sdl.Rect{
		X: input.X + m.Left,
		Y: input.Y + m.Top,
		W: input.W - (m.Left + m.Right),
		H: input.H - (m.Top + m.Bottom),
	}
}
