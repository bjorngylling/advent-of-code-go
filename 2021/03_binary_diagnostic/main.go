package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

func mostCommon(l [][]int, pos int) int {
	m := [2]int{}
	for _, v := range l {
		m[v[pos]]++
	}
	if m[0] > m[1] {
		return 0
	}

	return 1
}

func filter(l [][]int, f int, pos int) [][]int {
	var result [][]int
	for _, v := range l {
		if v[pos] != f {
			continue
		}
		result = append(result, v)
	}
	return result
}

func part1(nums [][]int) int {
	gamma := 0b0
	epsilon := 0b0
	for i := range nums[0] {
		c := mostCommon(nums, i)
		gamma <<= 1
		gamma += c
		epsilon <<= 1
		epsilon += c ^ 1
	}
	return gamma * epsilon
}

func part2(input [][]int, f func(int) int) int {
	for i := range input {
		c := mostCommon(input, i)
		input = filter(input, f(c), i)
		if len(input) == 1 {
			break
		}
	}
	r := 0b0
	for _, v := range input[0] {
		r <<= 1
		r += v
	}
	return r
}

func solve(input string) (string, string) {
	lns := strings.Split(input, "\n")
	nums := make([][]int, len(lns))
	for i, ln := range lns {
		for _, c := range ln {
			nums[i] = append(nums[i], util.GetInt(string(c)))
		}
	}

	oxygen := part2(nums, func(c int) int {
		return c
	})
	c02 := part2(nums, func(c int) int {
		return c ^ 1
	})

	return strconv.Itoa(part1(nums)), strconv.Itoa(oxygen * c02)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
