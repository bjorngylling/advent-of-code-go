package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

func solve(input string, preambleSize int) (string, string) {
	var data []int
	for _, ln := range strings.Split(input, "\n") {
		data = append(data, util.GetInt(ln))
	}
	solution1 := firstInvalid(data, preambleSize)

	solution2 := 0
outer:
	for a := 0; a < len(data); a++ {
		for b := a; b < len(data); b++ {
			sum := 0
			for _, i := range data[a:b] {
				sum += i
			}
			if sum == solution1 {
				solution2 = util.Min(data[a:b]...) + util.Max(data[a:b]...)
				break outer
			}
		}
	}

	return strconv.Itoa(solution1), strconv.Itoa(solution2)
}

func firstInvalid(data []int, preambleSize int) int {
	for i := preambleSize; i < len(data); i++ {
		if !validate(data[i], data[i-preambleSize:i]) {
			return data[i]
		}
	}
	return 0
}

func validate(num int, preamble []int) bool {
	for _, i := range preamble {
		p := num - i
		for _, j := range preamble {
			if p == j && j != i {
				return true
			}
		}
	}
	return false
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle, 25)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
