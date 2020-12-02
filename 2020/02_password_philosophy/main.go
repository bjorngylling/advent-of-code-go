package main

import (
	"fmt"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	solution1 := 0
	solution2 := 0
	for _, ln := range strings.Split(input, "\n") {
		var min, max int
		var chr, pass string
		fmt.Sscanf(ln, "%d-%d %1s: %s", &min, &max, &chr, &pass)
		n := strings.Count(pass, chr)
		if n >= min && n <= max {
			solution1++
		}
		if (string(pass[min-1]) == chr && string(pass[max-1]) != chr) ||
			(string(pass[min-1]) != chr && string(pass[max-1]) == chr) {
			solution2++
		}
	}
	return fmt.Sprint(solution1), fmt.Sprint(solution2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
