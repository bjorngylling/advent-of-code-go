package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

func part1(scan []int) int {
	incr := 0
	for i, ln := range scan {
		if i > 0 && ln > scan[i-1] {
			incr++
		}
	}
	return incr
}

func part2(scan []int) int {
	incr := 0
	for i := 0; i < len(scan); i++ {
		if i > 2 && scan[i] > scan[i-3] {
			incr++
		}
	}
	return incr
}

func solve(input string) (string, string) {
	lns := strings.Split(input, "\n")
	var scan []int
	for _, ln := range lns {
		scan = append(scan, util.GetInt(ln))
	}

	return strconv.Itoa(part1(scan)), strconv.Itoa(part2(scan))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
