package data

import (
	"numberjungle/internal/vars"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
	"numberjungle/pkg/world"
)

type Tile struct {
	Coords   world.Coords
	Object   *object.Object
	Sprite   *img.Sprite
}

type Puzzle struct {
	Tiles [vars.Height][vars.Width]Tile
	Start world.Coords
}