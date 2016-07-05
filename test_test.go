package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

var font *ttf.Font

func init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	err = ttf.Init()
	if err != nil {
		panic(err)
	}
	font, err = ttf.OpenFont("/usr/share/fonts/truetype/ttf-bitstream-vera/Vera.ttf", 11)
	if err != nil {
		panic(err)
	}
}
