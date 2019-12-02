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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.reg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
