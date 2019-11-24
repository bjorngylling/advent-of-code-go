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

func Test_parse(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{name: "simple", args: args{"abba[mnop]"}, want: []string{"abba"}, want1: []string{"mnop"}},
		{name: "two_regular", args: args{"abba[mnop]qrst"}, want: []string{"abba", "qrst"}, want1: []string{"mnop"}},
		{name: "two_hypernet", args: args{"abba[mnop][qrst]"}, want: []string{"abba"}, want1: []string{"mnop", "qrst"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.args.addr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_abba(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "abba", args: args{"abba"}, want: true},
		{name: "abba", args: args{"ioxxoj"}, want: true},
		{name: "not_abba", args: args{"mnop"}, want: false},
		{name: "not_abba", args: args{"aaaa"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abba(tt.args.seq); got != tt.want {
				t.Errorf("abba() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aba(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "simple", args: args{"aba"}, want: []string{"aba"}},
		{name: "two", args: args{"abaxulu"}, want: []string{"aba", "ulu"}},
		{name: "two_with_overlap", args: args{"zazbz"}, want: []string{"zaz", "zbz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aba(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aba() = %v, want %v", got, tt.want)
			}
		})
	}
}
