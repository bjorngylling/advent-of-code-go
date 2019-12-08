package main

import (
	"reflect"
	"testing"
)

func Test_readLayers(t *testing.T) {
	type args struct {
		input string
		w     int
		h     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "basic", args: args{input: "122156789012", w: 3, h: 2}, want: [][]int{{1, 2, 2, 1, 5, 6}, {7, 8, 9, 0, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readLayers(tt.args.input, tt.args.w, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readLayers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imgHash(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "basic", args: args{img: [][]int{{1, 2, 2, 1, 5, 6}, {7, 8, 9, 0, 1, 2}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imgHash(tt.args.img); got != tt.want {
				t.Errorf("imgHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
