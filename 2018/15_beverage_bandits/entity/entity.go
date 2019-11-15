package entity

import (
	"fmt"
	"image"
	"sort"
)

type Type = rune

const (
	GOBLIN Type = 'G'
	ELF    Type = 'E'
)

type Position = image.Point

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func dist(a, b Position) int {
	return abs(b.X-a.X) + abs(b.Y-a.Y)
}

type Entity struct {
	Type Type
	Pos  Position
	HP   int
	AP   int
}

func (entity *Entity) Attack(target *Entity) error {
	if d := dist(entity.Pos, target.Pos); d != 1 {
		return NotInReachError{Source: entity, Target: target, Dist: d}
	}
	if !Alive(entity) {
		return SourceDead{}
	}
	if !Alive(target) {
		return TargetDead{}
	}
	target.HP -= entity.AP
	return nil
}

func (entity *Entity) Move(p Position) {
	entity.Pos = p
}

func New(t Type, p Position, AP int) *Entity {
	return &Entity{Type: t, Pos: p, HP: 200, AP: AP}
}

type NotInReachError struct {
	Source *Entity
	Target *Entity
	Dist   int
}

func (e NotInReachError) Error() string {
	return fmt.Sprintf("Entity [%+v] unable to reach [%+v], dist=%d", e.Source, e.Target, e.Dist)
}

type SourceDead struct{}

func (e SourceDead) Error() string {
	return "Source is dead"
}

type TargetDead struct{}

func (e TargetDead) Error() string {
	return "Target is dead"
}

type Entities []*Entity

func (ent Entities) Filter(filters ...func(*Entity) bool) (r Entities) {
EntityLoop:
	for _, e := range ent {
		for _, f := range filters {
			if !f(e) {
				continue EntityLoop
			}
		}
		r = append(r, e)
	}
	return
}

func (ent Entities) Any(f func(*Entity) bool) bool {
	for _, e := range ent {
		if f(e) {
			return true
		}
	}
	return false
}

func (ent Entities) Sort() {
	sort.Slice(ent, func(i, j int) bool {
		return LessThan(ent[i].Pos, ent[j].Pos)
	})
}

func LessThan(pos Position, other Position) bool {
	return pos.Y < other.Y || (pos.Y == other.Y && pos.X < other.X)
}

func At(x, y int) func(*Entity) bool {
	p := Pos(x, y)
	return func(e *Entity) bool { return e.Pos.Eq(p) }
}

func NotOfType(t Type) func(*Entity) bool {
	return func(e *Entity) bool { return e.Type != t }
}

func OfType(t Type) func(*Entity) bool {
	return func(e *Entity) bool { return e.Type == t }
}

func Alive(e *Entity) bool {
	return e.HP > 0
}
