package main

import (
	"github.com/bjorngylling/advent-of-code-2018/17_reservoir_research/grid"
	"testing"
)

var input = `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504`

func TestParseInput(t *testing.T) {
	scan := parseInput(input)
	cases := []struct {
		in   grid.Position
		want grid.Type
	}{
		{grid.Pos(495, 7), grid.Clay},
		{grid.Pos(500, 7), grid.Clay},
		{grid.Pos(501, 7), grid.Clay},
		{grid.Pos(506, 1), grid.Clay},
		{grid.Pos(506, 2), grid.Clay},
		{grid.Pos(507, 2), grid.Sand},
	}
	for _, c := range cases {
		got := scan.At(c.in)
		if got != c.want {
			t.Errorf("world.At(Pos%s) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSimulate(t *testing.T) {
	world := parseInput(input)
	source := grid.Pos(500, 0)
	simulate(source, world, world.ViewPort(0).Max.Y)

	got := count(world, func(t grid.Type) bool { return t == grid.WetSand || t == grid.Water })
	expected := 57
	if got != expected {
		t.Fatalf("WetSand(world) + Water(world) == %d, want %d\nWorld state:\n%s", got, expected,
			renderWorld(world, source, world.ViewPort(2)))
	}
	got = count(world, func(t grid.Type) bool { return t == grid.Water })
	expected = 29
	if got != expected {
		t.Fatalf("Water(world) == %d, want %d\nWorld state:\n%s", got, expected,
			renderWorld(world, source, world.ViewPort(2)))
	}
}

func TestSimulateSplitFlow(t *testing.T) {
	input := `x=496, y=4..6
x=504, y=4..6
y=6, x=496..504
x=500, y=2..2`
	world := parseInput(input)
	source := grid.Pos(500, 0)
	simulate(source, world, world.ViewPort(0).Max.Y)

	got := count(world, func(t grid.Type) bool { return t == grid.WetSand || t == grid.Water })
	expected := 33
	if got != expected {
		t.Fatalf("countWetPositions(world) == %d, want %d\nWorld state:\n%s", got, expected,
			renderWorld(world, source, world.ViewPort(2)))
	}
}

func TestSimulatePrematureEnd(t *testing.T) {
	input := `x=495, y=4..6
x=500, y=4..6
y=6, x=495..500`
	world := parseInput(input)
	source := grid.Pos(500, 0)
	simulate(source, world, world.ViewPort(0).Max.Y)

	got := count(world, func(t grid.Type) bool { return t == grid.WetSand || t == grid.Water })
	expected := 14
	if got != expected {
		t.Fatalf("countWetPositions(world) == %d, want %d\nWorld state:\n%s", got, expected,
			renderWorld(world, source, world.ViewPort(2)))
	}
}
