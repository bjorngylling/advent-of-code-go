package main

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	for i, tt := range testPuzzles {
		t.Run(fmt.Sprint("example ", i), func(t *testing.T) {
			part1, part2 := solve(tt.Puzzle)
			if part1 != tt.Solution1 {
				t.Errorf("solve() part1 = %v, want %v", part1, tt.Solution1)
			}
			if part2 != tt.Solution2 {
				t.Errorf("solve() part2 = %v, want %v", part2, tt.Solution2)
			}
		})
	}
}

func Test_turn(t *testing.T) {
	type args struct {
		turn rune
		dir  direction
	}
	tests := []struct {
		name string
		args args
		want direction
	}{
		{name: "turn_right_from_north", args: args{'R', NORTH}, want: EAST},
		{name: "turn_right_from_west", args: args{'R', WEST}, want: NORTH},
		{name: "turn_left_from_east", args: args{'L', EAST}, want: NORTH},
		{name: "turn_left_from_north", args: args{'L', NORTH}, want: WEST},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := turn(tt.args.turn, tt.args.dir); got != tt.want {
				t.Errorf("turn() = %v, want %v", got, tt.want)
			}
		})
	}
}
