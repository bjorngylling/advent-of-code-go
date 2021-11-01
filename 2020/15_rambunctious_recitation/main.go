package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	spl := strings.Split(input, ",")
	var initial []int
	for _, s := range spl {
		d, _ := strconv.Atoi(s)
		initial = append(initial, d)
	}

	f := func(iterations int) int {
		curr := make([]int, iterations, iterations)
		prev := make([]int, iterations, iterations)
		for i, d := range initial {
			curr[d] = i + 1
			prev[d] = -1
		}
		var last int
		for i := len(initial); i < iterations; i++ {
			if curr[last] == 0 || prev[last] == -1 {
				last = 0
			} else {
				last = curr[last] - prev[last]
			}
			if curr[last] == 0 {
				prev[last] = -1
				curr[last] = i + 1
			} else {
				prev[last] = curr[last]
				curr[last] = i + 1
			}
		}
		return last
	}


	return strconv.Itoa(f(2020)), strconv.Itoa(f(30000000))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
