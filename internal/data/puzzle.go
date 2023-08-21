package data

import (
	"github.com/bytearena/ecs"
	"numberjungle/internal/vars"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
	"numberjungle/pkg/world"
)

type Tile struct {
	Coords world.Coords
	Object *object.Object
	Sprite *img.Sprite
	Entity *ecs.Entity
}

var (
	CurrPuzzle *Puzzle
	EditPuzzle *Puzzle
)

type Puzzle struct {
	Tiles [vars.RealHeight][vars.Width]*Tile
	Start world.Coords
	End   world.Coords
	Done  bool
}

type TileType int

const (
	Beach = iota
	Water
	Grass
	Stone
	LastTileType
)

func (t TileType) String() string {
	switch t {
	case Beach:
		return "beach"
	case Water:
		return "water"
	case Grass:
		return "grass"
	case Stone:
		return "stone"
	default:
		return "unknown"
	}
}