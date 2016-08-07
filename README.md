# ui

An experiment in building a simple ui toolkit using [github.com/veandco/go-sdl2](https://github.com/veandco/go-sdl2) as a renderer.

## What does it look like?

Pretty ugly.

Buttons and Spinner:
![Buttons and spinner](http://g.recordit.co/O9S5Cf5DkG.gif)

## Status

- Basic layouting works using a simple box model
- Simple label widget to render text
- Click and toggle buttons
- Numeric Spinner

## Using ui

- `go get github.com/ilikeorangutans/ui` 
- If you're using vendoring, make sure you add only `github.com/ilikeorangutans/ui` and not any of the SDL dependencies, otherwise you might get errors stating `...cannot use renderer (type *".../your_project/vendor/github.com/veandco/go-sdl2/sdl".Renderer) as type "github.com/ilikeorangutans/ui/vendor/github.com/veandco/go-sdl2/sdl".Renderer ...`
