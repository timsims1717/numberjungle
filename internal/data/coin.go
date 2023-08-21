package data

import (
	"github.com/bytearena/ecs"
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
	LastCoin
)

type Coin struct {
	Coords world.Coords
	Object *object.Object
	Value  CoinValue
	Sprs   []*img.Sprite
	Entity *ecs.Entity
}

func NewRandomCoin(coords world.Coords) *Coin {
	v := CoinValue(rand.Intn(LastCoin))
	obj := object.New()
	obj.Pos = world.MapToWorld(coords)
	return &Coin{
		Coords: coords,
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

func (c CoinValue) Rune() rune {
	switch c {
	case One:
		return '1'
	case Two:
		return '2'
	case Three:
		return '3'
	case Four:
		return '4'
	case Five:
		return '5'
	case Six:
		return '6'
	case Seven:
		return '7'
	case Eight:
		return '8'
	case Nine:
		return '9'
	case Zero:
		return '0'
	case Plus:
		return '+'
	case Minus:
		return '-'
	case Equals:
		return '='
	case Times:
		return '*'
	case Divide:
		return '/'
	default:
		return '?'
	}
}

func (c CoinValue) Int() int {
	switch c {
	case One:
		return 1
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Zero:
		return 0
	default:
		return -1
	}
}

func (c CoinValue) IsInt() bool {
	switch c {
	case One,Two,Three,Four,Five,Six,Seven,Eight,Nine,Zero:
		return true
	default:
		return false
	}
}

func (c CoinValue) IsOperator() bool {
	switch c {
	case Equals,Plus,Minus,Times,Divide:
		return true
	default:
		return false
	}
}

func (c CoinValue) Priority() int {
	switch c {
	case Plus,Minus:
		return 1
	case Times,Divide:
		return 2
	case Equals:
		return 4
	default:
		return -1
	}
}