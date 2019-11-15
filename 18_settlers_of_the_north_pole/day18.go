package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code-2018/18_settlers_of_the_north_pole/grid"
	"image"
	"io/ioutil"
	"log"
	"strings"
)

func parseInput(input string) *grid.Grid {
	s := strings.Split(input, "\n")
	g := grid.New(len(s[0]), len(s))

	for y, ln := range s {
		for x, c := range ln {
			switch c {
			case rune('#'):
				g.Set(grid.Pos(x, y), grid.Lumberyard)
			case rune('|'):
				g.Set(grid.Pos(x, y), grid.Trees)
			}
		}
	}

	return g
}

func step(world *grid.Grid) *grid.Grid {
	updated := grid.New(world.Width, world.Height)

	for x := 0; x < world.Width; x++ {
		for y := 0; y < world.Height; y++ {
			p := grid.Pos(x, y)
			neighbours := countNeighbours(p, world)
			switch world.At(p) {
			case grid.Open:
				if neighbours[grid.Trees] >= 3 {
					updated.Set(p, grid.Trees)
				} else {
					updated.Set(p, grid.Open)
				}
			case grid.Trees:
				if neighbours[grid.Lumberyard] >= 3 {
					updated.Set(p, grid.Lumberyard)
				} else {
					updated.Set(p, grid.Trees)
				}
			case grid.Lumberyard:
				if neighbours[grid.Lumberyard] >= 1 && neighbours[grid.Trees] >= 1 {
					updated.Set(p, grid.Lumberyard)
				} else {
					updated.Set(p, grid.Open)
				}
			}
		}
	}
	return updated
}

func neighbours(p grid.Position) (n []grid.Position) {
	for y := p.Y - 1; y <= p.Y+1; y++ {
		for x := p.X - 1; x <= p.X+1; x++ {
			if !(x == p.X && y == p.Y) {
				n = append(n, grid.Pos(x, y))
			}
		}
	}
	return
}

func countNeighbours(pos grid.Position, world *grid.Grid) map[grid.Type]int {
	bounds := image.Rect(0, 0, world.Width, world.Height)
	result := map[grid.Type]int{grid.Open: 0, grid.Lumberyard: 0, grid.Trees: 0}
	for _, n := range neighbours(pos) {
		if n.In(bounds) {
			result[world.At(n)] += 1
		}
	}
	return result
}

func count(w *grid.Grid, p func(grid.Type) bool) (c int) {
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			if p(w.At(grid.Pos(x, y))) {
				c += 1
			}
		}
	}
	return
}

func main() {
	fileContents, err := ioutil.ReadFile("18_settlers_of_the_north_pole/day18_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	g := parseInput(string(fileContents))
	for i := 0; i < 10; i++ {
		g = step(g)
	}
	lumberyards := count(g, func(t grid.Type) bool { return t == grid.Lumberyard })
	trees := count(g, func(t grid.Type) bool { return t == grid.Trees })

	fmt.Printf("Day 18 part 1 result: %+v\n", lumberyards*trees)

	// Part 2
	g = parseInput(string(fileContents))
	hist := make(map[string]int)
	for i := 1; i < 1000000; i++ {
		g = step(g)
		if prev, ok := hist[g.Hash()]; ok {
			period := i - prev
			if (i % period) == (1000000000 % period) {
				break
			}
		} else {
			hist[g.Hash()] = i
		}
	}

	lumberyards = count(g, func(t grid.Type) bool { return t == grid.Lumberyard })
	trees = count(g, func(t grid.Type) bool { return t == grid.Trees })

	fmt.Printf("Day 18 part 2 result: %+v\n", lumberyards*trees)
}
