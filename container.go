package ui

import "github.com/veandco/go-sdl2/sdl"

func NewContainer() *Container {
	return &Container{
		sizeable: newSizeable(),
		layouter: &HorizontalStackLayouter{
			top: 3,
		},
	}
}

type Container struct {
	sizeable
	children []Widget
	layouter Layouter
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
