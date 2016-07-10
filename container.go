package ui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

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

func (c *Container) Layout() {
	c.layouter.Layout(c)
	for i := range c.children {
		c.children[i].Layout()
	}
}

func (c *Container) Draw(renderer *sdl.Renderer) {
	for i := range c.children {
		child := c.children[i]
		child.Draw(renderer)
	}
}

func (c *Container) Visit(visitor WidgetVisitor) {
	log.Printf("Container.Visit")
	visitNext := visitor.VisitContainer(c)
	for i := range visitNext {
		visitNext[i].Visit(visitor)
	}
}
