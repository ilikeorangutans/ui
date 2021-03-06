package ui

import "github.com/veandco/go-sdl2/sdl"

func NewHorizontalContainer() *Container {
	return &Container{
		BoxModel: newBoxModel(),
		layouter: &HorizontalStackLayouter{},
	}
}

func NewVerticalContainer() *Container {
	return &Container{
		BoxModel: newBoxModel(),
		layouter: &VerticalStackLayouter{},
	}
}

type Container struct {
	EventHandlers
	BoxModel
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

// Clear removes all children from a container and destroys them.
func (c *Container) Clear() {
	for i := range c.children {
		c.Get(i).Destroy()
		c.children[i] = nil
	}

	c.children = make([]Widget, 0)

}

func (c *Container) Layout() {
	c.BoxModel.Layout()
	c.layouter.Layout(c)
	for i := range c.children {
		c.children[i].Layout()
	}
}

func (c *Container) Draw(renderer *sdl.Renderer) {
	c.BoxModel.border.Draw(renderer)

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
	c.Clear()
}
