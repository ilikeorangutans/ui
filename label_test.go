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

	l.SetBounds(100, 200, 300, 400)

	assert.Equal(t, int32(100), l.textDimensions.X)
	assert.Equal(t, int32(200), l.textDimensions.Y)
}
