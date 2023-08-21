package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"numberjungle/internal/states"
	"numberjungle/internal/vars"
	"numberjungle/pkg/camera"
	"numberjungle/pkg/img"
	"numberjungle/pkg/state"
	"numberjungle/pkg/timing"
	"numberjungle/pkg/world"
	"time"
)

func run() {
	rand.Seed(time.Now().Unix())
	world.SetTileSize(32.)
	conf := pixelgl.WindowConfig{
		Title:     vars.Title,
		Bounds:    pixel.R(0, 0, 1600, 900),
		VSync:     true,
		Invisible: true,
	}
	win, err := pixelgl.NewWindow(conf)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(false)

	camera.Main = camera.New(true)
	camera.Main.SetILock(true)
	camera.Main.SetSize(1600, 900)
	camera.Main.SetZoom(1.8)

	uiSheet, err := img.LoadSpriteSheet("assets/img/ui.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("ui", uiSheet, true, true)
	tileSheet, err := img.LoadSpriteSheet("assets/img/tiles.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("tiles", tileSheet, true, true)
	jungleObjSheet, err := img.LoadSpriteSheet("assets/img/jungleobj.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(vars.JungleBatch, jungleObjSheet, true, true)
	pirateSheet, err := img.LoadSpriteSheet("assets/img/pirate.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("pirate", pirateSheet, true, true)
	img.AddBatcher("ui_top", uiSheet, true, true)

	state.Register("game", state.New(states.GameState))
	state.Register("editor", state.New(states.EditorState))
	state.SwitchState("editor")

	timing.SetTargetFPS(20)
	timing.Reset()
	win.Show()
	for !win.Closed() {
		timing.Update()
		img.FullClear()

		state.Update(win)

		win.Clear(vars.BGColor)

		state.Draw(win)

		win.Update()
		timing.Wait()
	}
}

func main() {
	pixelgl.Run(run)
}