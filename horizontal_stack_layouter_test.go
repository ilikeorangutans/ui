package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestHorizontalStackLayouterChildFillsFullWidth(t *testing.T) {
	c := NewContainer()
	c.Dimensions().W = 800
	c.Bounds().W = 800
	w := NewSpacer(0, 0)
	c.Add(w)
	layouter := &HorizontalStackLayouter{}

	layouter.Layout(c)

	assert.Equal(t, int32(800), c.Bounds().W)
}

func TestHorizontalStackLayouterChildrenFillFullWidth(t *testing.T) {
	c := NewContainer()
	c.WidgetArea().W = 800
	c.WidgetArea().H = 200
	c.WidgetArea().Y = 100
	w := NewSpacer(0, 300)
	w2 := NewSpacer(0, 0)
	c.Add(w)
	c.Add(w2)
	layouter := &HorizontalStackLayouter{}

	layouter.Layout(c)

	assert.Equal(t, int32(0), w.Bounds().X)
	assert.Equal(t, int32(100), w.Bounds().Y)
	assert.Equal(t, int32(200), w.Bounds().H)

	assert.Equal(t, int32(400), w2.Bounds().X)
	assert.Equal(t, int32(100), w2.Bounds().Y)
	assert.Equal(t, int32(200), w2.Bounds().H)
}

func TestHorizontalStackLayouterWithFixedWithChild(t *testing.T) {
	c := NewContainer()
	c.Dimensions().W = 800
	c.WidgetArea().W = 800
	w := NewSpacer(200, 0)
	w2 := NewSpacer(0, 0)
	c.Add(w)
	c.Add(w2)
	layouter := &HorizontalStackLayouter{}

	layouter.Layout(c)

	assert.Equal(t, int32(200), w.Bounds().W)
	assert.Equal(t, int32(600), w2.Bounds().W)
}

func TestHorizontalStackLayouterSetsXYAccordingly(t *testing.T) {
	c := NewContainer()
	c.widgetArea = &sdl.Rect{10, 20, 800, 200}
	w1 := NewSpacer(0, 300)
	w2 := NewSpacer(0, 0)
	w3 := NewSpacer(0, 0)
	c.Add(w1)
	c.Add(w2)
	c.Add(w3)
	layouter := &HorizontalStackLayouter{}

	layouter.Layout(c)

	assert.Equal(t, int32(10), w1.Bounds().X)
	assert.Equal(t, int32(276), w2.Bounds().X)
	assert.Equal(t, int32(542), w3.Bounds().X)
}

func TestDistributeSizes(t *testing.T) {
	data := []struct {
		available       int32
		input, expected []int32
	}{
		{800, []int32{400}, []int32{400}},
		{800, []int32{0}, []int32{800}},
		{800, []int32{0, 0}, []int32{400, 400}},
		{800, []int32{200, 0}, []int32{200, 600}},
		{800, []int32{200, 200}, []int32{200, 200}},
		{800, []int32{0, 200}, []int32{600, 200}},
		{800, []int32{0, 0, 0}, []int32{266, 266, 266}},
		{800, []int32{500, 500}, []int32{400, 400}},
		{800, []int32{200, 700}, []int32{177, 622}},
		{800, []int32{900, 0}, []int32{800, 0}},
		{800, []int32{500, 500, 0}, []int32{400, 400, 0}},
	}

	for i := range data {
		d := data[i]

		res := distributeSizes(d.available, d.input)

		assert.Equal(t, d.expected, res)
	}

}
