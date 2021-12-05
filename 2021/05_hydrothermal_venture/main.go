package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"image"
	"strconv"
	"strings"
	"time"
)

type line struct {
	start image.Point
	end   image.Point
}

func plotLine(start, end image.Point) []image.Point {
	dx := util.Abs(end.X - start.X)
	var sx, sy int
	if start.X < end.X {
		sx = 1
	} else {
		sx = -1
	}
	dy := -util.Abs(end.Y - start.Y)
	if start.Y < end.Y {
		sy = 1
	} else {
		sy = -1
	}
	err := dx + dy
	var result []image.Point
	for {
		result = append(result, start)
		if start.X == end.X && start.Y == end.Y {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			start.X += sx
		}
		if e2 <= dx {
			err += dx
			start.Y += sy
		}
	}
	return result
}

func findCollisions(lines []line, avoidDiagonal bool) int {
	collisionCount := map[image.Point]int{}
	for i, ln := range lines {
		if avoidDiagonal && !(ln.start.X == ln.end.X || ln.start.Y == ln.end.Y) {
			continue
		}
		plot := plotLine(ln.start, ln.end)
		for _, other := range lines[i+1:] {

			plotOther := plotLine(other.start, other.end)
			for _, p := range plot {
				for _, op := range plotOther {
					if p == op {
						if avoidDiagonal && !(other.start.X == other.end.X || other.start.Y == other.end.Y) {
							continue
						}
						collisionCount[p]++
					}
				}
			}
		}
	}
	return len(collisionCount)
}

func solve(input string) (string, string) {
	var lines []line
	for _, ln := range strings.Split(input, "\n") {
		line := line{start: image.Point{}, end: image.Point{}}
		fmt.Sscanf(ln, "%d,%d -> %d,%d", &line.start.X, &line.start.Y, &line.end.X, &line.end.Y)
		lines = append(lines, line)
	}

	return strconv.Itoa(findCollisions(lines, true)), strconv.Itoa(findCollisions(lines, false))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
