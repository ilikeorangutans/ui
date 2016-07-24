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

	counter := 0

	label := ui.NewLabel("horray", font, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	label.SetBorder(ui.NewBorder(10, sdl.Color{G: 255, A: 255}))
	label.Dimensions().W = 200
	label.SetAlignment(ui.Alignment{ui.Middle, ui.Center})

	label2 := ui.NewLabel("label two", font, sdl.Color{R: 255, G: 255, B: 0, A: 255})
	label2.SetBorder(ui.NewBorder(35, sdl.Color{R: 255, A: 255}))
	label3 := ui.NewLabel("label 3 this is awesome!!", font, sdl.Color{R: 255, G: 0, B: 255, A: 255})
	label3.SetBorder(ui.NewBorder(1, sdl.Color{R: 80, B: 255, A: 255}))
	label3.Dimensions().W = 300
	label3.SetAlignment(ui.Alignment{ui.Bottom, ui.Right})

	label3.AddEventHandler("MouseOver", func(e *ui.Event) bool {
		log.Printf("Got mouse over event")
		return true
	})
	label3.AddEventHandler("MouseOut", func(e *ui.Event) bool {
		log.Printf("Got mouse out event")
		return true
	})

	button := ui.NewClickButton("I'm a button!", font)
	button.AddEventHandler(ui.ButtonClicked, func(e *ui.Event) bool {
		counter++
		label.SetText(fmt.Sprintf("Button Clicked %d times", counter))
		return true
	})
	button.SetDimensions(200, 49)
	button.SetMargin(ui.Margin{3, 3, 3, 3})

	v := ui.NewVerticalContainer()
	v.SetBounds(0, 0, 800, 600)
	c := ui.NewContainer()
	c.SetDimensions(0, 200)
	c.Add(label)
	c.Add(label2)
	c.Add(label3)
	c.Add(button)

	v.Add(c)

	sdlEventHandler := &ui.SDLEventHandler{
		Root: v,
	}

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
			sdlEventHandler.Handle(event)
		}

		v.Layout()

		label3.SetText(fmt.Sprintf("c: %v", c.Bounds()))

		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		v.Draw(renderer)

		renderer.Present()
		sdl.Delay(33)
	}

	c.Destroy()
}
