package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code-2018/15_beverage_bandits/entity"
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
	if len(rEntities.Filter(func(e *entity.Entity) bool { return e.Type == entity.GOBLIN })) != 4 {
		t.Errorf("Expected entity count to be 6 but was %+v", len(rEntities))
	}
	if len(rEntities.Filter(func(e *entity.Entity) bool { return e.Type == entity.ELF })) != 2 {
		t.Errorf("Expected entity count to be 6 but was %+v", len(rEntities))
	}
}

func TestDijkstra(t *testing.T) {
	expected := `#######
#1G123#
#212EG#
#3#3#G#
#45G#E#
#56789#
#######
`
	cave, entities := parseInput(data)
	dist, _ := Dijkstra(cave, entities, entities[0])
	result := ""
	for y := 0; y < cave.Height; y++ {
		for x := 0; x < cave.Width; x++ {
			if cave.Blocked(x, y) {
				result = fmt.Sprintf("%s#", result)
			} else if l := entities.Filter(entity.At(x, y)); len(l) > 0 {
				result = fmt.Sprintf("%s%s", result, l[0].Type)
			} else {
				result = fmt.Sprintf("%s%d", result, dist[entity.Pos(x, y)])
			}
		}
		result += fmt.Sprint("\n")
	}
	if expected != result {
		t.Errorf("Expected result to be:\n%sbut was:\n%s", expected, result)
	}
}
