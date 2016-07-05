package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func NewLabel(text string, font *ttf.Font, color sdl.Color) *Label {
	return &Label{
		text:              text,
		font:              font,
		updateTexture:     true,
		dimensions:        &sdl.Rect{},
		color:             color,
		screenCoordinates: &sdl.Rect{},
	}
}

type Label struct {
	text              string
	font              *ttf.Font
	updateTexture     bool
	texture           *sdl.Texture
	color             sdl.Color
	dimensions        *sdl.Rect
	screenCoordinates *sdl.Rect
}

func (l *Label) SetText(text string) {
	if l.text == text {
		return
	}

	l.text = text
	l.updateTexture = true
}

func (l *Label) Dimensions() *sdl.Rect {
	return l.dimensions
}

func (l *Label) ScreenDimensions() *sdl.Rect {
	return l.screenCoordinates
}

func (l *Label) Draw(renderer *sdl.Renderer) {
	if l.updateTexture {
		surface, err := l.font.RenderUTF8_Solid(l.text, l.color)
		if err != nil {
			panic(err)
		}
		defer surface.Free()

		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			panic(err)
		}
		l.texture = texture
		l.dimensions.W = surface.W
		l.dimensions.H = surface.H
		l.updateTexture = false
	}

	renderer.Copy(l.texture, nil, l.screenCoordinates)
}
