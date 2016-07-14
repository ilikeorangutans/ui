package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func NewButton(text string, font *ttf.Font) *Button {
	b := &Button{
		Label: NewLabel(text, font, sdl.Color{}),
	}

	border := NewBorder(2, sdl.Color{})
	border.style = RaisedBorderStyle{}
	b.SetBorder(border)

	b.AddEventHandler("MouseClickEvent", b.OnMouseClick)

	return b
}

type Button struct {
	*Label
}

func (b *Button) Draw(renderer *sdl.Renderer) {
	b.border.Draw(renderer)
	b.Label.Draw(renderer)
}

func (b *Button) OnMouseClick(e Event) {
	mouseClick := e.Data.(MouseClickEvent)
	if mouseClick.Button != LMB {
		return
	}

	if mouseClick.State == ButtonDown {
		b.border.style = LoweredBorderStyle{}
	} else if mouseClick.State == ButtonUp {
		b.border.style = RaisedBorderStyle{}
	}
}
