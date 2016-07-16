package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestAlignmentFillsParent(t *testing.T) {
	a := Alignment{}
	available := &sdl.Rect{11, 11, 100, 70}
	fill := &sdl.Rect{0, 0, 50, 50}

	result := a.Fill(available, fill)
	assert.Equal(t, available, result)
}

func TestAlignmentDealsWithNotEnoughSpace(t *testing.T) {
	a := Alignment{}
	available := &sdl.Rect{11, 11, 100, 70}
	fill := &sdl.Rect{0, 0, 50, 50}

	result := a.Fill(available, fill)
	assert.Equal(t, available, result)
}

func TestAlignmentTopLeft(t *testing.T) {
	a := Alignment{Horizontal: Left, Vertical: Top}
	available := &sdl.Rect{11, 11, 100, 50}
	fill := &sdl.Rect{0, 0, 50, 25}
	expected := &sdl.Rect{11, 11, 50, 25}

	result := a.Fill(available, fill)
	assert.Equal(t, expected, result)
}

func TestAlignmentBottomRight(t *testing.T) {
	a := Alignment{Horizontal: Right, Vertical: Bottom}
	available := &sdl.Rect{11, 11, 100, 50}
	fill := &sdl.Rect{0, 0, 50, 25}
	expected := &sdl.Rect{11 + 50, 11 + 25, 50, 25}

	result := a.Fill(available, fill)
	assert.Equal(t, expected, result)
}

func TestAlignmentMiddleCenter(t *testing.T) {
	a := Alignment{Horizontal: Center, Vertical: Middle}
	available := &sdl.Rect{11, 11, 100, 50}
	fill := &sdl.Rect{0, 0, 50, 25}
	expected := &sdl.Rect{11 + 25, 11 + 12, 50, 25}

	result := a.Fill(available, fill)
	assert.Equal(t, expected, result)
}
