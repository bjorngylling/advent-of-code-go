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

func Test_validate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid", args: args{s: "111111"}, want: true},
		{name: "valid", args: args{s: "111112"}, want: true},
		{name: "valid", args: args{s: "112233"}, want: true},
		{name: "valid", args: args{s: "111122"}, want: true},
		{name: "decreasing", args: args{s: "223450"}, want: false},
		{name: "no_double", args: args{s: "123789"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.s); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid", args: args{s: "111111"}, want: false},
		{name: "valid", args: args{s: "111112"}, want: false},
		{name: "valid", args: args{s: "112233"}, want: true},
		{name: "valid", args: args{s: "111122"}, want: true},
		{name: "decreasing", args: args{s: "223450"}, want: false},
		{name: "no_double", args: args{s: "123789"}, want: false},
		{name: "larger_grp", args: args{s: "123444"}, want: false},
		{name: "larger_grp", args: args{s: "444123"}, want: false},
		{name: "larger_grp", args: args{s: "443323"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate2(tt.args.s); got != tt.want {
				t.Errorf("validate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
