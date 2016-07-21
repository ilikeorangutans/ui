package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	ButtonPushed  = "ButtonPushed"
	ButtonClicked = "ButtonClicked"
)

// NewButton creates a new button with the given text as label.
func NewButton(text string, font *ttf.Font) *Button {
	b := &Button{
		Label: NewLabel(text, font, sdl.Color{}),
	}

	border := NewBorder(2, sdl.Color{})
	border.style = RaisedBorderStyle{}
	b.SetBorder(border)
	b.Label.SetAlignment(Alignment{Middle, Center})

	b.AddEventHandler("MouseClickEvent", b.OnMouseClick)
	b.AddEventHandler("MouseOver", b.OnMouseOver)
	b.AddEventHandler("MouseOut", b.OnMouseOut)

	return b
}

// A clickable button.
type Button struct {
	*Label
}

func (b *Button) Draw(renderer *sdl.Renderer) {
	b.border.Draw(renderer)
	b.Label.Draw(renderer)
}

func (b *Button) OnMouseClick(e *Event) bool {
	mouseClick := e.Data.(MouseClickEvent)
	if mouseClick.Button != LMB {
		return false
	}

	if mouseClick.State == ButtonDown {
		b.border.style = LoweredBorderStyle{}
	} else if mouseClick.State == ButtonUp {
		b.border.style = RaisedBorderStyle{}

		buttonEvent := &Event{
			Timestamp: e.Timestamp,
			Type:      ButtonClicked,
			Emitter:   b,
			Data:      ButtonClickEvent{},
		}
		b.OnEvent(buttonEvent)
	}
	return true
}

func (b *Button) OnMouseOver(e *Event) bool {
	return false
}

func (b *Button) OnMouseOut(e *Event) bool {
	return false
}

type ButtonClickEvent struct{}
