package data

type Direction int

const (
	Up = iota
	Down
	Left
	Right
)

type Movement struct {
	Direction Direction
	Return    bool
}