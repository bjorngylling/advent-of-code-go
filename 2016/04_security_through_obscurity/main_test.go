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

func Test_checksum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{"aaaaa-bbb-z-y-x"}, "abxyz"},
		{"tie", args{"a-b-c-d-e-f-g-h"}, "abcde"},
		{"words", args{"not-a-real-room"}, "oarel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksum(tt.args.s); got != tt.want {
				t.Errorf("checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decrypt(t *testing.T) {
	type args struct {
		name     string
		sectorId int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"qzmt-zixmtkozy-ivhz", 343}, "very encrypted name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.name, tt.args.sectorId); got != tt.want {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
