package states

import (
	"github.com/faiface/pixel"
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/internal/systems"
	"numberjungle/internal/vars"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
	"numberjungle/pkg/world"
)

func startPuzzle(coords world.Coords) {
	data.NewExpression()
	setPlayerStart(coords)
}

func setPlayerStart(coords world.Coords) {
	if data.Player != nil {
		if data.Player.Move != nil {
			data.Player.Move = nil
		}
		if data.Player.Char != nil {
			myecs.Manager.DisposeEntity(data.Player.Char)
			data.Player.Char = nil
		}
	}
	data.NewPlayer()
	shipCoords := coords
	shipCoords.Y++
	pirateSpr := img.Batchers["pirate"].GetSprite("standing")
	pObj := object.New()
	pObj.Pos = world.MapToWorld(coords)
	data.Player.Char = myecs.Manager.NewEntity()
	data.Player.Char.AddComponent(myecs.Drawable, pirateSpr).
		AddComponent(myecs.Object, pObj).
		AddComponent(myecs.Coords, coords).
		AddComponent(myecs.Player, myecs.Has{})
	shipSpr := img.Batchers[vars.JungleBatch].GetSprite("ship")
	shipObj := object.New()
	shipObj.Pos = world.MapToWorld(shipCoords)
	myecs.Manager.NewEntity().AddComponent(myecs.Drawable, shipSpr).
		AddComponent(myecs.Object, shipObj).
		AddComponent(myecs.Coords, shipCoords)
}

func updatePuzzle() {
	if data.Player.Move == nil {
		if data.TheInput.Get("moveUp").JustPressed() {
			data.Player.Move = &data.Movement{Direction: data.Up}
		} else if data.TheInput.Get("moveDown").JustPressed() {
			data.Player.Move = &data.Movement{Direction: data.Down}
		} else if data.TheInput.Get("moveLeft").JustPressed() {
			data.Player.Move = &data.Movement{Direction: data.Left}
		} else if data.TheInput.Get("moveRight").JustPressed() {
			data.Player.Move = &data.Movement{Direction: data.Right}
		}
	}
	if !data.Player.Char.HasComponent(myecs.Moving) && data.Player.Move != nil {
		data.Player.Char.AddComponent(myecs.Moving, data.Player.Move)
		data.Player.Move = nil
	}

	systems.TemporarySystem()
	systems.MovementSystem()
	systems.CollectingSystem()
	systems.ExprPosSystem()
	systems.ExprCheckSystem()
	systems.FullTransformSystem()
	systems.AnimationSystem()
}

func blankPuzzle() *data.Puzzle {
	coords := world.Coords{
		X: 0,
		Y: vars.Height - 1,
	}
	shipCoords := coords
	shipCoords.Y++

	puzzle := &data.Puzzle{
		Tiles: [vars.RealHeight][vars.Width]*data.Tile{},
		Start: coords,
		End:   shipCoords,
	}

	for y, row := range puzzle.Tiles {
		for x := range row {
			tCoords := world.Coords{X: x, Y: y}
			obj := object.New()
			obj.Pos = world.MapToWorld(tCoords)
			obj.Rect = pixel.R(-16., -16., 16., 16.)
			key := "beach"
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
			if y == vars.Height && x != coords.X {
				t.AddComponent(myecs.Occupy, myecs.Has{})
			}
			tile.Entity = t
			puzzle.Tiles[y][x] = tile
		}
	}
	return puzzle
}