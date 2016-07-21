package ui

// WidgetStack is a stacked list of widgets.
type WidgetStack []Widget

// Notify gives each widget a chance to handle the given event, starting at the top most widget.
func (s WidgetStack) Notify(event *Event) {
	for i := range s {
		index := len(s) - 1 - i
		w := s[index]
		w.OnEvent(event)

		if event.Handled {
			break
		}
	}
}

// Top returns the widget at the top of the stack or nil if the stack is empty
func (s WidgetStack) Top() Widget {
	if len(s) == 0 {
		return nil
	}

	return s[len(s)-1]
}
