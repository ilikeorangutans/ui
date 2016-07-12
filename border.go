package ui

import "github.com/veandco/go-sdl2/sdl"

func NewBorder(w Widget, width int32, color sdl.Color) *Border {
	return &Border{
		sizeable:  newSizeable(),
		child:     w,
		thickness: width,
		color:     color,
	}
}

type Border struct {
	sizeable
	child     Widget
	thickness int32
	color     sdl.Color
}

func (b *Border) SetBounds(x, y, w, h int32) {
	b.sizeable.SetBounds(x, y, w, h)
	b.child.SetBounds(x+b.thickness, y+b.thickness, w-(2*b.thickness), h-(2*b.thickness))
}

func (b *Border) Dimensions() *sdl.Rect {
	c := b.child.Dimensions()
	var w int32
	if c.W > 0 {
		w = b.thickness + b.thickness + c.W
	}

	var h int32
	if c.H > 0 {
		h = b.thickness + b.thickness + c.H
	}
	dimensions := &sdl.Rect{
		X: c.X,
		Y: c.Y,
		W: w,
		H: h,
	}
	return dimensions
}

func (b *Border) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	renderer.DrawRect(b.bounds)
	b.child.Draw(renderer)
}

func (b *Border) Layout() {
	b.child.Layout()
}

func (b *Border) OnEvent(event Event) {}

func (b *Border) Visit(visitor WidgetVisitor) {
	// TODO: either we add a visitBorder or we have to represent borders differnetly
	visitor.VisitWidget(b.child)
}
