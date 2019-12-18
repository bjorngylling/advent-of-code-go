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

func fakeFft(signal []int, cycles int) []int {
	for ; cycles > 0; cycles-- {
		sum := 0
		for n := len(signal) - 1; n >= 0; n-- {
			sum += signal[n]
			signal[n] = util.Abs(sum) % 10
		}
	}
	return signal
}

func part1(input string, phaseCount int) string {
	in := parse(input)
	part1 := ""
	for _, i := range fft(in, phaseCount)[:8] {
		part1 += strconv.Itoa(i)
	}
	return part1
}

func part2(input string, phaseCount int) string {
	in := parse(input)
	offset := util.GetInt(input[:7])
	signal := make([]int, len(in)*10000-offset)
	for i := range signal {
		signal[i] = in[(offset+i)%len(in)]
	}
	part2 := ""
	for _, i := range fakeFft(signal, phaseCount)[:8] {
		part2 += strconv.Itoa(i)
	}

	return part2
}

func solve(input string, phaseCount int) (string, string) {
	return part1(input, phaseCount), part2(input, phaseCount)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle, 100)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
