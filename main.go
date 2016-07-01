package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()

	font, err := ttf.OpenFont("/usr/share/fonts/truetype/ttf-bitstream-vera/Vera.ttf", 11)
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

	var texture *sdl.Texture
	surface, err := font.RenderUTF8_Solid("Horray", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	texture, err = renderer.CreateTextureFromSurface(surface)
	surface.Free()
	running := true
	var event sdl.Event
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyDownEvent:
				log.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				fmt.Sprintf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

			}
		}

		renderer.Clear()

		err = renderer.Copy(texture, nil, nil)
		if err != nil {
			log.Println(err)
		}

		renderer.Present()
		sdl.Delay(33)
	}

}
