package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"image"
	"math"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	space := make(map[image.Point]map[image.Point]struct{})
	for y, ln := range strings.Split(input, "\n") {
		for x, chr := range ln {
			if chr == '#' {
				space[image.Pt(x, y)] = make(map[image.Point]struct{})
			}
		}
	}

	lst := make([]image.Point, len(space))
	i := 0
	for a := range space {
		lst[i] = a
		i++
	}
	for _, pair := range pairs(lst) {
		blocked := false
		for c := range space {
			if pair[0] == c || pair[1] == c {
				continue
			}
			if intersects(pair[0], pair[1], c) {
				blocked = true
				break
			}
		}
		if !blocked {
			space[pair[0]][pair[1]] = struct{}{}
			space[pair[1]][pair[0]] = struct{}{}
		}
	}

	best := lst[0]
	for pt, inLOS := range space {
		if len(inLOS) > len(space[best]) {
			best = pt
		}
	}

	return strconv.Itoa(len(space[best])), ""
}

func pairs(lst []image.Point) [][2]image.Point {
	var pairs [][2]image.Point
	for i, a := range lst {
		for _, b := range lst[i+1:] {
			pairs = append(pairs, [2]image.Point{a, b})
		}
	}
	return pairs
}

// True if pt(c) is exactly on the line between pt(a) and pt(b)
func intersects(a, b, c image.Point) bool {
	return angleRad(a, b) == angleRad(a, c) && util.ManhattanDistance(a, c) < util.ManhattanDistance(a, b)
}

func angleRad(a, b image.Point) float64 {
	return math.Atan2(float64(b.Y-a.Y), float64(b.X-a.X))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
