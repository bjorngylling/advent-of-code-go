package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bjorngylling/advent-of-code/util"
)

func part1(input [][4]string) int {
	count := make([]int, 10)
	for _, l := range input {
		for _, s := range l {
			count[len(s)]++
		}
	}
	return count[2] + count[4] + count[3] + count[7]
}

func part2(patterns [][10]string, values [][4]string) int {
	sum := 0
	for i, p := range patterns {
		mapping := correctPattern(p)
		val := ""
		for _, v := range values[i] {
			patt := ""
			for _, c := range []rune(v) {
				patt += string(mapping[c])
			}
			val += patternToNum(patt)
		}
		sum += util.GetInt(val)
	}

	return sum
}

func patternToNum(pattern string) string {
	pattToNum := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}
	s := []rune(pattern)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return pattToNum[string(s)]
}

func correctPattern(patterns [10]string) map[rune]rune {
	mapping := map[rune]rune{}
	countToPatt := map[int][]map[rune]struct{}{}
	for _, s := range patterns {
		set := map[rune]struct{}{}
		for _, c := range []rune(s) {
			set[c] = struct{}{}
		}
		countToPatt[len(s)] = append(countToPatt[len(s)], set)
	}
	// line a
	for c := range countToPatt[3][0] {
		if _, ok := countToPatt[2][0][c]; !ok {
			mapping[c] = 'a'
		}
	}
	// line e
	count := map[rune]int{}
	for _, p := range countToPatt[5] {
		for c := range p {
			if _, ok := countToPatt[4][0][c]; !ok {
				count[c]++
			}
		}
	}
	for k, v := range count {
		if v == 1 {
			mapping[k] = 'e'
		}
	}
	// line g
	for c := range countToPatt[7][0] {
		if _, ok := countToPatt[4][0][c]; !ok {
			if _, ok := mapping[c]; !ok {
				mapping[c] = 'g'
			}
		}
	}
	// line d
	count = map[rune]int{}
	for _, p := range countToPatt[6] {
		for c := range p {
			// not known or in pattern 7
			if _, ok := mapping[c]; !ok {
				if _, ok := countToPatt[3][0][c]; !ok {
					count[c]++
				}
			}
		}
	}
	for k, v := range count {
		if v == 2 {
			mapping[k] = 'd'
		}
	}
	// line b
	for _, p := range countToPatt[5] {
		for c := range p {
			// not known or in pattern 7
			if _, ok := mapping[c]; !ok {
				if _, ok := countToPatt[3][0][c]; !ok {
					mapping[c] = 'b'
				}
			}
		}
	}
	// line c and f
	count = map[rune]int{}
	for _, p := range countToPatt[6] {
		for c := range p {
			if _, ok := mapping[c]; !ok {
				count[c]++
			}
		}
	}
	for k, v := range count {
		if v == 2 {
			mapping[k] = 'c'
		} else if v == 3 {
			mapping[k] = 'f'
		}
	}

	return mapping
}

func solve(input string) (string, string) {
	patterns := [][10]string{}
	values := [][4]string{}
	for _, ln := range strings.Split(input, "\n") {
		spl := strings.Split(ln, "|")
		var p [10]string
		var v [4]string
		for i, s := range strings.Fields(spl[0]) {
			p[i] = s
		}
		for i, s := range strings.Fields(spl[1]) {
			v[i] = s
		}
		patterns = append(patterns, p)
		values = append(values, v)
	}

	return strconv.Itoa(part1(values)), strconv.Itoa(part2(patterns, values))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
