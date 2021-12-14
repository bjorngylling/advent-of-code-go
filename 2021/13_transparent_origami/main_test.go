package main

import (
	"fmt"
	"image"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_fold(t *testing.T) {
	got := fold(map[image.Point]struct{}{image.Pt(0, 14): {}, image.Pt(0, 9): {}}, image.Pt(0, 7))
	want := map[image.Point]struct{}{image.Pt(0, 0): {}, image.Pt(0, 5): {}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("fold() mismatch (-want +got):\n%s", diff)
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
