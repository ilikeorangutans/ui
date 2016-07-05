package ui

type HorizontalStackLayouter struct {
	x                        int32
	top, left, right, bottom int32
}

func (l *HorizontalStackLayouter) Layout(parent *Container) {
	desiredWidths := make([]int32, len(parent.Children()))

	for i := range parent.Children() {
		child := parent.Get(i)
		desiredWidths[i] = child.Dimensions().W
	}

	widths := distributeSizes(parent.Bounds().W, desiredWidths)

	x := parent.Bounds().X
	y := parent.Bounds().Y
	for i := range widths {
		child := parent.Get(i)
		heights := distributeSizes(parent.Bounds().H, []int32{child.Dimensions().H})
		child.SetBounds(x, y, widths[i], heights[0])
		x += widths[i]
	}
}

// distributeSizes attempts to equally distribute the sizes given based on the
// following assumptions:
// 1) If a width is given as 0, it's assumed to be a filler, i.e. equally
// distribute remaining space among all fillers.
// 2) if a fixed width > 0 is requested, attempt to keep that size, unless
// 3) all requested fixed widths exceed the available space, then they will be
// all equally resized to fit the available space.
func distributeSizes(available int32, sizes []int32) []int32 {
	var result []int32 = make([]int32, len(sizes))
	var remaining, required int32
	var fillers []int

	remaining = available

	for i := range sizes {
		s := sizes[i]
		result[i] = s
		remaining -= s
		required += s

		if s == 0 {
			fillers = append(fillers, i)
		}
	}

	if remaining < 0 {
		ratio := float32(available) / float32(required)
		for i := range result {
			result[i] = int32(float32(result[i]) * ratio)
		}
	}

	if len(fillers) > 0 && remaining > 0 {
		fillerWidth := remaining / int32(len(fillers))

		for _, i := range fillers {
			result[i] = fillerWidth
			remaining -= fillerWidth
		}
	}

	return result
}
