package main

import (
	"fmt"
	"strconv"
	"time"
)

func solve(input string) (string, string) {
	var cur, end int
	fmt.Sscanf(input, "%d-%d", &cur, &end)

	part1, part2 := 0, 0
	for ; cur <= end; cur++ {
		if validate(strconv.Itoa(cur)) {
			part1++
		}
		if validate2(strconv.Itoa(cur)) {
			part2++
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func validate(s string) bool {
	var ints [6]int
	for i := 0; i < 6; i++ {
		ints[i], _ = strconv.Atoi(string(s[i]))
	}
	pairs := false
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] > ints[i+1] {
			return false
		}
		pairs = pairs || ints[i] == ints[i+1]
	}
	return pairs
}

func validate2(s string) bool {
	var ints [6]int
	for i := 0; i < 6; i++ {
		ints[i], _ = strconv.Atoi(string(s[i]))
	}
	pairs := false
	for i := 0; i < len(ints)-1; {
		if ints[i] > ints[i+1] {
			return false
		}
		grp := 1
		for j := i; j < 5 && ints[j] == ints[j+1]; j++ {
			grp++
		}
		i += grp
		if grp > 1 {
			i--
		}
		pairs = pairs || grp == 2
	}
	return pairs
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
