package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/ilikeorangutans/ui"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func loadFont() *ttf.Font {
	var fontPath string
	if runtime.GOOS == "darwin" {
		fontPath = "/Library/Fonts/Verdana.ttf"
	} else if runtime.GOOS == "linux" {
		fontPath = "/usr/share/fonts/truetype/ttf-bitstream-vera/Vera.ttf"
	}
	font, err := ttf.OpenFont(fontPath, 12)
	if err != nil {
		panic(err)
	}

	return font
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()
	ttf.Init()
	defer ttf.Quit()

	font := loadFont()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	label := ui.NewLabel("horray", font, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	label.Dimensions().W = 200
	label2 := ui.NewLabel("label two", font, sdl.Color{R: 255, G: 255, B: 0, A: 255})
	label3 := ui.NewLabel("label 3 this is awesome!!", font, sdl.Color{R: 255, G: 0, B: 255, A: 255})
	label3.Dimensions().W = 300

	c := ui.NewContainer()
	c.SetBounds(0, 0, 800, 200)
	c.Add(ui.NewBorder(label, 5, sdl.Color{G: 255, A: 255}))
	c.Add(ui.NewBorder(label2, 10, sdl.Color{R: 255, A: 255}))
	c.Add(label3)
	running := true
	var event sdl.Event
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
				log.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				s := fmt.Sprintf("sym:%c", t.Keysym.Sym)
				label.SetText(s)

			case *sdl.MouseButtonEvent:
				// state 1 = down, state 0 = up
				if t.State == 0 {
					log.Printf("[%d ms] mouse button %d state %d", t.Timestamp, t.Button, t.State)
					w := findComponentUnder(c, t.X, t.Y)
					log.Printf("[%d ms] found %d components", t.Timestamp, len(w))
					for _, w := range w {
						log.Printf("[%d ms] found component %v", t.Timestamp, w.Bounds())
					}

					v := &ui.LoggingVisitor{}
					c.Visit(v)
				}

			case *sdl.MouseWheelEvent:
				log.Printf("[%d ms] mouse wheel x %d y %d", t.Timestamp, t.X, t.Y)

			case *sdl.WindowEvent:
				switch t.Event {
				case sdl.WINDOWEVENT_RESIZED:
					c.SetBounds(0, 0, t.Data1, t.Data2)
				case sdl.WINDOWEVENT_SIZE_CHANGED:
					c.SetBounds(0, 0, t.Data1, t.Data2)
				}
			}
		}

		c.Layout()

		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		c.Draw(renderer)

		renderer.Present()
		sdl.Delay(33)
	}
}

func findComponentUnder(root *ui.Container, x, y int32) []ui.Widget {
	f := &ui.LocatingFinder{
		X: x,
		Y: y,
	}
	root.Visit(f)
	return f.Stack
}
