package ui

import "github.com/veandco/go-sdl2/sdl"

var borderLightColor = sdl.Color{236, 240, 241, 255}
var borderDarkColor = sdl.Color{108, 122, 137, 255}
var borderRaisedFill = sdl.Color{189, 195, 199, 255}
var borderLoweredFill = sdl.Color{171, 183, 183, 255}


func EmptyBorder() *Border {
	return &Border{Margin: Margin{}}
}

func NewBorder(width int32, color sdl.Color, style BorderStyle) *Border {
	return &Border{
		Margin:     Margin{width, width, width, width},
		drawBorder: style,
		color:      color,
	}
}

func NewFlatBorder(width int32, color sdl.Color) *Border {
	return NewBorder(width, color, FlatBorderStyle)
}

type Border struct {
	Margin
	drawBorder BorderStyle
	bounds     *sdl.Rect
	color      sdl.Color
}

func (b *Border) SetBounds(bounds *sdl.Rect) {
	if b == nil {
		return
	}
	b.bounds = bounds
}

func (b *Border) SetStyle(style BorderStyle) {
	b.drawBorder = style
}

func (b *Border) Draw(renderer *sdl.Renderer) {
	if b == nil || b.Empty() {
		return
	}

	b.drawBorder(renderer, b)
}

// BorderStyle is a function that draws the given border
type BorderStyle func(renderer *sdl.Renderer, border *Border)

