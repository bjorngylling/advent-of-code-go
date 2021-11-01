package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestParseMask(t *testing.T) {
	testCases := []struct {
		desc       string
		mask       string
		wantIgnore int64
		wantMask   int64
	}{
        {desc: "", mask: "", wantIgnore: 0, wantMask: 0},
		{desc: "", mask: "1XXXX0X", wantIgnore: 2, wantMask: 64},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			gotIgnore, gotMask := parseMask(tC.mask)
			if gotIgnore != tC.wantIgnore {
                t.Errorf("parseMask: got ignore=%b but want=%b", gotIgnore, tC.wantIgnore)
			}
			if gotMask != tC.wantMask {
                t.Errorf("parseMask: got mask=%b but want=%b", gotMask, tC.wantMask)
			}
		})
	}
}

func TestApplyMask(t *testing.T) {
	testCases := []struct {
		desc string
		mask string
		val  int64
		want int64
	}{
		{desc: "empty", mask: "", val: 0, want: 0},
		{desc: "basic", mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", val: 11, want: 73},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := applyMask(tC.mask, tC.val)
			if got != tC.want {
				t.Errorf("applyMask: got %d but want %d", got, tC.want)
			}
		})
	}
}

func TestParseMaskPart2(t *testing.T) {
	testCases := []struct {
		desc       string
		mask       string
		wantIgnore int64
		wantMask   int64
	}{
		{desc: "", mask: "", wantIgnore: 0, wantMask: 0},
		{desc: "", mask: "000000000000000000000000000000X1001X", wantIgnore: 33, wantMask: 51},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			gotIgnore, gotMask := parseMaskPart2(tC.mask)
			if gotIgnore != tC.wantIgnore {
				t.Errorf("parseMaskPart2: got ignore=%b but want=%b", gotIgnore, tC.wantIgnore)
			}
			if gotMask != tC.wantMask {
				t.Errorf("parseMaskPart2: got mask=%b but want=%b", gotMask, tC.wantMask)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(testPuzzles[1].Puzzle, "\n"))

	want := map[int64]int64 {
		16: 1,
		17: 1,
		18: 1,
		19: 1,
		24: 1,
		25: 1,
		26: 1,
		27: 1,
		58: 100,
		59: 100,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("part2() mismatch (-want +got):\n%s", diff)
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
