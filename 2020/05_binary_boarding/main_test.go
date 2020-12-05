package main

import (
	"fmt"
	"testing"
)

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

func Test_calcSeatId(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: struct{ input string }{input: "FBFBBFFRLR"}, want: 357},
		{name: "2", args: struct{ input string }{input: "FFFBBBFRRR"}, want: 119},
		{name: "3", args: struct{ input string }{input: "BBFFBBFRLL"}, want: 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSeatId(tt.args.input); got != tt.want {
				t.Errorf("calcSeatId() = %v, want %v", got, tt.want)
			}
		})
	}
}
