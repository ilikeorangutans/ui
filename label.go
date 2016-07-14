package ui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func NewLabel(text string, font *ttf.Font, color sdl.Color) *Label {
	return &Label{
		sizeable:       newSizeable(),
		text:           text,
		font:           font,
		updateTexture:  true,
		color:          color,
		textDimensions: &sdl.Rect{},
		EventHandlers:  &EventHandlers{},
	}
}

type Label struct {
	sizeable
	*EventHandlers
	text           string
	font           *ttf.Font
	updateTexture  bool
	texture        *sdl.Texture
	color          sdl.Color
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

func (l *Label) SetBounds(x, y, w, h int32) {
	l.sizeable.SetBounds(x, y, w, h)
	l.textDimensions.X = l.drawArea.X
	l.textDimensions.Y = l.drawArea.Y
	l.border.SetBounds(l.bounds)
}

func (l *Label) renderTexture(renderer *sdl.Renderer) {
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
	l.textDimensions.W = surface.W
	l.textDimensions.H = surface.H
	l.updateTexture = false
}

func (l *Label) Draw(renderer *sdl.Renderer) {
	if l.updateTexture {
		l.renderTexture(renderer)
	}

	l.border.Draw(renderer)

	var boundsToUse *sdl.Rect
	if l.stretchToFill {
		boundsToUse = l.drawArea
	} else {
		boundsToUse = l.textDimensions
	}

	renderer.Copy(l.texture, nil, boundsToUse)
}

func (l *Label) OnMouseClick(e MouseClickEvent) {
	log.Printf("Label.OnMouseClick() \n")
}

func (l *Label) Visit(visitor WidgetVisitor) {
	visitor.VisitWidget(l)
}
