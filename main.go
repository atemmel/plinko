package main

import "github.com/veandco/go-sdl2/sdl"

var ball []sdl.Point
var pin  []sdl.Point

func initCircle(radius float64) []sdl.Point {
	pts := make([]sdl.Point, 0, 64)


	return pts
}

func movePts(pts []sdl.Point, x, y int32) []sdl.Point { 
	mvpts := make([]sdl.Point, len(pts))

	for i := range mvpts {
		p0 := pts[i]
		p0.X += x
		p0.Y += y
		mvpts[i] = p0
	}

	return mvpts
}

func init() {

}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rend, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer rend.Destroy()

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255}
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
		rend.Clear()
		rend.Present()
	}
}
