package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestNewLabelSetsUpdateTextureFlag(t *testing.T) {
	l := NewLabel("test", font, sdl.Color{255, 255, 255, 255})
	assert.True(t, l.updateTexture)
}

func TestLabelSetsUpdateTextureFlagAfterChangingText(t *testing.T) {
	l := NewLabel("test", font, sdl.Color{255, 255, 255, 255})
	l.updateTexture = false
	l.SetText("test")
	assert.False(t, l.updateTexture)
	l.SetText("different text")
	assert.True(t, l.updateTexture)
}

func TestLabelSetBoundsUpdatesTextDimensions(t *testing.T) {
	l := NewLabel("test", font, sdl.Color{255, 255, 255, 255})
	l.Dimensions().W = 300
	l.Dimensions().H = 400
	l.SetBounds(100, 200, 300, 400)
	l.SetAlignment(Alignment{FillVertical, FillHorizontal})
	l.Layout()

	assert.Equal(t, &sdl.Rect{0, 0, 300, 400}, l.Dimensions())
	assert.Equal(t, &sdl.Rect{100, 200, 300, 400}, l.BorderArea())
	assert.Equal(t, &sdl.Rect{100, 200, 300, 400}, l.WidgetArea())
}

func TestLabelWithMarginBorderAndPadding(t *testing.T) {
	l := NewLabel("test", font, sdl.Color{255, 255, 255, 255})
	l.Dimensions().W = 200
	l.Dimensions().H = 100
	l.SetBounds(11, 13, 200, 100)
	l.SetAlignment(Alignment{FillVertical, FillHorizontal})
	l.SetMargin(Margin{3, 3, 3, 3})
	l.SetBorder(NewFlatBorder(5, sdl.Color{}))
	l.SetPadding(Margin{7, 7, 7, 7})

	l.Layout()

	assert.Equal(t, &sdl.Rect{11, 13, 200, 100}, l.Bounds())
	assert.Equal(t, &sdl.Rect{11 + 3, 13 + 3, 194, 94}, l.BorderArea())
	assert.Equal(t, &sdl.Rect{11 + 3 + 5, 13 + 3 + 5, 184, 84}, l.PaddingArea())
	assert.Equal(t, &sdl.Rect{11 + 3 + 5 + 7, 13 + 3 + 5 + 7, 170, 70}, l.WidgetArea())
}

func TestLabelTextDimensionsWithAlignment(t *testing.T) {
	l := NewLabel("test", font, sdl.Color{255, 255, 255, 255})
	l.Dimensions().W = 200
	l.Dimensions().H = 100
	l.SetBounds(11, 13, 200, 100)
	l.SetAlignment(Alignment{Middle, Center})

	l.textDimensions.W = 100
	l.textDimensions.H = 50

	l.Layout()

	assert.Equal(t, &sdl.Rect{11, 13, 200, 100}, l.Bounds())
	assert.Equal(t, &sdl.Rect{11, 13, 200, 100}, l.BorderArea())
	assert.Equal(t, &sdl.Rect{11, 13, 200, 100}, l.PaddingArea())
	assert.Equal(t, &sdl.Rect{61, 38, 100, 50}, l.WidgetArea())
}
