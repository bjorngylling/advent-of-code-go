package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func calcSeatId(input string) int {
	rowMin, rowMax := 0, 128
	colMin, colMax := 0, 8
	for _, a := range input[0:7] {
		switch a {
		case 'B':
			rowMin += (rowMax - rowMin) / 2
			break
		case 'F':
			rowMax -= (rowMax - rowMin) / 2
			break
		}
	}
	for _, a := range input[7:] {
		switch a {
		case 'R':
			colMin += (colMax - colMin) / 2
			break
		case 'L':
			colMax -= (colMax - colMin) / 2
			break
		}
	}
	return rowMin*8 + colMin
}

func solve(input string) (string, string) {
	solution1 := 0
	var occupiedSeats []int
	for _, ln := range strings.Split(input, "\n") {
		id := calcSeatId(ln)
		occupiedSeats = append(occupiedSeats, id)
		if id > solution1 {
			solution1 = id
		}
	}
	var solution2 int
	sort.Ints(occupiedSeats)
	for i, id := range occupiedSeats {
		if i > 0 && id != occupiedSeats[i-1]+1 {
			solution2 = occupiedSeats[i-1] + 1
		}
	}
	return strconv.Itoa(solution1), strconv.Itoa(solution2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
