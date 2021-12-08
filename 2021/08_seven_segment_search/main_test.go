package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_correctPattern(t *testing.T) {
	input := [10]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	want := map[rune]rune{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}
	got := correctPattern(input)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("correctPattern() mismatch (-want +got):\n%s", diff)
	}
}

func Test_solve(t *testing.T) {
	for i, tt := range testPuzzles {
		t.Run(fmt.Sprint("example ", i), func(t *testing.T) {
			part1, part2 := solve(tt.Puzzle)
			if tt.Solution1 != "" && part1 != tt.Solution1 {
				t.Errorf("solve() part1 = %v, want %v", part1, tt.Solution1)
			}
			if tt.Solution2 != "" && part2 != tt.Solution2 {
				t.Errorf("solve() part2 = %v, want %v", part2, tt.Solution2)
			}
		})
	}
}
