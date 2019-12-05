package main

import (
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
			if got := run(tt.args.reg); !reflect.DeepEqual(got, tt.want) {
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
