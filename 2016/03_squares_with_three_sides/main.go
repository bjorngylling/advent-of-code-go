package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	part1 := 0
	part2 := 0

	var columnTriangles [3][3]int
	for lineNumber, ln := range strings.Split(input, "\n") {
		ln = strings.TrimSpace(ln)
		var sides [3]int
		i := 0
		for _, s := range strings.Split(ln, " ") {
			if side, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
				sides[i] = side
				i++
			}
		}
		if valid(sides) {
			part1++
		}
		for i := 0; i < 3; i++ {
			columnTriangles[i][lineNumber%3] = sides[i]
		}
		if (lineNumber+1)%3 == 0 {
			for i := 0; i < 3; i++ {
				if valid(columnTriangles[i]) {
					part2++
				}
			}
		}
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func valid(sides [3]int) bool {
	return sides[0]+sides[1] > sides[2] &&
		sides[0]+sides[2] > sides[1] &&
		sides[1]+sides[2] > sides[0]
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
