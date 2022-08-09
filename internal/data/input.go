package data

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/timsims1717/pixel-go-input"
)

var TheInput = &pxginput.Input{
	Buttons: map[string]*pxginput.ButtonSet{
		"moveLeft":  pxginput.NewJoyless(pixelgl.KeyLeft),
		"moveRight": pxginput.NewJoyless(pixelgl.KeyRight),
		"moveUp":    pxginput.NewJoyless(pixelgl.KeyUp),
		"moveDown":  pxginput.NewJoyless(pixelgl.KeyDown),
		"camLeft":   pxginput.NewJoyless(pixelgl.KeyA),
		"camRight":  pxginput.NewJoyless(pixelgl.KeyD),
		"camUp":     pxginput.NewJoyless(pixelgl.KeyW),
		"camDown":   pxginput.NewJoyless(pixelgl.KeyS),
		"zoomIn":    pxginput.NewJoyless(pixelgl.KeyR),
		"zoomOut":   pxginput.NewJoyless(pixelgl.KeyF),
		"click": {
			Keys:   []pixelgl.Button{pixelgl.MouseButtonLeft, pixelgl.KeySpace},
			Scroll: 1,
		},
		"back": pxginput.NewJoyless(pixelgl.KeyEscape),
	},
	Mode: pxginput.KeyboardMouse,
}
