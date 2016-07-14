package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestNewSizeableWithDimensions(t *testing.T) {
	s := newSizeableWithDimensions(200, 100)
	s.Layout()

	assert.Equal(t, int32(200), s.Dimensions().W)
	assert.Equal(t, int32(100), s.Dimensions().H)

	assert.Equal(t, int32(200), s.DrawArea().W)
	assert.Equal(t, int32(100), s.DrawArea().H)
}

func TestSizeableWithBorderHasSmallerDrawArea(t *testing.T) {
	s := newSizeableWithDimensions(200, 100)
	s.SetBorder(&Border{Margin: Margin{10, 20, 5, 15}})
	s.Layout()

	assert.Equal(t, int32(200), s.Dimensions().W)
	assert.Equal(t, int32(100), s.Dimensions().H)

	assert.Equal(t, int32(15), s.DrawArea().X)
	assert.Equal(t, int32(10), s.DrawArea().Y)
	assert.Equal(t, int32(165), s.DrawArea().W)
	assert.Equal(t, int32(85), s.DrawArea().H)
}

func TestMarginReduce(t *testing.T) {
	data := []struct {
		input  *sdl.Rect
		margin Margin
		output *sdl.Rect
	}{
		{&sdl.Rect{0, 0, 100, 100}, Margin{}, &sdl.Rect{0, 0, 100, 100}},
		{&sdl.Rect{0, 0, 100, 100}, Margin{10, 10, 10, 10}, &sdl.Rect{10, 10, 80, 80}},
	}

	for _, d := range data {
		output := d.margin.Reduce(d.input)

		assert.Equal(t, d.output, output)
	}
}
