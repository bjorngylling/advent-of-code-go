package grid

import (
	"crypto"
	_ "crypto/sha1"
	"fmt"
	"image"
)

type Type = byte
type Position = image.Point
type Grid struct {
	m      []Type
	Width  int
	Height int
}

const (
	Open Type = iota
	Trees
	Lumberyard
)

func TypeToString(t Type) (s string) {
	switch t {
	case 0:
		s = "open"
	case 1:
		s = "trees"
	case 2:
		s = "lumberyard"
	}
	return
}

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}

func New(w, h int) *Grid {
	return &Grid{m: make([]Type, w*h), Width: w, Height: h}
}

func (c *Grid) At(pos Position) Type {
	return c.m[pos.X+pos.Y*c.Width]
}

func (c *Grid) Set(pos Position, val Type) {
	c.m[pos.X+pos.Y*c.Width] = val
}

func (c *Grid) Hash() string {
	sha1 := crypto.SHA1.New()
	sha1.Write(c.m)
	return fmt.Sprintf("%x", sha1.Sum(nil))
}
