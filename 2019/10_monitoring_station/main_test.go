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

func Test_pairs(t *testing.T) {
	type args struct {
		lst []image.Point
	}
	tests := []struct {
		name string
		args args
		want [][2]image.Point
	}{
		{
			name: "a_pair",
			args: args{[]image.Point{image.Pt(1, 1), image.Pt(2, 2)}},
			want: [][2]image.Point{{image.Pt(1, 1), image.Pt(2, 2)}},
		},
		{
			name: "three_pairs",
			args: args{[]image.Point{image.Pt(1, 1), image.Pt(2, 2), image.Pt(3, 3)}},
			want: [][2]image.Point{
				{image.Pt(1, 1), image.Pt(2, 2)},
				{image.Pt(1, 1), image.Pt(3, 3)},
				{image.Pt(2, 2), image.Pt(3, 3)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairs(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersects(t *testing.T) {
	type args struct {
		a image.Point
		b image.Point
		c image.Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "intersection", args: args{a: image.Pt(0, 0), b: image.Pt(0, 10), c: image.Pt(0, 5)}, want: true},
		{name: "intersection", args: args{a: image.Pt(0, 0), b: image.Pt(10, 10), c: image.Pt(5, 5)}, want: true},
		{name: "no_intersection", args: args{a: image.Pt(0, 0), b: image.Pt(0, 10), c: image.Pt(1, 5)}, want: false},
		{name: "no_intersection", args: args{a: image.Pt(0, 0), b: image.Pt(10, 10), c: image.Pt(6, 5)}, want: false},
		{name: "negative", args: args{a: image.Pt(10, 0), b: image.Pt(0, 0), c: image.Pt(3, 0)}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersects(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}
