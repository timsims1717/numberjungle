package states

import (
	"github.com/bytearena/ecs"
	"math"
	"math/rand"
	"numberjungle/internal/myecs"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
)

var (
	ty = 396.
	by = -12.
	lx = -140.
	rx = 492.
	tys = 404.
	bys = -20.
	lxs = -148.
	rxs = 500.

	gamePanel   *ecs.Entity
	editorPanel *ecs.Entity
	scrollBar   *ecs.Entity
)

func initUIBorder() {
	// top row
	for i := 0; i < 19; i++ {
		mt := img.Batchers["ui_top"].GetSprite("straight_map")
		mtObj := object.New()
		mtObj.Pos.X = lx + float64(i+1) * 32.
		mtObj.Pos.Y = tys
		mtObj.Flip = rand.Intn(2) == 0
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, mtObj).
			AddComponent(myecs.Drawable, mt)
	}
	// bottom row
	for i := 0; i < 19; i++ {
		mb := img.Batchers["ui_top"].GetSprite("straight_map")
		mbObj := object.New()
		mbObj.Pos.X = lx + float64(i+1) * 32.
		mbObj.Pos.Y = bys
		mbObj.Flop = true
		mbObj.Flip = rand.Intn(2) == 0
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, mbObj).
			AddComponent(myecs.Drawable, mb)
	}
	// left side
	for i := 0; i < 12; i++ {
		ml := img.Batchers["ui_top"].GetSprite("straight_map")
		mlObj := object.New()
		mlObj.Pos.X = lxs
		mlObj.Pos.Y = by + float64(i+1) * 32.
		mlObj.Rot = math.Pi * 0.5
		mlObj.Flop = rand.Intn(2) == 0
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, mlObj).
			AddComponent(myecs.Drawable, ml)
	}
	// right side
	for i := 0; i < 12; i++ {
		mr := img.Batchers["ui_top"].GetSprite("straight_map")
		mrObj := object.New()
		mrObj.Pos.X = rxs
		mrObj.Pos.Y = by + float64(i+1) * 32.
		mrObj.Rot = math.Pi * -0.5
		mrObj.Flop = rand.Intn(2) == 0
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, mrObj).
			AddComponent(myecs.Drawable, mr)
	}

	mtl := img.Batchers["ui_top"].GetSprite("corner_map")
	mtlObj := object.New()
	mtlObj.Pos.X = lx
	mtlObj.Pos.Y = ty
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mtlObj).
		AddComponent(myecs.Drawable, mtl)
	mbl := img.Batchers["ui_top"].GetSprite("corner_map")
	mblObj := object.New()
	mblObj.Pos.X = lx
	mblObj.Pos.Y = by
	mblObj.Flop = true
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mblObj).
		AddComponent(myecs.Drawable, mbl)
	mtr := img.Batchers["ui_top"].GetSprite("corner_map")
	mtrObj := object.New()
	mtrObj.Pos.X = rx
	mtrObj.Pos.Y = ty
	mtrObj.Flip = true
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mtrObj).
		AddComponent(myecs.Drawable, mtr)
	mbr := img.Batchers["ui_top"].GetSprite("corner_map")
	mbrObj := object.New()
	mbrObj.Pos.X = rx
	mbrObj.Pos.Y = by
	mbrObj.Flip = true
	mbrObj.Flop = true
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mbrObj).
		AddComponent(myecs.Drawable, mbr)
}

func initGamePanel() {
	uiPanel := img.Batchers["ui"].GetSprite("game_panel")
	uiObj := object.New()
	uiObj.Pos.X = -80.
	uiObj.Pos.Y = 192.
	gamePanel = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, uiObj).
		AddComponent(myecs.Drawable, uiPanel)
}

func initEditorPanel() {
	uiPanel := img.Batchers["ui"].GetSprite("editor_panel")
	uiObj := object.New()
	uiObj.Pos.X = -80.
	uiObj.Pos.Y = 192.
	editorPanel = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, uiObj).
		AddComponent(myecs.Drawable, uiPanel)
	scrollTop := img.Batchers["ui"].GetSprite("scroll_bar_top")
	tObj := object.New()
	tObj.Pos.X = -20.
	tObj.Pos.Y = 284.
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tObj).
		AddComponent(myecs.Drawable, scrollTop)
	scrollMid := img.Batchers["ui"].GetSprite("scroll_bar_middle")
	mObj := object.New()
	mObj.Pos.X = -20.
	mObj.Pos.Y = 138.
	mObj.Scalar.Y = 18.
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mObj).
		AddComponent(myecs.Drawable, scrollMid)
	scrollBot := img.Batchers["ui"].GetSprite("scroll_bar_bottom")
	bObj := object.New()
	bObj.Pos.X = -20.
	bObj.Pos.Y = -8.
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, bObj).
		AddComponent(myecs.Drawable, scrollBot)
	scrollBarSpr := img.Batchers["ui_top"].GetSprite("scroll_bar")
	sObj := object.New()
	sObj.Pos.X = -20.
	sObj.Pos.Y = 284.
	scrollBar = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, sObj).
		AddComponent(myecs.Drawable, scrollBarSpr)
}