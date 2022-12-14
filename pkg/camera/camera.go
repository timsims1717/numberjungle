package camera

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
	"numberjungle/pkg/timing"
	"time"
)

var (
	Main *Camera
)

type Camera struct {
	Height float64
	Width  float64
	Mat    pixel.Matrix
	Pos    pixel.Vec
	APos   pixel.Vec
	Zoom   float64
	AZoom  float64
	zStep  float64
	Opt    Options
	Mask   color.RGBA
	IsWin  bool
	iLock  bool

	lock   bool
	random *rand.Rand
}

type Options struct {
	ScrollSpeed float64
	ZoomStep    float64
	ZoomSpeed   float64
	WindowScale float64
}

func New(isWin bool) *Camera {
	return &Camera{
		Mat:   pixel.IM,
		Pos:   pixel.ZV,
		Zoom:  1.0,
		zStep: 1.0,
		Opt: Options{
			ScrollSpeed: 40.0,
			ZoomStep:    1.2,
			ZoomSpeed:   0.2,
			WindowScale: 900.,
		},
		Mask:   colornames.White,
		IsWin:  isWin,
		random: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (c *Camera) SetSize(width, height float64) {
	c.Width = width
	c.Height = height
}

func (c *Camera) GetZoomScale() float64 {
	return 1 / c.Zoom
}

func (c *Camera) GetZoom() float64 {
	return c.Zoom
}

func (c *Camera) Moving() bool {
	return c.lock
}

func (c *Camera) Restrict(bl, tr pixel.Vec) {
	world := c.Pos
	if bl.X <= tr.X {
		if bl.X > world.X {
			world.X = bl.X
		} else if tr.X < world.X {
			world.X = tr.X
		}
	}
	if bl.Y <= tr.Y {
		if bl.Y > world.Y {
			world.Y = bl.Y
		} else if tr.Y < world.Y {
			world.Y = tr.Y
		}
	}
	c.Pos = world
}

func (c *Camera) Update(win *pixelgl.Window) {
	if c.IsWin {
		c.SetSize(win.Bounds().W(), win.Bounds().H())
	}
	c.APos = c.Pos
	c.AZoom = c.Zoom
	if c.iLock {
		c.APos.X = math.Round(c.APos.X)
		c.APos.Y = math.Round(c.APos.Y)
	}
	c.Mat = pixel.IM.Scaled(c.APos, c.AZoom).Moved(win.Bounds().Center().Sub(c.APos))
	win.SetMatrix(c.Mat)
	win.SetColorMask(c.Mask)
}

func (c *Camera) SnapTo(v pixel.Vec) {
	if !c.lock {
		c.Pos.X = v.X
		c.Pos.Y = v.Y
	}
}

func (c *Camera) StayWithin(v pixel.Vec, d float64) {
	if !c.lock {
		if c.Pos.X >= v.X+d {
			c.Pos.X = v.X + d
		} else if c.Pos.X <= v.X-d {
			c.Pos.X = v.X - d
		}
		if c.Pos.Y >= v.Y+d {
			c.Pos.Y = v.Y + d
		} else if c.Pos.Y <= v.Y-d {
			c.Pos.Y = v.Y - d
		}
	}
}

func (c *Camera) Follow(v pixel.Vec, spd float64) {
	if !c.lock {
		c.Pos.X += spd * timing.DT * (v.X - c.Pos.X)
		c.Pos.Y += spd * timing.DT * (v.Y - c.Pos.Y)
	}
}

func (c *Camera) CenterOn(points []pixel.Vec) {
	if !c.lock {
		if points == nil || len(points) == 0 {
			return
		} else if len(points) == 1 {
			c.Pos = points[0]
		} else {
			// todo: center on multiple points + change zoom
		}
	}
}

func (c *Camera) Left() {
	if !c.lock {
		c.Pos.X -= c.Opt.ScrollSpeed * timing.DT
	}
}

func (c *Camera) Right() {
	if !c.lock {
		c.Pos.X += c.Opt.ScrollSpeed * timing.DT
	}
}

func (c *Camera) Down() {
	if !c.lock {
		c.Pos.Y -= c.Opt.ScrollSpeed * timing.DT
	}
}

func (c *Camera) Up() {
	if !c.lock {
		c.Pos.Y += c.Opt.ScrollSpeed * timing.DT
	}
}

func (c *Camera) SetZoom(zoom float64) {
	c.Zoom = zoom
	c.zStep = zoom
}

func (c *Camera) SetILock(b bool) {
	c.iLock = b
}

func (c *Camera) GetColor() color.RGBA {
	return c.Mask
}

func (c *Camera) SetColor(col color.RGBA) {
	c.Mask = col
}

//var (
//	Main *Camera
//)
//
//type Camera struct {
//	Canvas *pixelgl.Canvas
//	Height float64
//	Width  float64
//	Mat    pixel.Matrix
//	Pos    pixel.Vec
//	APos   pixel.Vec
//	Zoom   float64
//	AZoom  float64
//	zStep  float64
//	Opt    Options
//	Mask   color.RGBA
//	iLock  bool
//
//	lock   bool
//	random *rand.Rand
//}
//
//type Options struct {
//	ScrollSpeed float64
//	ZoomStep    float64
//	ZoomSpeed   float64
//}
//
//func New(canvas *pixelgl.Canvas) *Camera {
//	return &Camera{
//		Canvas: canvas,
//		Mat:    pixel.IM,
//		Pos:    pixel.ZV,
//		Zoom:   1.0,
//		zStep:  1.0,
//		Opt: Options{
//			ScrollSpeed: 40.0,
//			ZoomStep:    1.2,
//			ZoomSpeed:   0.2,
//		},
//		Mask:   colornames.White,
//		random: rand.New(rand.NewSource(time.Now().Unix())),
//	}
//}
//
//func (c *Camera) Update(win *pixelgl.Window) {
//	c.APos = c.Pos
//	if c.iLock {
//		c.APos.X = math.Round(c.APos.X)
//		c.APos.Y = math.Round(c.APos.Y)
//	}
//	w := c.Width * 0.5
//	h := c.Height * 0.5
//	r := pixel.R(c.APos.X - w, c.APos.Y - h, c.APos.X + w, c.APos.Y + h)
//	c.Canvas.SetBounds(r)
//	c.Mat = pixel.IM.Scaled(c.APos, c.AZoom).Moved(win.Bounds().Center().Sub(c.APos))
//	win.SetMatrix(c.Mat)
//	win.SetColorMask(c.Mask)
//}
//
//func (c *Camera) SnapTo(v pixel.Vec) {
//	c.Pos.X = v.X
//	c.Pos.Y = v.Y
//}
//
//func (c *Camera) SetSize(width, height float64) {
//	c.Width = width
//	c.Height = height
//}
//
//func (c *Camera) GetZoomScale() float64 {
//	return 1 / c.Zoom
//}
//
//func (c *Camera) GetZoom() float64 {
//	return c.Zoom
//}
//
//func (c *Camera) SetZoom(zoom float64) {
//	c.Zoom = zoom
//	c.zStep = zoom
//}
//
//func (c *Camera) Left() {
//	if !c.lock {
//		c.Pos.X -= c.Opt.ScrollSpeed * timing.DT
//	}
//}
//
//func (c *Camera) Right() {
//	if !c.lock {
//		c.Pos.X += c.Opt.ScrollSpeed * timing.DT
//	}
//}
//
//func (c *Camera) Down() {
//	if !c.lock {
//		c.Pos.Y -= c.Opt.ScrollSpeed * timing.DT
//	}
//}
//
//func (c *Camera) Up() {
//	if !c.lock {
//		c.Pos.Y += c.Opt.ScrollSpeed * timing.DT
//	}
//}
//
////func (c *Camera) ZoomIn(zoom float64) {
////	if !c.lock {
////		c.zStep *= math.Pow(c.Opt.ZoomStep, zoom)
////		c.interZ = gween.New(c.Zoom, c.zStep, c.Opt.ZoomSpeed, ease.OutQuad)
////	}
////}
//
//func (c *Camera) SetILock(b bool) {
//	c.iLock = b
//}
//
//func (c *Camera) SetMask(col color.RGBA) {
//	c.Mask = col
//}