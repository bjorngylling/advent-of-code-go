package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"image"
	"math"
	"sort"
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
	part1 := len(space[best])

	if len(space) < 200 {
		return strconv.Itoa(part1), "" // part two cant be done if the asteroid field is smaller than 200
	}

	// Calculate the slope of all asteroids from the station, if there's more than one asteroid on the same slope, sort them by distance.
	targets := map[float64][]image.Point{}
	for tar := range space {
		slope := angleRad(best, tar)
		if _, found := targets[slope]; found {
			targets[slope] = append(targets[slope], tar)
			sort.Slice(targets[slope], func(i, j int) bool {
				return util.ManhattanDistance(best, targets[slope][i]) < util.ManhattanDistance(best, targets[slope][j])
			})
		} else {
			targets[slope] = []image.Point{tar}
		}
	}

	// Get a list of all slopes and sort it in order they will be targeted, the cannon will start aiming at slope atan2(-1, 0) since the
	// coordinate system is "inverted", ie y=1 is south of y=0. 🤯
	allSlopes := make([]float64, len(targets))
	i = 0
	for slope := range targets {
		allSlopes[i] = slope
		i++
	}
	sort.Float64s(allSlopes)
	northPos := -1
	for i, slope := range allSlopes {
		if slope == math.Atan2(-1, 0) {
			northPos = i
			break
		}
	}
	allSlopes = append(allSlopes[northPos:], allSlopes[:northPos]...)

	// Spin the cannon and remove asteroids until 200 of them have been removed!
	var lastDestroy image.Point
outer:
	for destroyed := 0; ; {
		for _, slope := range allSlopes {
			if destroyed >= 200 {
				break outer
			}
			if len(targets[slope]) > 0 {
				lastDestroy = targets[slope][0]
				targets[slope] = targets[slope][1:]
				destroyed++
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(lastDestroy.X*100 + lastDestroy.Y)
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
