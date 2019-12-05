package main

import (
	"bytes"
	"reflect"
	"strings"
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
		input string
		want  string
	}{
		{name: "simple_io", args: args{reg: memory{3, 0, 4, 0, 99}}, input: "5", want: "5"},

		{name: "pos_mode_eq_8", args: args{reg: memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}}, input: "5", want: "0"},
		{name: "pos_mode_eq_8", args: args{reg: memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}}, input: "8", want: "1"},
		{name: "pos_mode_lt_8", args: args{reg: memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}}, input: "5", want: "1"},
		{name: "pos_mode_lt_8", args: args{reg: memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}}, input: "8", want: "0"},
		{name: "pos_mode_lt_8", args: args{reg: memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}}, input: "42", want: "0"},

		{name: "intermediate_mode_eq_8", args: args{reg: memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}}, input: "42", want: "0"},
		{name: "intermediate_mode_eq_8", args: args{reg: memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}}, input: "8", want: "1"},
		{name: "intermediate_mode_lt_8", args: args{reg: memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}}, input: "3", want: "1"},
		{name: "intermediate_mode_lt_8", args: args{reg: memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}}, input: "8", want: "0"},
		{name: "intermediate_mode_lt_8", args: args{reg: memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}}, input: "69", want: "0"},

		{name: "pos_mode_if_in_0", args: args{reg: memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}}, input: "0", want: "0"},
		{name: "pos_mode_if_in_0", args: args{reg: memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}}, input: "10", want: "1"},
		{name: "intermediate_mode_if_in_0", args: args{reg: memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}}, input: "0", want: "0"},
		{name: "intermediate_mode_if_in_0", args: args{reg: memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}}, input: "10", want: "1"},

		{name: "compare_to_8", args: args{reg: memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			input: "2", want: "999"},
		{name: "compare_to_8", args: args{reg: memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			input: "8", want: "1000"},
		{name: "compare_to_8", args: args{reg: memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			input: "200", want: "1001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var writer bytes.Buffer
			run(tt.args.reg, strings.NewReader(tt.input), &writer)
			if got := writer.String(); !reflect.DeepEqual(got, tt.want) {
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
