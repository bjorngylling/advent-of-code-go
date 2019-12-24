package main

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	states := [][5 * 5]bool{{}}
	for y, ln := range strings.Split(input, "\n") {
		for x, c := range ln {
			states[0][5*y+x] = c == '#'
		}
	}

	adjacent := []image.Point{
		image.Pt(0, -1),
		image.Pt(-1, 0), image.Pt(1, 0),
		image.Pt(0, 1),
	}
	area := image.Rect(0, 0, 5, 5)
outer:
	for i := 0; ; i++ {
		newState := [5 * 5]bool{}
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				cur := image.Pt(x, y)
				adjBugCount := 0
				for _, d := range adjacent {
					p := cur.Add(d)
					if p.In(area) && states[i][5*p.Y+p.X] {
						adjBugCount++
					}
				}
				if states[i][5*cur.Y+cur.X] {
					newState[5*cur.Y+cur.X] = adjBugCount == 1
				} else {
					newState[5*cur.Y+cur.X] = adjBugCount == 1 || adjBugCount == 2
				}
			}
		}
		for _, s := range states {
			if s == newState {
				states = append(states, newState)
				break outer
			}
		}
		states = append(states, newState)
	}

	rating := 0
	for i, bug := range states[len(states)-1] {
		if bug {
			rating += int(math.Pow(2, float64(i)))
		}
	}

	return strconv.Itoa(rating), ""
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
