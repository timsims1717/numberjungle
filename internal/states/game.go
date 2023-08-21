package states

import (
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/internal/systems"
	"numberjungle/internal/vars"
	"numberjungle/pkg/camera"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
	"numberjungle/pkg/state"
	"numberjungle/pkg/world"
)

var GameState = &gameState{}

type gameState struct {
	*state.AbstractState
}

func (s *gameState) Unload() {
	systems.ClearSystem()
}

func (s *gameState) Load() {
	camera.Main.Pos.X = (vars.Width - 5) * world.TileSize * 0.5
	camera.Main.Pos.Y = (vars.Height) * world.TileSize * 0.5

	initUIBorder()
	initGamePanel()

	coords := world.Coords{
		X: rand.Intn(vars.Width),
		Y: vars.Height - 1,
	}
	shipCoords := coords
	shipCoords.Y = vars.Height

	data.CurrPuzzle = &data.Puzzle{
		Tiles: [vars.RealHeight][vars.Width]*data.Tile{},
		Start: coords,
		End:   shipCoords,
	}

	for y, row := range data.CurrPuzzle.Tiles {
		for x := range row {
			tCoords := world.Coords{X: x, Y: y}
			obj := object.New()
			obj.Pos = world.MapToWorld(tCoords)
			key := "beach"
			if tCoords != coords {
				c := rand.Intn(8)
				switch c {
				case 0, 1:
					key = "water"
				case 2, 3:
					key = "grass"
				case 4:
					key = "stone"
				}
			}
			if y == vars.Height {
				key = "water"
			}
			tile := &data.Tile{
				Coords: tCoords,
				Object: obj,
				Sprite: img.Batchers["tiles"].GetSprite(key),
			}
			t := myecs.Manager.NewEntity().
				AddComponent(myecs.Drawable, tile.Sprite).
				AddComponent(myecs.Object, tile.Object).
				AddComponent(myecs.Coords, tCoords)
			if key == "stone" || (y == vars.Height && x != coords.X) {
				t.AddComponent(myecs.Occupy, myecs.Has{})
			} else if key == "grass" || key == "beach" && rand.Intn(50) == 0 {
				coin := data.NewRandomCoin(tCoords)
				coin.Entity = myecs.Manager.NewEntity().
					AddComponent(myecs.Object, coin.Object).
					AddComponent(myecs.Coords, tCoords).
					AddComponent(myecs.Drawable, coin.Sprs).
					AddComponent(myecs.Collect, coin)
			}
			tile.Entity = t
			data.CurrPuzzle.Tiles[y][x] = tile
		}
	}
	startPuzzle(coords)
}

func (s *gameState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win, camera.Main.Mat)
	updatePuzzle()
	camera.Main.Update(win)
}

func (s *gameState) Draw(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win)
	img.Draw(win)
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}