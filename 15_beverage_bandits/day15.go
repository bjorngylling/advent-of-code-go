package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"strings"
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

func (c *Cave) SetBlocked(x, y int, val bool) {
	c.m[x+y*c.Width] = val
}

func parseInput(in string) (*Cave, Entities) {
	lines := strings.Split(in, "\n")
	cave := NewCave(len(lines[0]), len(lines))
	var entities Entities

	for y, ln := range lines {
		for x, c := range ln {
			switch c {
			case '#':
				cave.SetBlocked(x, y, true)
			case 'G':
				entities = append(entities, NewEntity(GOBLIN, Pos(x, y)))
			case 'E':
				entities = append(entities, NewEntity(ELF, Pos(x, y)))
			}
		}
	}

	return cave, entities
}

func main() {
	fileContent, err := ioutil.ReadFile("15_beverage_bandits/day15_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	cave, entities := parseInput(string(fileContent))
	fmt.Printf("Day 15 part 1 result: %+v, %+v\n", len(entities), cave.Blocked(0, 0))

	fmt.Printf("Day 15 part 2 result: %+v\n", nil)
}
