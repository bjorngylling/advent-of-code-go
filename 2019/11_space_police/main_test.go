package main

import (
	"fmt"
	"image"
	"reflect"
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

func Test_turn(t *testing.T) {
	type args struct {
		dir  image.Point
		turn int
	}
	var (
		north = image.Pt(0, 1)
		east  = image.Pt(1, 0)
		south = image.Pt(0, -1)
		west  = image.Pt(-1, 0)
	)
	tests := []struct {
		name string
		args args
		want image.Point
	}{
		{name: "north_right", args: args{dir: north, turn: right}, want: east},
		{name: "north_left", args: args{dir: north, turn: left}, want: west},
		{name: "east_right", args: args{dir: east, turn: right}, want: south},
		{name: "east_left", args: args{dir: east, turn: left}, want: north},
		{name: "south_right", args: args{dir: south, turn: right}, want: west},
		{name: "south_left", args: args{dir: south, turn: left}, want: east},
		{name: "west_right", args: args{dir: west, turn: right}, want: north},
		{name: "west_left", args: args{dir: west, turn: left}, want: south},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := turn(tt.args.dir, tt.args.turn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("turn() = %v, want %v", got, tt.want)
			}
		})
	}
}
