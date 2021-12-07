package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bjorngylling/advent-of-code/util"
)

func calcDist(pos []int, f func(int) int) int {
	max := util.Max(pos...)
	dist := make([][]int, max)
	for i := 0; i < max; i++ {
		dist[i] = make([]int, len(pos))
		for j, other := range pos {
			dist[i][j] = f(i - other)
		}
	}
	sums := make([]int, len(dist))
	for i, d := range dist {
		for _, n := range d {
			sums[i] += n
		}
	}
	return util.Min(sums...)
}

func partialSum(n int) int {
	n = util.Abs(n)
	return n * (n + 1) / 2
}

func solve(input string) (string, string) {
	var positions []int
	for _, n := range strings.Split(input, ",") {
		positions = append(positions, util.GetInt(n))
	}
	return strconv.Itoa(util.Min(calcDist(positions, util.Abs))),
		strconv.Itoa(util.Min(calcDist(positions, partialSum)))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
