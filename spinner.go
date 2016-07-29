package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

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

	plusButton := NewClickButton("+", font)
	plusButton.AddEventHandler(ButtonReleased, func(e *Event) bool {
		s.Increment()
		return true
	})
	minusButton := NewClickButton("-", font)
	minusButton.AddEventHandler(ButtonReleased, func(e *Event) bool {
		s.Decrement()
		return true
	})
	buttonContainer.Add(plusButton)
	buttonContainer.Add(minusButton)

	s.Layout()

	return s
}

type Spinner struct {
	Container
	Initial, Value, Max, Min, Step float64
	Format                         string
	label                          *Label
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
}

func (s *Spinner) Decrement() {
	if s.Value-s.Step < s.Min {
		return
	}
	s.Value -= s.Step
	s.updateLabel()
}

func (s *Spinner) formattedValue() string {
	format := s.Format
	if format == "" {
		format = "%d"
	}
	return fmt.Sprintf(format, s.Value)
}
