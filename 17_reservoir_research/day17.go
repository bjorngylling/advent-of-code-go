package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code-2018/17_reservoir_research/grid"
	"image"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(s string) *grid.Grid {
	c := grid.New(700, 2000)
	for _, ln := range strings.Split(s, "\n") {
		lh, rh := splitPair(ln, ", ")
		var x1, x2, y1, y2 int

		n, v := splitPair(lh, "=")
		if n == "x" {
			x1, _ = strconv.Atoi(v)
			x2 = x1
		} else if n == "y" {
			y1, _ = strconv.Atoi(v)
			y2 = y1
		}

		n, v = splitPair(rh, "=")
		v1, v2 := splitPair(v, "..")
		if n == "x" {
			x1, _ = strconv.Atoi(v1)
			x2, _ = strconv.Atoi(v2)
		} else if n == "y" {
			y1, _ = strconv.Atoi(v1)
			y2, _ = strconv.Atoi(v2)
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				c.Set(grid.Pos(x, y), grid.Clay)
			}
		}
	}
	return c
}

func splitPair(s string, sep string) (lh string, rh string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

func renderWorld(w *grid.Grid, source grid.Position, viewPort image.Rectangle) (s string) {
	for y := viewPort.Min.Y; y <= viewPort.Max.Y; y++ {
		for x := viewPort.Min.X; x <= viewPort.Max.X; x++ {
			if grid.Pos(x, y).Eq(source) {
				s += "+"
			} else {
				switch w.At(grid.Pos(x, y)) {
				case grid.Sand:
					s += "."
				case grid.WetSand:
					s += "|"
				case grid.Clay:
					s += "#"
				case grid.Water:
					s += "~"
				}
			}
		}
		s += "\n"
	}
	return
}

var gravity = grid.Pos(0, 1)
var left = grid.Pos(-1, 0)
var right = grid.Pos(1, 0)

func simulate(source grid.Position, world *grid.Grid, maxY int) {
	var flow func(grid.Position)
	flow = func(pos grid.Position) {
		// End this flow if we reach the bottom or end up on a blocked position
		if pos.Y > maxY || world.Blocked(pos) {
			return
		}

		if world.Blocked(pos.Add(gravity)) {
			// If below is blocked flow to the left and right
			// Attempt to find boundary or space that isn't blocked below to the left and right
			l := pos
			for !world.Blocked(l) && world.Blocked(l.Add(gravity)) {
				world.Set(l, grid.WetSand)
				l = l.Add(left)
			}
			r := pos
			for !world.Blocked(r) && world.Blocked(r.Add(gravity)) {
				world.Set(r, grid.WetSand)
				r = r.Add(right)
			}

			if !world.Blocked(l.Add(gravity)) || !world.Blocked(r.Add(gravity)) {
				// If open space below flow from here (i.e. recurse and continue flowing down)
				flow(l)
				flow(r)
			} else if world.At(l) == grid.Clay && world.At(r) == grid.Clay {
				// If bounded by clay to the left and right this row should fill with water completely
				for p := l.Add(right); p.X < r.X; p.X++ {
					world.Set(p, grid.Water)
				}
			}
		} else if world.At(pos) == grid.Sand {
			// If flowing into sand, make it wet and continue flowing down
			world.Set(pos, grid.WetSand)
			flow(pos.Add(gravity))
			if world.At(pos.Add(gravity)) == grid.Water {
				// Flow on top of stable water (to end up flowing left and right on top of water)
				flow(pos)
			}
		}
	}

	flow(source)
}

func count(w *grid.Grid, p func(grid.Type) bool) (c int) {
	bounds := w.ViewPort(0)
	for x := bounds.Min.X - 1; x <= bounds.Max.X+1; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			if p(w.At(grid.Pos(x, y))) {
				c += 1
			}
		}
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("17_reservoir_research/day17_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	w := parseInput(string(fileContent))
	source := grid.Pos(500, 0)
	simulate(source, w, w.ViewPort(0).Max.Y)

	wetCount := count(w, func(t grid.Type) bool { return t == grid.WetSand || t == grid.Water })
	fmt.Printf("Day 17 part 1 result: %+v\n", wetCount)

	waterCount := count(w, func(t grid.Type) bool { return t == grid.Water })
	fmt.Printf("Day 17 part 2 result: %+v\n", waterCount)
}
