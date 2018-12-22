package entity

import "image"

type Type = string

const (
	GOBLIN Type = "G"
	ELF         = "E"
)

type Position = image.Point

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}

type Entity struct {
	Type Type
	Pos  Position
	HP   int
	AP   int
}

func New(t Type, p Position) *Entity {
	return &Entity{Type: t, Pos: p, HP: 200, AP: 3}
}

type Entities []*Entity

func (entities Entities) Filter(f func(*Entity) bool) (r Entities) {
	for _, e := range entities {
		if f(e) {
			r = append(r, e)
		}
	}
	return
}

func (entities Entities) Any(f func(*Entity) bool) bool {
	for _, e := range entities {
		if f(e) {
			return true
		}
	}
	return false
}

func At(x, y int) func(*Entity) bool {
	p := Pos(x, y)
	return func(e *Entity) bool { return e.Pos.Eq(p) }
}
