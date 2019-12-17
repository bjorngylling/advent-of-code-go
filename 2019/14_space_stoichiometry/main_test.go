package main

import (
	"fmt"
	"strings"
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

func Test_revTopologialSort(t *testing.T) {
	type args struct {
		n *node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "sorted", args: args{n: buildTree(testPuzzles[0].Puzzle)["FUEL"]}, want: "FUEL E D C B A ORE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := revTopologialSort(tt.args.n)
			var lst []string
			for _, n := range sorted {
				lst = append(lst, n.name)
			}
			got := strings.Join(lst, " ")
			if got != tt.want {
				t.Errorf("revTopologialSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
