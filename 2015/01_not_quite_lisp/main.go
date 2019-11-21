package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	current := 0
	basementIteration := 0
	for i, c := range input {
		switch c {
		case '(':
			current++
		case ')':
			current--
		}
		if current < 0 {
			basementIteration = i
			break
		}
	}
	return strconv.Itoa(strings.Count(input, "(") - strings.Count(input, ")")), strconv.Itoa(basementIteration+1)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
