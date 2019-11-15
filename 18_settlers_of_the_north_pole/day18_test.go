package main

import (
	"github.com/bjorngylling/advent-of-code-2018/18_settlers_of_the_north_pole/grid"
	"reflect"
	"testing"
)

var input = `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`

func TestParseInput(t *testing.T) {
	g := parseInput(input)

	got := g.Height
	want := 10
	if got != want {
		t.Fatalf("grid.Height == %d, want %d", got, want)
	}
	got = g.Width
	want = 10
	if got != want {
		t.Fatalf("grid.Height == %d, want %d", got, want)
	}

	cases := []struct {
		in   grid.Position
		want grid.Type
	}{
		{grid.Pos(0, 0), grid.Open},
		{grid.Pos(1, 0), grid.Lumberyard},
		{grid.Pos(6, 0), grid.Open},
		{grid.Pos(7, 0), grid.Trees},
		{grid.Pos(8, 0), grid.Lumberyard},
		{grid.Pos(5, 4), grid.Trees},
	}
	for _, c := range cases {
		got := g.At(c.in)
		if got != c.want {
			t.Errorf("grid.At(Pos%s) == %q, want %q", c.in, grid.TypeToString(got), grid.TypeToString(c.want))
		}
	}
}

func TestNeighbours(t *testing.T) {
	want := []grid.Position{grid.Pos(0, 0), grid.Pos(1, 0), grid.Pos(2, 0),
		grid.Pos(0, 1), grid.Pos(2, 1),
		grid.Pos(0, 2), grid.Pos(1, 2), grid.Pos(2, 2)}
	in := grid.Pos(1, 1)
	got := neighbours(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("neightbours(Pos%s) == %q, want %q", in, got, want)
	}
}

func TestStep(t *testing.T) {
	g := parseInput(input)

	g = step(g)

	cases := []struct {
		in   grid.Position
		want grid.Type
	}{
		{grid.Pos(0, 0), grid.Open},
		{grid.Pos(1, 0), grid.Open},
		{grid.Pos(6, 0), grid.Open},
		{grid.Pos(7, 0), grid.Lumberyard},
		{grid.Pos(8, 0), grid.Lumberyard},
		{grid.Pos(3, 4), grid.Lumberyard},
		{grid.Pos(5, 4), grid.Trees},
		{grid.Pos(2, 9), grid.Open},
		{grid.Pos(5, 9), grid.Trees},
	}
	for _, c := range cases {
		got := g.At(c.in)
		if got != c.want {
			t.Errorf("grid.At(Pos%s) == %q, want %q", c.in, grid.TypeToString(got), grid.TypeToString(c.want))
		}
	}

	for i := 0; i < 9; i++ {
		g = step(g)
	}

	got := count(g, func(t grid.Type) bool { return t == grid.Lumberyard })
	want := 31
	if got != want {
		t.Errorf("count(Lumberyard) == %d, want %d", got, want)
	}
	got = count(g, func(t grid.Type) bool { return t == grid.Trees })
	want = 37
	if got != want {
		t.Errorf("count(Trees) == %d, want %d", got, want)
	}
}
