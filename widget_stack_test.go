package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyStopsWhenEventIsHandled(t *testing.T) {
	handled := false
	ws := WidgetStack{}
	topOfStack := &TestWidget{}
	topOfStack.AddEventHandler("test", func(event *Event) bool {
		handled = true
		return true
	})
	bottomOfStack := &TestWidget{}
	bottomOfStack.AddEventHandler("test", func(event *Event) bool {
		assert.Fail(t, "this handler should not be called")
		return false
	})

	ws = append(ws, bottomOfStack)
	ws = append(ws, topOfStack)

	ws.Notify(&Event{Type: "test"})
}
