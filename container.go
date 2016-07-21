package ui

import "github.com/veandco/go-sdl2/sdl"

func NewContainer() *Container {
	return &Container{
		sizeable: newSizeable(),
		layouter: &HorizontalStackLayouter{},
	}
}

func NewVerticalContainer() *Container {
	return &Container{
		sizeable: newSizeable(),
		layouter: &VerticalStackLayouter{},
	}
}

type Container struct {
	EventHandlers
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

func (c *Container) Add(child Widget) Widget {
	c.children = append(c.children, child)
	return child
}

func (c *Container) Layout() {
	c.sizeable.Layout()
	c.layouter.Layout(c)
	for i := range c.children {
		c.children[i].Layout()
	}
}

func (c *Container) Draw(renderer *sdl.Renderer) {
	c.sizeable.border.Draw(renderer)

	for i := range c.children {
		child := c.children[i]
		child.Draw(renderer)
	}
}

func (c *Container) Visit(visitor WidgetVisitor) {
	visitNext := visitor.VisitContainer(c)
	for i := range visitNext {
		visitNext[i].Visit(visitor)
	}
}

func (c *Container) Destroy() {
	for i := range c.children {
		c.children[i].Destroy()
	}
}
