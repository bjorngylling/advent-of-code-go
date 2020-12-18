package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"math"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	spl := strings.Split(input, "\n")
	atStop := util.GetInt(spl[0])
	var busLines []int
	var intervals []int
	for i, b := range strings.Split(spl[1], ",") {
		if b != "x" {
			intervals = append(intervals, i)
			busLines = append(busLines, util.GetInt(b))
		}
	}

	return strconv.Itoa(part1(atStop, busLines)), strconv.Itoa(part2(busLines, intervals))
}

func part1(atStop int, busLines []int) int {
	solution1 := 0
	best := math.MaxInt32
	for _, b := range busLines {
		dep := (atStop/b + 1) * b
		if dep-atStop < best {
			solution1 = b
			best = dep - atStop
		}
	}
	return solution1 * best
}

func part2(busLines []int, intervals []int) int {
	step := busLines[0]
	it := 1
	for t := step; ; t += step {
		success := true
		for i := 0; i <= it; i++ {
			if (t+intervals[i])%busLines[i] != 0 {
				success = false
				break
			}
		}
		if success && it == len(busLines)-1 {
			return t
		}
		if success {
			step *= busLines[it]
			it++
		}
	}
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
