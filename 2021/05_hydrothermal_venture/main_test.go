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

func Test_plotLine(t *testing.T) {
	type args struct {
		start image.Point
		end   image.Point
	}
	tests := []struct {
		name string
		args args
		want []image.Point
	}{
		{
			name: "horizontal",
			args: args{start: image.Point{X: 0, Y: 0}, end: image.Point{X: 10, Y: 0}},
			want: []image.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 5, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0}, {X: 10, Y: 0}},
		},
		{
			name: "vertical",
			args: args{start: image.Point{X: 1, Y: 1}, end: image.Point{X: 1, Y: 6}},
			want: []image.Point{{X: 1, Y: 1}, {X: 1, Y: 2}, {X: 1, Y: 3}, {X: 1, Y: 4}, {X: 1, Y: 5}, {X: 1, Y: 6}},
		},
		{
			name: "diagonal",
			args: args{start: image.Point{X: 1, Y: 1}, end: image.Point{X: 6, Y: 6}},
			want: []image.Point{{X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 4}, {X: 5, Y: 5}, {X: 6, Y: 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plotLine(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plotLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
