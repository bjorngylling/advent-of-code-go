package grid

import "image"

type Type uint8
type Position = image.Point
type Grid struct {
	m      []Type
	Width  int
	Height int
}

const (
	Sand Type = iota
	WetSand

	Clay
	Water
)

func (t Type) String() (s string) {
	switch t {
	case 0:
		s = "sand"
	case 1:
		s = "wet sand"
	case 2:
		s = "clay"
	case 3:
		s = "water"
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

func (c *Grid) Blocked(pos Position) bool {
	return c.m[pos.X+pos.Y*c.Width] >= Clay
}

func (c *Grid) ViewPort(padding int) image.Rectangle {
	view := image.Rectangle{Min: image.Pt(c.Width, c.Height), Max: image.Pt(0, 0)}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			if c.At(Pos(x, y)) == Clay {
				if view.Min.X > x {
					view.Min.X = x
				}
				if view.Max.X < x {
					view.Max.X = x
				}
				if view.Min.Y > y {
					view.Min.Y = y
				}
				if view.Max.Y < y {
					view.Max.Y = y
				}
			}
		}
	}
	return view.Inset(-padding).Intersect(image.Rect(0, 0, c.Width, c.Height))
}
