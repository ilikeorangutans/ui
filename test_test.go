package ui

import (
	"runtime"

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
	font = loadFont()
}

func loadFont() *ttf.Font {
	var fontPath string
	if runtime.GOOS == "darwin" {
		fontPath = "/Library/Fonts/Verdana.ttf"
	} else if runtime.GOOS == "linux" {
		fontPath = "/usr/share/fonts/truetype/droid/DroidSansFallbackFull.ttf"
	}
	font, err := ttf.OpenFont(fontPath, 12)
	if err != nil {
		panic(err)
	}

	return font
}

type TestWidget struct {
	BoxModel
	EventHandlers
}

func (w *TestWidget) Draw(renderer *sdl.Renderer) {}
func (w *TestWidget) Visit(v WidgetVisitor)       {}
