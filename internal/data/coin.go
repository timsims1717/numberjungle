package data

import (
	"golang.org/x/image/colornames"
	"math/rand"
	"numberjungle/internal/vars"
	"numberjungle/pkg/img"
	"numberjungle/pkg/object"
	"numberjungle/pkg/world"
)

type CoinValue int

const (
	One = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Zero
	Plus
	Minus
	Equals
	Times
	Divide
	Unknown
)

type Coin struct {
	Coords world.Coords
	Object *object.Object
	Value  CoinValue
}

func NewRandomCoin(coords world.Coords) *Coin {
	obj := object.New()
	obj.Pos = world.MapToWorld(coords)
	return &Coin{
		Coords: coords,
		Object: obj,
		Value:  CoinValue(rand.Intn(Unknown)),
	}
}

func (c *Coin) Sprites() []*img.Sprite {
	return []*img.Sprite{
		{
			Key:   "coin",
			Color: colornames.White,
			Batch: vars.JungleBatch,
		},
		{
			Key:   c.Value.String(),
			Color: colornames.White,
			Batch: vars.JungleBatch,
		},
	}
}

func (c CoinValue) String() string {
	switch c {
	case One:
		return "one"
	case Two:
		return "two"
	case Three:
		return "three"
	case Four:
		return "four"
	case Five:
		return "five"
	case Six:
		return "six"
	case Seven:
		return "seven"
	case Eight:
		return "eight"
	case Nine:
		return "nine"
	case Zero:
		return "zero"
	case Plus:
		return "plus"
	case Minus:
		return "minus"
	case Equals:
		return "equals"
	case Times:
		return "times"
	case Divide:
		return "divide"
	default:
		return "unknown"
	}
}