package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/internal/systems"
	"numberjungle/internal/vars"
	"numberjungle/pkg/camera"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
	"numberjungle/pkg/timing"
	"numberjungle/pkg/world"
)

func run() {
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
	camera.Main.SetZoom(2.)

	pirate, err := img.LoadImage("assets/img/pirate.png")
	if err != nil {
		panic(err)
	}
	pirateSpr := pixel.NewSprite(pirate, pirate.Bounds())
	pObj := object.New()
	pSpd := 50.
	//pirateSheet, err := img.LoadSpriteSheet("assets/img/pirate.json")
	//if err != nil {
	//	panic(err)
	//}
	//img.AddBatcher("pirate", pirateSheet, true, true)
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Drawable, pirateSpr).
		AddComponent(myecs.Object, pObj)

	timing.SetTargetFPS(20)
	timing.Reset()
	win.Show()
	for !win.Closed() {
		timing.Update()
		data.TheInput.Update(win, camera.Main.Mat)

		if data.TheInput.Get("moveUp").Pressed() {
			pObj.Pos.Y += pSpd * timing.DT
		}
		if data.TheInput.Get("moveDown").Pressed() {
			pObj.Pos.Y -= pSpd * timing.DT
		}
		if data.TheInput.Get("moveLeft").Pressed() {
			pObj.Pos.X -= pSpd * timing.DT
		}
		if data.TheInput.Get("moveRight").Pressed() {
			pObj.Pos.X += pSpd * timing.DT
		}

		systems.FullTransformSystem()
		systems.AnimationSystem()
		camera.Main.Update(win)
		win.Clear(vars.BGColor)
		systems.DrawSystem(win)
		win.Update()
		timing.Wait()
	}
}

func main() {
	pixelgl.Run(run)
}