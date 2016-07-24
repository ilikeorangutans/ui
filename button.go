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
func NewClickButton(text string, font *ttf.Font) *Button {
	b := newButton(text, font)

	states := make(map[string]ButtonState)
	states["default"] = &ClickButtonDefaultState{EmptyButtonState: EmptyButtonState{Button: b}}
	states["hover"] = &ClickButtonHoverState{EmptyButtonState: EmptyButtonState{Button: b}}
	states["push"] = &ClickButtonPushState{EmptyButtonState: EmptyButtonState{Button: b}}
	states["click"] = &ClickButtonClickState{EmptyButtonState: EmptyButtonState{Button: b}}
	b.states = states

	b.transition("default")

	return b
}

func newButton(text string, font *ttf.Font) *Button {
	b := &Button{
		Label: NewLabel(text, font, sdl.Color{}),
	}

	border := NewBorder(2, sdl.Color{})
	border.Style = RaisedBorderStyle{}
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
	state  ButtonState
	states map[string]ButtonState
}

func (b *Button) transition(name string) {
	if b.state != nil {
		b.state.End()
	}
	b.state = b.states[name]
	b.state.Begin()
}

func (b *Button) Draw(renderer *sdl.Renderer) {
	b.state.Tick()
	b.border.Draw(renderer)
	b.Label.Draw(renderer)
}

func (b *Button) OnMouseClick(e *Event) bool {
	return b.state.OnMouseClick(e)
}

func (b *Button) OnMouseOver(e *Event) bool {
	return b.state.OnMouseOver(e)
}

func (b *Button) OnMouseOut(e *Event) bool {
	return b.state.OnMouseOut(e)
}

type ButtonClickEvent struct{}
