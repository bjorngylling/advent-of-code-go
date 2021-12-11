package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"
	"time"

	"github.com/bjorngylling/advent-of-code/util"
)

var adjacent = []image.Point{
	image.Pt(0, 1),
	image.Pt(1, 1),
	image.Pt(1, 0),
	image.Pt(1, -1),
	image.Pt(0, -1),
	image.Pt(-1, -1),
	image.Pt(-1, 0),
	image.Pt(-1, 1),
}

func step(octopi map[image.Point]int) int {
	flashpoints := map[image.Point]struct{}{}
	var flash func(p image.Point)
	flash = func(p image.Point) {
		if _, ok := octopi[p]; !ok {
			return
		}
		octopi[p]++
		if _, flashed := flashpoints[p]; !flashed && octopi[p] > 9 {
			flashpoints[p] = struct{}{}
			for _, adj := range adjacent {
				flash(p.Add(adj))
			}
		}
	}
	for pt := range octopi {
		flash(pt)
	}
	for pt := range flashpoints {
		octopi[pt] = 0
	}
	return len(flashpoints)
}

func solve(input string) (string, string) {
	octopi := map[image.Point]int{}
	for y, ln := range strings.Split(input, "\n") {
		for x, v := range []rune(ln) {
			octopi[image.Pt(x, y)] = util.GetInt(string(v))
		}
	}

	sum := 0
	allFlashStep := -1
	i := 0
	for ; i < 100; i++ {
		flashCount := step(octopi)
		sum += flashCount
		if allFlashStep == -1 && flashCount == len(octopi) {
			allFlashStep = i
		}
	}
	for ; allFlashStep == -1; i++ {
		flashCount := step(octopi)
		if allFlashStep == -1 && flashCount == len(octopi) {
			allFlashStep = i+1
		}
	}

	return strconv.Itoa(sum), strconv.Itoa(allFlashStep)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
