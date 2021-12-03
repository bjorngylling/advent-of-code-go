package main

import (
	"fmt"
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

func Test_filter(t *testing.T) {
	type args struct {
		l   [][]int
		f   int
		pos int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "basic",
			args: args{
				l:   [][]int{{0, 0, 1, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 1, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 1, 0, 1}, {0, 1, 1, 1, 1}, {0, 0, 1, 1, 1}, {1, 1, 1, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 1, 0, 1, 0}},
				f:   1,
				pos: 0},
			want: [][]int{{1, 1, 1, 1, 0}, {1, 0, 1, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args.l, tt.args.f, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
