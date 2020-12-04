package main

import (
	"fmt"
	"image"
	"strings"
	"time"
)

type World struct {
	m map[image.Point]bool
	w int
	h int
}

func (w *World) get(x, y int) bool {
	return w.m[image.Pt(x%w.w, y)]
}

func solve(input string) (string, string) {
	world := &World{m: map[image.Point]bool{}}
	lines := strings.Split(input, "\n")
	world.w = len(lines[0])
	world.h = len(lines)
	for y, ln := range lines {
		for x, c := range ln {
			b := c == '#'
			world.m[image.Point{X: x, Y: y}] = b
		}
	}

	solution1 := 0
	for x, y := 0, 0; y < world.h; x, y = x+3, y+1 {
		if world.get(x, y) {
			solution1++
		}
	}

	solution2 := 1
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, s := range slopes {
		sum := 0
		for x, y := 0, 0; y < world.h; x, y = x+s[0], y+s[1] {
			if world.get(x, y) {
				sum++
			}
		}
		solution2 *= sum
	}

	return fmt.Sprint(solution1), fmt.Sprint(solution2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
