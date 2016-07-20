package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func NewLabel(text string, font *ttf.Font, color sdl.Color) *Label {
	return &Label{
		sizeable: sizeable{
			dimensions: &sdl.Rect{},
			bounds:     &sdl.Rect{},
			border:     EmptyBorder(),
			alignment:  Alignment{Top, Left},
		},
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
}

func (l *Label) SetText(text string) {
	if l.text == text {
		return
	}

	l.text = text
	l.updateTexture = true
}

func (l *Label) Layout() {
	l.sizeable.Layout()
	l.widgetArea = l.alignment.Fill(l.WidgetArea(), l.textDimensions)
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
		// TODO could this be deferred?
		l.renderTexture(renderer)
	}

	l.border.Draw(renderer)
	renderer.Copy(l.texture, nil, l.widgetArea)
}

func (l *Label) Visit(visitor WidgetVisitor) {
	visitor.VisitWidget(l)
}

func (l *Label) Destroy() {
	l.texture.Destroy()
}
