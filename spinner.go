package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	SpinnerChanged EventType = "SpinnerChanged"
)

type SpinnerChangedEvent struct {
	NewValue float64
}

func NewSpinner(font *ttf.Font) *Spinner {
	s := &Spinner{
		Container: Container{
			sizeable: newSizeable(),
			layouter: &HorizontalStackLayouter{},
		},
		Step:   1.0,
		Format: "%9.0f",
		Max:    1000000,
	}

	s.SetDimensions(0, 34)

	label := NewLabel("", font, sdl.Color{255, 255, 255, 255})
	label.SetDimensions(0, 34)
	label.SetText(s.formattedValue())
	label.SetAlignment(Alignment{Middle, Right})
	label.SetPadding(Margin{3, 3, 3, 3})
	s.label = label
	s.Container.Add(label)

	buttonContainer := NewVerticalContainer()
	s.Container.Add(buttonContainer)
	buttonContainer.SetDimensions(17, 34)

	s.plusButton = NewClickButton("+", font)
	s.plusButton.AddEventHandler(ButtonReleased, func(e *Event) bool {
		s.Increment()
		return true
	})
	s.minusButton = NewClickButton("-", font)
	s.minusButton.AddEventHandler(ButtonReleased, func(e *Event) bool {
		s.Decrement()
		return true
	})
	buttonContainer.Add(s.plusButton)
	buttonContainer.Add(s.minusButton)

	s.Layout()

	return s
}

type Spinner struct {
	Container
	Initial, Value, Max, Min, Step float64
	Format                         string
	label                          *Label
	plusButton, minusButton        *Button
}

func (s *Spinner) SetValue(value float64) {
	s.Value = value
	s.updateLabel()
}

func (s *Spinner) updateLabel() {
	s.label.SetText(s.formattedValue())
}

func (s *Spinner) Increment() {
	if s.Value+s.Step > s.Max {
		return
	}
	s.Value += s.Step
	s.updateLabel()
	s.fireChangedEvent()
}

func (s *Spinner) Decrement() {
	if s.Value-s.Step < s.Min {
		return
	}
	s.Value -= s.Step
	s.updateLabel()
	s.fireChangedEvent()
}

func (s *Spinner) fireChangedEvent() {
	s.OnEvent(&Event{
		Timestamp: 0,
		Type:      SpinnerChanged,
		Emitter:   s,
		Data: SpinnerChangedEvent{
			NewValue: s.Value,
		},
	})
}

func (s *Spinner) formattedValue() string {
	format := s.Format
	if format == "" {
		format = "%d"
	}
	return fmt.Sprintf(format, s.Value)
}

func (s *Spinner) Destroy() {
	s.label.Destroy()
}
