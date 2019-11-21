package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	wrappingPaper := 0
	ribbon := 0
	for _, p := range parsePresents(input) {
		wrappingPaper += p[0]*p[1] + p[0]*p[1]*2 + p[0]*p[2]*2 + p[1]*p[2]*2
		ribbon += p[0]*2 + p[1]*2 + p[0]*p[1]*p[2]
	}
	return strconv.Itoa(wrappingPaper), strconv.Itoa(ribbon)
}

func parsePresents(input string) [][]int {
	var presents [][]int
	for _, ln := range strings.Split(input, "\n") {
		b := []int{0, 0, 0}
		if _, err := fmt.Sscanf(ln, "%dx%dx%d", &b[0], &b[1], &b[2]); err == nil {
			sort.Ints(b)
			presents = append(presents, b)
		}
	}
	return presents
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
