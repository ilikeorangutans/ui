package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMouseClickEventDoesNothingWithoutHandlerMethod(t *testing.T) {
	w := NewSpacer(100, 100)

	event := MouseClickEvent{}
	event.Notify(w)
}

type WidgetWithClickHandler struct {
	Spacer
	clicked bool
}

func (w *WidgetWithClickHandler) OnMouseClick(e MouseClickEvent) {
	w.clicked = true
}

func TestMouseClickEventCallsHandler(t *testing.T) {
	w := &WidgetWithClickHandler{}

	event := MouseClickEvent{}
	event.Notify(w)

	assert.True(t, w.clicked)
}
