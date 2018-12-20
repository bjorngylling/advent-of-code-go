package main

import (
	"testing"
)

var data = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

func TestParseInput(t *testing.T) {
	rCave, rEntities := parseInput(data)
	if rCave.Width != 7 {
		t.Errorf("Expected cave width to be 7 but was %+v", rCave.Width)
	}
	if rCave.Height != 7 {
		t.Errorf("Expected cave height to be 7 but was %+v", rCave.Height)
	}
	if len(rEntities) != 6 {
		t.Errorf("Expected entity count to be 6 but was %+v", len(rEntities))
	}
	if len(rEntities.Filter(func(e *Entity) bool { return e.Type == GOBLIN })) != 4 {
		t.Errorf("Expected entity count to be 6 but was %+v", len(rEntities))
	}
	if len(rEntities.Filter(func(e *Entity) bool { return e.Type == ELF })) != 2 {
		t.Errorf("Expected entity count to be 6 but was %+v", len(rEntities))
	}
}
