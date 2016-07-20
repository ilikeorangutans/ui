package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestNewSizeableWithDimensions(t *testing.T) {
	s := newSizeableWithDimensions(200, 100)
	s.SetBounds(0, 0, 200, 100)
	s.Layout()

	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, s.Dimensions())
	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, s.BorderArea())
	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, s.PaddingArea())
	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, s.WidgetArea())
}

func TestSizeableLayoutWithMarginAndBorderAndPadding(t *testing.T) {
	s := newSizeableWithDimensions(200, 100)
	s.SetBounds(0, 0, 200, 100)
	s.SetMargin(Margin{3, 3, 3, 3})
	s.SetBorder(NewBorder(13, sdl.Color{}))
	s.SetPadding(Margin{11, 11, 11, 11})
	s.Layout()

	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, s.Dimensions())
	assert.Equal(t, &sdl.Rect{3, 3, 194, 94}, s.BorderArea())
	assert.Equal(t, &sdl.Rect{16, 16, 168, 68}, s.PaddingArea())
	assert.Equal(t, &sdl.Rect{27, 27, 146, 46}, s.WidgetArea())
}

func TestAlignmentDoesNothingIfWidgetFillsSpace(t *testing.T) {
	s := newSizeableWithDimensions(200, 100)
	s.SetBounds(0, 0, 100, 50)
	s.Layout()

	assert.Equal(t, &sdl.Rect{0, 0, 200, 100}, s.Dimensions())
	assert.Equal(t, &sdl.Rect{0, 0, 100, 50}, s.WidgetArea())
}

func TestSizeableWithBorderHasSmallerDrawArea(t *testing.T) {
	s := newSizeableWithDimensions(200, 100)
	s.SetBorder(&Border{Margin: Margin{10, 20, 5, 15}})
	s.Layout()

	assert.Equal(t, int32(200), s.Dimensions().W)
	assert.Equal(t, int32(100), s.Dimensions().H)

	assert.Equal(t, int32(15), s.WidgetArea().X)
	assert.Equal(t, int32(10), s.WidgetArea().Y)
	assert.Equal(t, int32(165), s.WidgetArea().W)
	assert.Equal(t, int32(85), s.WidgetArea().H)
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
