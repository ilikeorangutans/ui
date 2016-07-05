package ui

import "github.com/veandco/go-sdl2/sdl"

func NewContainer() *Container {
	return &Container{
		screenDimensions: &sdl.Rect{},
		layouter: &HorizontalStackLayouter{
			top: 3,
		},
	}
}

type Container struct {
	children         []Widget
	layouter         Layouter
	screenDimensions *sdl.Rect
}

func (c *Container) ScreenDimensions() *sdl.Rect {
	return c.screenDimensions
}

func (c *Container) Add(child *Label) {
	c.children = append(c.children, child)
}

func (c *Container) Draw(renderer *sdl.Renderer) {
	for i := range c.children {
		child := c.children[i]
		c.layouter.Layout(c.ScreenDimensions(), child)
		child.Draw(renderer)
	}

	c.layouter.Reset(c.screenDimensions)
}
