package ui

import "time"

// ButtonState represents the different states a button can be in.
type ButtonState interface {
	Begin()
	End()
	Tick()
	OnMouseClick(e *Event) bool
	OnMouseOver(e *Event) bool
	OnMouseOut(e *Event) bool
}

// EmptyButtonState is a skeleton for button states. It doesn't handle any events or does anything when transitioned to.
type EmptyButtonState struct {
	Button *Button
}

func (s *EmptyButtonState) Begin()                     {}
func (s *EmptyButtonState) End()                       {}
func (s *EmptyButtonState) Tick()                      {}
func (s *EmptyButtonState) OnMouseClick(e *Event) bool { return false }
func (s *EmptyButtonState) OnMouseOver(e *Event) bool  { return false }
func (s *EmptyButtonState) OnMouseOut(e *Event) bool   { return false }

// ClickButtonDefaultStyle is the default state for click buttons, i.e. raised border style.
type ClickButtonDefaultState struct {
	EmptyButtonState
}

func (s *ClickButtonDefaultState) Begin() {
	s.Button.border.Style = RaisedBorderStyle{}
}

func (s *ClickButtonDefaultState) OnMouseClick(e *Event) bool {
	data := e.Data.(MouseClickEvent)
	if data.Button != LMB || data.State != ButtonDown {
		return false
	}

	s.Button.transition("push")
	return true
}

func (s *ClickButtonDefaultState) OnMouseOver(e *Event) bool {
	s.Button.transition("hover")
	return false
}

func (s *ClickButtonDefaultState) String() string {
	return "default"
}

// ClickButtonHoverState is the click button state when the mouse cursor hovers
type ClickButtonHoverState struct {
	EmptyButtonState
}

func (s *ClickButtonHoverState) Begin() {
	// TODO set a highlight state
	s.Button.border.Style = RaisedBorderStyle{}
}

func (s *ClickButtonHoverState) String() string {
	return "hover"
}

func (s *ClickButtonHoverState) OnMouseOut(e *Event) bool {
	s.Button.transition("default")
	return false
}

func (s *ClickButtonHoverState) OnMouseClick(e *Event) bool {
	data := e.Data.(MouseClickEvent)
	if data.Button != LMB || data.State != ButtonDown {
		return false
	}

	s.Button.transition("push")
	return true
}

type ClickButtonPushState struct {
	EmptyButtonState
}

func (s *ClickButtonPushState) String() string {
	return "push"
}

func (s *ClickButtonPushState) Begin() {
	// TODO fire push event
	// TODO how to deal with autorepeat?
	s.Button.border.Style = &LoweredBorderStyle{}
}

func (s *ClickButtonPushState) OnMouseClick(e *Event) bool {
	data := e.Data.(MouseClickEvent)

	if data.Button != LMB {
		return false
	}

	if data.Button == LMB && data.State == ButtonUp {
		s.Button.transition("click")
		return true
	}

	s.Button.transition("hover")
	return true
}

func (s *ClickButtonPushState) OnMouseOut(e *Event) bool {
	s.Button.transition("default")
	return false
}

// ClickButtonClickState is the short transient state a button is in when it has been clicked.
type ClickButtonClickState struct {
	EmptyButtonState
}

func (s *ClickButtonClickState) String() string {
	return "click"
}

func (s *ClickButtonClickState) Begin() {
	// TODO set highlight active state
	event := &Event{
		Timestamp: uint32(time.Now().UnixNano()),
		Type:      ButtonReleased,
		Emitter:   s.Button,
		Data:      ButtonClickEvent{},
	}
	s.Button.OnEvent(event)
}

func (s *ClickButtonClickState) Tick() {
	s.Button.transition("hover")
}

type ToggleButtonDefaultState struct {
	EmptyButtonState
}

func (s *ToggleButtonDefaultState) Begin() {
	s.Button.border.Style = RaisedBorderStyle{}
}

func (s *ToggleButtonDefaultState) OnMouseOver(e *Event) bool {
	s.Button.transition("hover")
	return false
}

type ToggleButtonHoverState struct {
	EmptyButtonState
}

func (s *ToggleButtonHoverState) Begin() {
	s.Button.border.Style = RaisedBorderStyle{}
}

func (s *ToggleButtonHoverState) OnMouseOut(e *Event) bool {
	s.Button.transition("default")
	return false
}

func (s *ToggleButtonHoverState) OnMouseClick(e *Event) bool {
	data := e.Data.(MouseClickEvent)
	if data.Button != LMB || data.State != ButtonDown {
		return false
	}

	s.Button.OnEvent(&Event{
		Timestamp: e.Timestamp,
		Type:      ButtonToggled,
		Emitter:   s.Button,
		Data: ButtonToggleEvent{
			Pushed: true,
		},
	})
	s.Button.transition("pushhover")
	return true
}

//
type ToggleButtonPushedState struct {
	EmptyButtonState
}

func (s *ToggleButtonPushedState) Begin() {
	s.Button.border.Style = &LoweredBorderStyle{}
}

func (s *ToggleButtonPushedState) OnMouseOver(e *Event) bool {
	s.Button.transition("pushhover")
	return false
}

//
type ToggleButtonPushedHoverState struct {
	EmptyButtonState
}

func (s *ToggleButtonPushedHoverState) Begin() {
	s.Button.border.Style = &LoweredBorderStyle{}
}

func (s *ToggleButtonPushedHoverState) OnMouseOut(e *Event) bool {
	s.Button.transition("push")
	return false
}

func (s *ToggleButtonPushedHoverState) OnMouseClick(e *Event) bool {
	data := e.Data.(MouseClickEvent)
	if data.Button != LMB || data.State != ButtonDown {
		return false
	}

	s.Button.OnEvent(&Event{
		Timestamp: e.Timestamp,
		Type:      ButtonToggled,
		Emitter:   s.Button,
		Data: ButtonToggleEvent{
			Pushed: false,
		},
	})
	s.Button.transition("hover")
	return true
}
