package ui

type VerticalStackLayouter struct {
	y                        int
	top, left, right, bottom int
}

func (l *VerticalStackLayouter) Layout(c *Container) {
	desiredHeights := make([]int32, len(c.Children()))

	for i := range c.Children() {
		child := c.Get(i)
		desiredHeights[i] = child.Dimensions().H
	}

	heights := distributeSizes(c.Bounds().H, desiredHeights)

	x := c.Bounds().X
	y := c.Bounds().Y
	for i := range heights {
		child := c.Get(i)
		widths := distributeSizes(c.Bounds().W, []int32{child.Dimensions().W})
		child.SetBounds(x, y, widths[0], heights[i])
		y += heights[i]
	}
}
