package main

import (
	"fmt"
	"image"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	seats := map[image.Point]struct{}{}
	split := strings.Split(input, "\n")
	width, height := len(split[0]), len(split)
	for y, ln := range split {
		for x, c := range ln {
			if c == 'L' {
				seats[image.Pt(x, y)] = struct{}{}
			}
		}
	}

	adjacent := [8]image.Point{
		image.Pt(-1, -1), image.Pt(0, -1), image.Pt(1, -1),
		image.Pt(-1, 0), image.Pt(1, 0),
		image.Pt(-1, 1), image.Pt(0, 1), image.Pt(1, 1),
	}
	simpleAdj := func(pos image.Point, world map[image.Point]struct{}) int {
		adjCount := 0
		for _, adj := range adjacent {
			if _, ok := world[pos.Add(adj)]; ok {
				adjCount++
			}
		}
		return adjCount
	}
	losAdj := func(pos image.Point, world map[image.Point]struct{}) int {
		adjCount := 0
		for _, adj := range adjacent {
			p := pos
			for {
				p = p.Add(adj)
				if p.X < 0 || p.X > width || p.Y < 0 || p.Y > height {
					break
				}
				if _, ok := seats[p]; ok {
					if _, ok := world[p]; ok {
						adjCount++
					}
					break
				}
			}
		}
		return adjCount
	}

	return strconv.Itoa(run(seats, simpleAdj, 4)), strconv.Itoa(run(seats, losAdj, 5))
}

func run(seats map[image.Point]struct{}, countAdj func(image.Point, map[image.Point]struct{}) int, adjLimit int) int {
	var states []map[image.Point]struct{}
	for i := 0; ; i++ {
		prevState := map[image.Point]struct{}{}
		if len(states) > 0 {
			prevState = states[len(states)-1]
		}
		newState := map[image.Point]struct{}{}
		for pos := range seats {
			adjCount := countAdj(pos, prevState)
			if _, ok := prevState[pos]; !ok && adjCount == 0 {
				newState[pos] = struct{}{}
			} else if ok && adjCount < adjLimit {
				newState[pos] = struct{}{}
			}
		}
		if reflect.DeepEqual(prevState, newState) {
			break
		}
		states = append(states, newState)
	}
	return len(states[len(states)-1])
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
