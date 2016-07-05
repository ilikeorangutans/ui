package main

import (
	"fmt"
	"log"

	"github.com/ilikeorangutans/ui"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()

	font, err := ttf.OpenFont("/usr/share/fonts/truetype/ttf-bitstream-vera/Vera.ttf", 12)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

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

			}
		}

		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		c.Draw(renderer)

		renderer.Present()
		sdl.Delay(33)
	}

}