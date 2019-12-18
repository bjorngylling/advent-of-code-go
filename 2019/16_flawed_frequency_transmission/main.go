package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"time"
)

func parse(input string) []int {
	var result []int
	for _, c := range input {
		result = append(result, int(c)-48)
	}
	return result
}

func fft(signal []int, cycles int) []int {
	base := []int{0, 1, 0, -1}
	for phase := 0; phase < cycles; phase++ {
		out := make([]int, len(signal))
		for n := 1; n <= len(signal); n++ {
			sum := 0
			for i := 1; i <= len(signal); i++ {
				sum += signal[i-1] * base[i/n%4]
			}
			out[n-1] = util.Abs(sum) % 10
		}
		signal = out
	}
	return signal
}

func solve(input string, phaseCount int) (string, string) {
	signal := fft(parse(input), phaseCount)
	res := ""
	for _, i := range signal[:8] {
		res += strconv.Itoa(i)
	}
	return res, ""
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle, 100)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
