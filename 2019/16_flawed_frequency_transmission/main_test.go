package main

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	for i, tt := range testPuzzles {
		t.Run(fmt.Sprint("example ", i), func(t *testing.T) {
			if tt.Solution1 != "" {
				got := part1(tt.Puzzle, tt.PhaseCount)
				if got != tt.Solution1 {
					t.Errorf("solve() part1 = %v, want %v", got, tt.Solution1)
				}
			}
			if tt.Solution2 != "" {
				got := part2(tt.Puzzle, tt.PhaseCount)
				if got != tt.Solution2 {
					t.Errorf("solve() part2 = %v, want %v", got, tt.Solution2)
				}
			}
		})
	}
}
