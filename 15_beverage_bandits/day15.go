package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code-2018/15_beverage_bandits/cave"
	. "github.com/bjorngylling/advent-of-code-2018/15_beverage_bandits/entity"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func parseInput(in string) (*cave.Cave, Entities) {
	lines := strings.Split(in, "\n")
	c := cave.New(len(lines[0]), len(lines))
	var entities Entities

	for y, ln := range lines {
		for x, chr := range ln {
			switch chr {
			case '#':
				c.SetBlocked(x, y, true)
			case 'G':
				entities = append(entities, New(GOBLIN, Pos(x, y)))
			case 'E':
				entities = append(entities, New(ELF, Pos(x, y)))
			}
		}
	}

	return c, entities
}

func neighbours(p Position) [4]Position {
	return [4]Position{p.Add(Pos(1, 0)), p.Add(Pos(0, 1)),
		p.Sub(Pos(1, 0)), p.Sub(Pos(0, 1))}
}

func Dijkstra(cave *cave.Cave, entities Entities, source *Entity) (map[Position]int, map[Position]Position) {
	q := map[Position]struct{}{source.Pos: {}} // Set containing all unvisited positions
	dist := map[Position]int{source.Pos: 0}    // Position -> cost to move there from source
	prev := make(map[Position]Position)        // Position -> previous position

	// Add all unblocked positions to q and set the distance there to "infinity"
	for y := 0; y < cave.Height; y++ {
		for x := 0; x < cave.Width; x++ {
			if !cave.Blocked(x, y) && !entities.Any(At(x, y)) {
				p := Pos(x, y)
				q[p] = struct{}{}
				dist[p] = math.MaxInt32
			}
		}
	}
	for len(q) > 0 {
		// Find the unvisited position with the lowest distance from source
		u := Pos(-1, -1)
		for pos := range q {
			if u == Pos(-1, -1) || dist[pos] < dist[u] {
				u = pos
			}
		}
		delete(q, u)

		// Check all neighbours of u if there is a shorter path there
		for _, v := range neighbours(u) {
			if _, ok := q[v]; ok {
				alt := dist[u] + 1
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}
	return dist, prev
}

func main() {
	fileContent, err := ioutil.ReadFile("15_beverage_bandits/day15_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	cave, entities := parseInput(string(fileContent))
	fmt.Printf("Day 15 part 1 result: %+v, %+v\n", len(entities), cave.Blocked(0, 0))

	fmt.Printf("Day 15 part 2 result: %+v\n", nil)
}
