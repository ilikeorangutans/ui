package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestNewBorderWrapsWidget(t *testing.T) {
	w := NewSpacer(100, 10)
	b := NewBorder(w, 0, sdl.Color{})

	assert.Equal(t, w, b.child)

	assert.Equal(t, int32(100), b.Dimensions().W)
}

func TestNewBorderCalculatesDimensionsCorrectly(t *testing.T) {
	w := NewSpacer(100, 10)
	b := NewBorder(w, 10, sdl.Color{})

	assert.Equal(t, w, b.child)

	assert.Equal(t, int32(120), b.Dimensions().W)
	assert.Equal(t, int32(30), b.Dimensions().H)
}

func TestNewBorderDoesNotOverrideFill(t *testing.T) {
	w := NewSpacer(0, 0)
	b := NewBorder(w, 10, sdl.Color{})

	assert.Equal(t, int32(0), b.Dimensions().W)
	assert.Equal(t, int32(0), b.Dimensions().H)
}

func TestNewBorderCalculatesBoundsCorrectly(t *testing.T) {
	w := NewSpacer(100, 10)
	b := NewBorder(w, 10, sdl.Color{})

	b.SetBounds(10, 20, 200, 50)

	assert.Equal(t, int32(10), b.Bounds().X)
	assert.Equal(t, int32(20), b.Bounds().Y)
	assert.Equal(t, int32(200), b.Bounds().W)
	assert.Equal(t, int32(50), b.Bounds().H)

	assert.Equal(t, int32(120), b.Dimensions().W)
	assert.Equal(t, int32(30), b.Dimensions().H)

	assert.Equal(t, int32(20), b.child.Bounds().X)
	assert.Equal(t, int32(30), b.child.Bounds().Y)
	assert.Equal(t, int32(180), b.child.Bounds().W)
	assert.Equal(t, int32(30), b.child.Bounds().H)
}
