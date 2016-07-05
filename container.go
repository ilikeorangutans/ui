package ui

import "github.com/veandco/go-sdl2/sdl"

func NewContainer() *Container {
	return &Container{
		dimensions: &sdl.Rect{},
		bounds:     &sdl.Rect{},
		layouter: &HorizontalStackLayouter{
			top: 3,
		},
	}
}

type Container struct {
	children   []Widget
	layouter   Layouter
	dimensions *sdl.Rect
	bounds     *sdl.Rect
}

func (c *Container) Dimensions() *sdl.Rect {
	return c.dimensions
}

func (c *Container) Bounds() *sdl.Rect {
	return c.bounds
}

func (c *Container) SetBounds(x, y, w, h int32) {
	c.bounds.X = x
	c.bounds.Y = y
	c.bounds.W = w
	c.bounds.H = h
}

func (c *Container) Get(n int) Widget {
	return c.children[n]
}

func (c *Container) Children() []Widget {
	return c.children
}

func (c *Container) Add(child Widget) {
	c.children = append(c.children, child)
}

func (c *Container) Draw(renderer *sdl.Renderer) {

	c.layouter.Layout(c)

	for i := range c.children {
		child := c.children[i]
		child.Draw(renderer)
	}

}
