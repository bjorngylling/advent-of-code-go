package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	return strconv.Itoa(part1(input)), strconv.Itoa(part2(input))
}

func part1(input string) int {
	split := strings.Split(input, "\n")
	t := map[string]struct{}{}
	var groups []map[string]struct{}
	for i, ln := range split {
		if i > 0 && split[i] == "" {
			groups = append(groups, t)
			t = map[string]struct{}{}
			continue
		}
		for _, c := range ln {
			t[string(c)] = struct{}{}
		}
	}
	groups = append(groups, t)

	solution1 := 0
	for _, grp := range groups {
		solution1 += len(grp)
	}
	return solution1
}

func part2(input string) int {
	split := strings.Split(input, "\n\n")
	groups := make([][]string, len(split))
	for i, grp := range split {
		for _, ln := range strings.Split(grp, "\n") {
			groups[i] = append(groups[i], ln)
		}
	}

	solution := 0
	for _, grp := range groups {
		solution += len(union(grp))
	}
	return solution
}

func union(str []string) string {
	result := str[0]
	for _, s := range str[1:] {
		for _, c := range result {
			if !strings.Contains(s, string(c)) {
				result = strings.ReplaceAll(result, string(c), "")
			}
		}
	}
	return result
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
