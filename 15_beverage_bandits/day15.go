package main

import (
	"fmt"
	"image"
)

type EntityType = int

const (
	GOBLIN EntityType = iota
	ELF
)

type Position = image.Point

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}

type Entity struct {
	Type EntityType
	Pos  Position
	HP   int
	AP   int
}

func NewEntity(t EntityType, p Position) *Entity {
	return &Entity{Type: t, Pos: p, HP: 200, AP: 3}
}

type Cave struct {
	m      []bool
	Width  int
	Height int
}

func NewCave(w, h int) *Cave {
	return &Cave{m: make([]bool, w*h), Width: w, Height: h}
}

func (c *Cave) Blocked(x, y int) bool {
	return c.m[x+y*c.Width]
}

func main() {
	fmt.Printf("Day 15 part 1 result: %+v\n", nil)

	fmt.Printf("Day 15 part 2 result: %+v\n", nil)
}
