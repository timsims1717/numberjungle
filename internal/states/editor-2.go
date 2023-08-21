package states

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"numberjungle/internal/data"
	"numberjungle/internal/myecs"
	"numberjungle/internal/vars"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
)

var (
	edx = -122.
	edy = 266.
	edd = 38.

	editorCanvas *pixelgl.Canvas
	editorTile   *data.Tile
	editorCoin   *data.Coin
	editorSelect *ecs.Entity
)

func initPicker() {
	index := 0
	for i := 0; i < data.LastTileType; i++ {
		obj := object.New()
		obj.Pos.X = edx + edd* float64(index % 3)
		obj.Pos.Y = edy - edd* float64(index / 3)
		obj.Rect = pixel.R(-16., -16., 16., 16.)
		tile := &data.Tile{
			Object: obj,
			Sprite: img.Batchers["tiles"].GetSprite(data.TileType(index).String()),
		}
		t := myecs.Manager.NewEntity().
			AddComponent(myecs.Drawable, tile.Sprite).
			AddComponent(myecs.Object, tile.Object)
		tile.Entity = t
		EditorState.Picker = append(EditorState.Picker, tile)
		if i == 0 {
			esObj := object.New()
			esObj.Pos = obj.Pos
			editorTile = tile
			editorCoin = nil
			editorSelect = myecs.Manager.NewEntity().
				AddComponent(myecs.Drawable, img.Batchers["ui"].GetSprite("editor_select")).
				AddComponent(myecs.Object, esObj)
		}
		index++
	}
	index += 3 - (index % 3)
	for i := 0; i < data.LastCoin; i++ {
		v := data.CoinValue(i)
		obj := object.New()
		obj.Pos.X = edx + edd* float64(index % 3)
		obj.Pos.Y = edy - edd* float64(index / 3)
		obj.Rect = pixel.R(-16., -16., 16., 16.)
		coin := &data.Coin{
			Object: obj,
			Value:  v,
			Sprs:   []*img.Sprite{
				{
					Key:   "coin",
					Color: colornames.White,
					Batch: vars.JungleBatch,
				},
				{
					Key:   v.String(),
					Color: colornames.White,
					Batch: vars.JungleBatch,
				},
			},
		}
		coin.Entity = myecs.Manager.NewEntity().
			AddComponent(myecs.Object, coin.Object).
			AddComponent(myecs.Drawable, coin.Sprs)
		EditorState.Picker = append(EditorState.Picker, coin)
		index++
	}
}