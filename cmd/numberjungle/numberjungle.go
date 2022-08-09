package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/internal/systems"
	"numberjungle/internal/vars"
	"numberjungle/pkg/camera"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
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
	camera.Main.SetZoom(2.)
	camera.Main.Pos.X = (vars.Width - 1) * world.TileSize * 0.5
	camera.Main.Pos.Y = (vars.Height - 1) * world.TileSize * 0.5

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

	pirateSpr := img.Batchers["pirate"].GetSprite("standing")
	pObj := object.New()
	pirate := myecs.Manager.NewEntity()
	pirate.AddComponent(myecs.Drawable, pirateSpr).
		AddComponent(myecs.Object, pObj).
		AddComponent(myecs.Coords, world.Coords{
			X: 0,
			Y: 0,
		}).
		AddComponent(myecs.Player, myecs.Has{})


	puzzle := data.Puzzle{
		Tiles: [12][16]data.Tile{},
		Start: world.Coords{
			X: 0,
			Y: 0,
		},
	}

	for y, row := range puzzle.Tiles {
		for x := range row {
			coords := world.Coords{X: x, Y: y}
			obj := object.New()
			obj.Pos = world.MapToWorld(coords)
			key := "water"
			c := rand.Intn(4)
			switch c {
			case 0:
				key = "grass"
			case 1:
				key = "beach"
			case 2:
				key = "stone"
			}
			tile := data.Tile{
				Coords: coords,
				Object: obj,
				Sprite: img.Batchers["tiles"].GetSprite(key),
			}
			puzzle.Tiles[y][x] = tile
			t := myecs.Manager.NewEntity().
				AddComponent(myecs.Drawable, tile.Sprite).
				AddComponent(myecs.Object, tile.Object).
				AddComponent(myecs.Coords, coords)
			if key == "stone" {
				t.AddComponent(myecs.Occupy, myecs.Has{})
			} else if key == "grass" || key == "beach" && rand.Intn(20) == 0 {
				coin := data.NewRandomCoin(coords)
				myecs.Manager.NewEntity().
					AddComponent(myecs.Object, coin.Object).
					AddComponent(myecs.Coords, coords).
					AddComponent(myecs.Drawable, coin.Sprites()).
					AddComponent(myecs.Collect, coin)
			}
		}
	}
	var move *data.Movement

	timing.SetTargetFPS(20)
	timing.Reset()
	win.Show()
	for !win.Closed() {
		timing.Update()
		img.FullClear()
		data.TheInput.Update(win, camera.Main.Mat)

		if move == nil {
			if data.TheInput.Get("moveUp").JustPressed() {
				move = &data.Movement{Direction: data.Up}
			} else if data.TheInput.Get("moveDown").JustPressed() {
				move = &data.Movement{Direction: data.Down}
			} else if data.TheInput.Get("moveLeft").JustPressed() {
				move = &data.Movement{Direction: data.Left}
			} else if data.TheInput.Get("moveRight").JustPressed() {
				move = &data.Movement{Direction: data.Right}
			}
		}
		if !pirate.HasComponent(myecs.Moving) && move != nil {
			pirate.AddComponent(myecs.Moving, move)
			move = nil
		}

		systems.TemporarySystem()
		systems.MovementSystem()
		systems.CollectingSystem()
		systems.FullTransformSystem()
		systems.AnimationSystem()
		camera.Main.Update(win)
		win.Clear(vars.BGColor)
		systems.DrawSystem(win)
		img.Draw(win)
		win.Update()
		timing.Wait()
	}
}

func main() {
	pixelgl.Run(run)
}