package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2018/15_beverage_bandits/entity"
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
	rCave, rEntities := parseInput(data, 3)
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
	cave, entities := parseInput(data, 3)
	dist, _ := Dijkstra(cave, entities, entities[0].Pos)
	result := ""
	for y := 0; y < cave.Height; y++ {
		for x := 0; x < cave.Width; x++ {
			if cave.Blocked(x, y) {
				result = fmt.Sprintf("%s#", result)
			} else if l := entities.Filter(entity.At(x, y)); len(l) > 0 {
				result = fmt.Sprintf("%s%s", result, string(l[0].Type))
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

func TestNextStep(t *testing.T) {
	cave, entities := parseInput(data, 3)
	dist, prev := Dijkstra(cave, entities, entities[0].Pos)
	g := entity.Pos(5, 1)
	r := nextStep(dist, prev, g)
	e := entity.Pos(3, 1)
	if !r.Eq(e) {
		t.Errorf("Expected first step towards %+v to be %+v but was %+v", g, e, r)
	}

	g = entity.Pos(1, 3)
	r = nextStep(dist, prev, g)
	e = entity.Pos(1, 1)
	if !r.Eq(e) {
		t.Errorf("Expected first step towards %+v to be %+v but was %+v", g, e, r)
	}

	d := `#####
#.E.#
#.#.#
#.#.#
#...#
#####`
	cave, entities = parseInput(d, 3)
	dist, prev = Dijkstra(cave, entities, entity.Pos(2, 1))
	g = entity.Pos(2, 4)
	r = nextStep(dist, prev, g)
	e = entity.Pos(1, 1)
	if !r.Eq(e) {
		t.Errorf("Expected first step towards %+v to be %+v but was %+v", g, e, r)
		printState(cave, entities, dist)
	}
}

func TestRunSimulation1(t *testing.T) {
	c, entities := parseInput(data, 3)
	steps := runSimulation(c, entities)
	entities = entities.Filter(entity.Alive)
	hpPool := 0
	for _, e := range entities {
		hpPool += e.HP
	}
	if steps != 46 {
		t.Errorf("Expected battle to end after 47 rounds but ended after %d", steps)
	}
	if hpPool != 590 {
		t.Errorf("Expected winning side HP pool to be 590 but was %d", hpPool)
	}
}

func TestRunSimulation2(t *testing.T) {
	d := `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`
	c, entities := parseInput(d, 3)
	steps := runSimulation(c, entities)
	entities = entities.Filter(entity.Alive)
	hpPool := 0
	for _, e := range entities {
		hpPool += e.HP
	}
	if steps != 37 {
		t.Errorf("Expected battle to end after 37 rounds but ended after %d", steps)
	}
	if hpPool != 982 {
		t.Errorf("Expected winning side HP pool to be 982 but was %d", hpPool)
	}
}

func TestRunSimulation3(t *testing.T) {
	d := `#######   
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#   
#######`
	c, entities := parseInput(d, 3)
	steps := runSimulation(c, entities)
	entities = entities.Filter(entity.Alive)
	hpPool := 0
	for _, e := range entities {
		hpPool += e.HP
	}
	if steps != 46 {
		t.Errorf("Expected battle to end after 46 rounds but ended after %d", steps)
	}
	if hpPool != 859 {
		t.Errorf("Expected winning side HP pool to be 859 but was %d", hpPool)
	}
}

func TestRunSimulation4(t *testing.T) {
	d := `#######   
#E.G#.#
#.#G..#
#G.#.G#   
#G..#.#
#...E.#
#######`
	c, entities := parseInput(d, 3)
	steps := 0
	for step(c, entities.Filter(entity.Alive)) {
		steps++
	}
	entities = entities.Filter(entity.Alive)
	hpPool := 0
	for _, e := range entities {
		hpPool += e.HP
	}
	if steps != 35 {
		t.Errorf("Expected battle to end after 35 rounds but ended after %d", steps)
	}
	if hpPool != 793 {
		t.Errorf("Expected winning side HP pool to be 793 but was %d", hpPool)
	}
}

func TestRunSimulation5(t *testing.T) {
	d := `#######   
#.E...#   
#.#..G#
#.###.#   
#E#G#G#   
#...#G#
#######`
	c, entities := parseInput(d, 3)
	steps := 0
	for step(c, entities.Filter(entity.Alive)) {
		steps++
	}
	entities = entities.Filter(entity.Alive)
	hpPool := 0
	for _, e := range entities {
		hpPool += e.HP
	}
	if steps != 54 {
		t.Errorf("Expected battle to end after 54 rounds but ended after %d", steps)
	}
	if hpPool != 536 {
		t.Errorf("Expected winning side HP pool to be 536 but was %d", hpPool)
	}
}

func TestRunSimulation6(t *testing.T) {
	d := `#########   
#G......#
#.E.#...#
#..##..G#
#...##..#   
#...#...#
#.G...G.#   
#.....G.#   
#########`
	c, entities := parseInput(d, 3)
	steps := 0
	for step(c, entities.Filter(entity.Alive)) {
		steps++
	}
	entities = entities.Filter(entity.Alive)
	hpPool := 0
	for _, e := range entities {
		hpPool += e.HP
	}
	if steps != 20 {
		t.Errorf("Expected battle to end after 20 rounds but ended after %d", steps)
	}
	if hpPool != 937 {
		t.Errorf("Expected winning side HP pool to be 937 but was %d", hpPool)
	}
}

func TestRunSimulationCheatingElves1(t *testing.T) {
	steps, hpPool := runSimulationCheatingElves(data)
	if steps != 29 {
		t.Errorf("Expected battle to end after 29 rounds but ended after %d", steps)
	}
	if hpPool != 172 {
		t.Errorf("Expected winning side HP pool to be 172 but was %d", hpPool)
	}
}
