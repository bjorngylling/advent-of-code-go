package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	var nums []int
	for _, n := range strings.Split(input, "\n") {
		if i, err := strconv.Atoi(n); err == nil {
			nums = append(nums, i)
		}
	}

	var first int
	for _, n := range nums {
		for _, m := range nums {
			if n+m == 2020 {
				first = n * m
			}
		}
	}

	var second int
	for _, n := range nums {
		for _, m := range nums {
			for _, o := range nums {
				if n+m+o == 2020 {
					second = n * m * o
				}
			}
		}
	}

	return fmt.Sprint(first), fmt.Sprint(second)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
