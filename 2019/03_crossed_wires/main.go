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

var vel = map[rune]image.Point{
	'U': image.Pt(0, 1),
	'R': image.Pt(1, 0),
	'D': image.Pt(0, -1),
	'L': image.Pt(-1, 0),
}

type pointSet = map[image.Point]int

func solve(input string) (string, string) {
	var cables []pointSet
	for _, instructions := range strings.Split(input, "\n") {
		cable := make(pointSet)
		pos := image.Point{}
		step := 1
		for _, instr := range strings.Split(instructions, ",") {
			dir := instr[0]
			d, _ := strconv.Atoi(instr[1:])
			for i := 0; i < d; i++ {
				pos = pos.Add(vel[rune(dir)])
				if _, found := cable[pos]; !found {
					cable[pos] = step
				}
				step++
			}
		}
		cables = append(cables, cable)
	}

	part1 := math.MaxInt32
	part2 := math.MaxInt32
	for pt, steps := range union(cables[0], cables[1]) {
		if d := util.ManhattanDistance(image.Point{}, pt); d < part1 {
			part1 = d
		}
		if steps < part2 {
			part2 = steps
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func union(a, b pointSet) pointSet {
	res := make(pointSet)
	for pt := range a {
		if _, found := b[pt]; found {
			res[pt] = a[pt] + b[pt]
		}
	}
	return res
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
