package main

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"time"
)

type direction int

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

var delta = map[direction]image.Point{
	NORTH: image.Pt(0, 1),
	EAST:  image.Pt(1, 0),
	SOUTH: image.Pt(0, -1),
	WEST:  image.Pt(-1, 0),
}

func solve(input string) (string, string) {
	commands := strings.Split(input, ", ")
	pos := image.Pt(0, 0)
	dir := NORTH
	visited := map[image.Point]struct{}{pos: {}}
	part2 := ""
	for _, c := range commands {
		dir = turn(rune(c[0]), dir)
		steps, err := strconv.Atoi(c[1:])
		if err != nil {
			panic(err)
		}
		for i := 0; i < steps; i++ {
			pos = pos.Add(delta[dir])
			if _, found := visited[pos]; found && part2 == "" {
				part2 = distance(pos)
			}
			visited[pos] = struct{}{}
		}
	}
	return distance(pos), part2
}

func distance(pos image.Point) string {
	return fmt.Sprint(math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y)))
}

func turn(turn rune, dir direction) direction {
	switch turn {
	case 'R':
		dir += 1
	case 'L':
		dir -= 1
	}
	switch dir {
	case -1:
		dir = WEST
	case 4:
		dir = NORTH
	}
	return dir
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
