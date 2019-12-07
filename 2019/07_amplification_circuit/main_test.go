package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		reg memory
	}
	tests := []struct {
		name string
		args args
		want memory
	}{
		{name: "example_01", args: args{reg: memory{1, 0, 0, 0, 99}}, want: memory{2, 0, 0, 0, 99}},
		{name: "example_02", args: args{reg: memory{2, 3, 0, 3, 99}}, want: memory{2, 3, 0, 6, 99}},
		{name: "example_03", args: args{reg: memory{2, 4, 4, 5, 99, 0}}, want: memory{2, 4, 4, 5, 99, 9801}},
		{name: "example_04", args: args{reg: memory{1, 1, 1, 4, 99, 5, 6, 0, 99}}, want: memory{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{name: "example_param_mode", args: args{reg: memory{1002, 4, 3, 4, 33}}, want: memory{1002, 4, 3, 4, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.reg, nil, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run_with_io(t *testing.T) {
	type args struct {
		reg memory
	}
	tests := []struct {
		name  string
		args  args
		input []int
		want  int
	}{
		{name: "simple_io", args: args{reg: memory{3, 0, 4, 0, 99}}, input: []int{5}, want: 5},

		{name: "pos_mode_eq_8", args: args{reg: memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}}, input: []int{5}, want: 0},
		{name: "pos_mode_eq_8", args: args{reg: memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}}, input: []int{8}, want: 1},
		{name: "pos_mode_lt_8", args: args{reg: memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}}, input: []int{5}, want: 1},
		{name: "pos_mode_lt_8", args: args{reg: memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}}, input: []int{8}, want: 0},
		{name: "pos_mode_lt_8", args: args{reg: memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}}, input: []int{42}, want: 0},

		{name: "intermediate_mode_eq_8", args: args{reg: memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}}, input: []int{42}, want: 0},
		{name: "intermediate_mode_eq_8", args: args{reg: memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}}, input: []int{8}, want: 1},
		{name: "intermediate_mode_lt_8", args: args{reg: memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}}, input: []int{3}, want: 1},
		{name: "intermediate_mode_lt_8", args: args{reg: memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}}, input: []int{8}, want: 0},
		{name: "intermediate_mode_lt_8", args: args{reg: memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}}, input: []int{69}, want: 0},

		{name: "pos_mode_if_in_0", args: args{reg: memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}}, input: []int{0}, want: 0},
		{name: "pos_mode_if_in_0", args: args{reg: memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}}, input: []int{10}, want: 1},
		{name: "intermediate_mode_if_in_0", args: args{reg: memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}}, input: []int{0}, want: 0},
		{name: "intermediate_mode_if_in_0", args: args{reg: memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}}, input: []int{10}, want: 1},

		{name: "compare_to_8", args: args{reg: memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			input: []int{2}, want: 999},
		{name: "compare_to_8", args: args{reg: memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			input: []int{8}, want: 1000},
		{name: "compare_to_8", args: args{reg: memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			input: []int{200}, want: 1001},

		{name: "max_thruster_sig_43210", args: args{reg: memory{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}},
			input: []int{4, 0}, want: 4},
		{name: "max_thruster_sig_43210", args: args{reg: memory{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}},
			input: []int{3, 4}, want: 43},
		{name: "max_thruster_sig_43210", args: args{reg: memory{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}},
			input: []int{2, 43}, want: 432},
		{name: "max_thruster_sig_43210", args: args{reg: memory{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}},
			input: []int{1, 432}, want: 4321},
		{name: "max_thruster_sig_43210", args: args{reg: memory{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}},
			input: []int{0, 4321}, want: 43210},

		{name: "max_thruster_sig_54321", args: args{reg: memory{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}},
			input: []int{0, 0}, want: 5},
		{name: "max_thruster_sig_54321", args: args{reg: memory{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}},
			input: []int{1, 5}, want: 54},
		{name: "max_thruster_sig_54321", args: args{reg: memory{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}},
			input: []int{2, 54}, want: 543},
		{name: "max_thruster_sig_54321", args: args{reg: memory{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}},
			input: []int{3, 543}, want: 5432},
		{name: "max_thruster_sig_54321", args: args{reg: memory{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}},
			input: []int{4, 5432}, want: 54321},

		{name: "max_thruster_sig_65210", args: args{reg: memory{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}},
			input: []int{1, 0}, want: 6},
		{name: "max_thruster_sig_65210", args: args{reg: memory{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}},
			input: []int{0, 6}, want: 65},
		{name: "max_thruster_sig_65210", args: args{reg: memory{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}},
			input: []int{4, 65}, want: 652},
		{name: "max_thruster_sig_65210", args: args{reg: memory{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}},
			input: []int{3, 652}, want: 6521},
		{name: "max_thruster_sig_65210", args: args{reg: memory{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}},
			input: []int{2, 6521}, want: 65210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan int, 10)
			for _, i := range tt.input {
				in <- i
			}
			out := make(chan int, 10)
			run(tt.args.reg, in, out)
			if got := <-out; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOp(t *testing.T) {
	type args struct {
		instr int
	}
	tests := []struct {
		name  string
		args  args
		want  instruction
		want1 int
		want2 int
		want3 int
	}{
		{name: "", args: args{instr: 1002}, want: 2, want1: 0, want2: 1, want3: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := parseOp(tt.args.instr)
			if got != tt.want {
				t.Errorf("parseOp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseOp() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseOp() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("parseOp() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

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
