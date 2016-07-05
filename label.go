package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func NewLabel(text string, font *ttf.Font, color sdl.Color) *Label {
	return &Label{
		text:           text,
		font:           font,
		updateTexture:  true,
		dimensions:     &sdl.Rect{},
		color:          color,
		bounds:         &sdl.Rect{},
		textDimensions: &sdl.Rect{},
	}
}

type Label struct {
	text           string
	font           *ttf.Font
	updateTexture  bool
	texture        *sdl.Texture
	color          sdl.Color
	dimensions     *sdl.Rect
	bounds         *sdl.Rect
	textDimensions *sdl.Rect
	stretchToFill  bool
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

func (l *Label) Bounds() *sdl.Rect {
	return l.bounds
}

func (l *Label) SetBounds(x, y, w, h int32) {
	l.bounds.X = x
	l.bounds.Y = y
	l.bounds.W = w
	l.bounds.H = h
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
		l.textDimensions.X = l.bounds.X
		l.textDimensions.Y = l.bounds.Y
		l.textDimensions.W = surface.W
		l.textDimensions.H = surface.H
		l.updateTexture = false
	}

	var boundsToUse *sdl.Rect
	if l.stretchToFill {
		boundsToUse = l.bounds
	} else {
		boundsToUse = l.textDimensions
	}

	renderer.Copy(l.texture, nil, boundsToUse)
}
