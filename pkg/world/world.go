package world

import "github.com/faiface/pixel"

var (
	TileSize float64
	Origin   = Coords{
		X: 0,
		Y: 0,
	}
)

func SetTileSize(s float64) {
	TileSize = s
}

func MapToWorld(a Coords) pixel.Vec {
	return pixel.V(float64(a.X)*TileSize, float64(a.Y)*TileSize)
}

func WorldToMap(x, y float64) (int, int) {
	return int(x / TileSize), int(y / TileSize)
}
