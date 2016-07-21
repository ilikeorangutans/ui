package ui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestPointInsideOfRect(t *testing.T) {
	data := []struct {
		p      sdl.Point
		r      *sdl.Rect
		inside bool
	}{
		{sdl.Point{0, 0}, &sdl.Rect{0, 0, 10, 10}, true},
		{sdl.Point{9, 9}, &sdl.Rect{0, 0, 10, 10}, true},
		{sdl.Point{10, 10}, &sdl.Rect{0, 0, 10, 10}, false},
		{sdl.Point{-1, -1}, &sdl.Rect{0, 0, 10, 10}, false},
	}

	for _, d := range data {
		inside := pointInsideOfRect(d.p, d.r)
		assert.Equal(t, d.inside, inside, fmt.Sprintf("expected %v to be inside %v: %t", d.p, d.r, d.inside))
	}
}
