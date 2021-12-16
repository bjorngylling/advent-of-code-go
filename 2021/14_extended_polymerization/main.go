package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func f(rules map[string]string, pairs map[string]int) map[string]int {
	next := map[string]int{}
	for k, v := range pairs {
		next[string(k[0])+rules[k]] += v
		next[rules[k]+string(k[1])] += v
	}
	return next
}

func mostCommonLeastCommon(pairs map[string]int, last string) int {
	count := map[string]int{}
	for k, v := range pairs {
		count[string(k[0])] += v
	}
	count[last]++

	min, max := math.MaxInt, 0
	for _, v := range count {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func solve(input string) (string, string) {
	spl := strings.Split(input, "\n\n")
	template := spl[0]
	rules := map[string]string{}
	for _, ln := range strings.Split(spl[1], "\n") {
		rules[ln[0:2]] = string(ln[6])
	}

	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}
	for step := 0; step < 10; step++ {
		pairs = f(rules, pairs)
	}
	p1 := mostCommonLeastCommon(pairs, string(template[len(template)-1]))
	for step := 0; step < 30; step++ {
		pairs = f(rules, pairs)
	}

	return strconv.Itoa(p1), strconv.Itoa(mostCommonLeastCommon(pairs, string(template[len(template)-1])))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
