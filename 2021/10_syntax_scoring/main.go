package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	start := func(r rune) bool {
		return r == '(' || r == '{' || r == '[' || r == '<'
	}
	pair := func(r rune, other rune) bool {
		return (r == '(' && other == ')') || (r == '{' && other == '}') || (r == '[' && other == ']') || (r == '<' && other == '>')
	}

	var invalid []rune
	var completion [][]rune
	for _, ln := range strings.Split(input, "\n") {
		var s []rune
		corrupt := false
		for _, r := range []rune(ln) {
			if start(r) {
				s = append(s, r)
			} else {
				other := s[len(s)-1]
				s = s[:len(s)-1]
				if !pair(other, r) {
					invalid = append(invalid, r)
					corrupt = true
					break
				}
			}
		}
		if !corrupt {
			var c []rune
			for i := len(s) - 1; i >= 0; i-- {
				switch s[i] {
				case '(':
					c = append(c, ')')
				case '[':
					c = append(c, ']')
				case '{':
					c = append(c, '}')
				case '<':
					c = append(c, '>')
				}
			}
			completion = append(completion, c)
		}
	}
	p1 := 0
	for _, r := range invalid {
		switch r {
		case ')':
			p1 += 3
		case ']':
			p1 += 57
		case '}':
			p1 += 1197
		case '>':
			p1 += 25137
		}
	}
	var lineScore []int
	for _, l := range completion {
		score := 0
		for _, r := range l {
			score *= 5
			switch r {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}
		lineScore = append(lineScore, score)
	}
	sort.Ints(lineScore)
	return strconv.Itoa(p1), strconv.Itoa(lineScore[len(lineScore)/2])
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
