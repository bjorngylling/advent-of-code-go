package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code-2018/15_beverage_bandits/cave"
	. "github.com/bjorngylling/advent-of-code-2018/15_beverage_bandits/entity"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"strings"
	"time"
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
			case 'G', 'E':
				entities = append(entities, New(chr, Pos(x, y)))
			}
		}
	}

	return c, entities
}

func neighbours(p Position) [4]Position {
	return [4]Position{p.Sub(Pos(0, 1)), p.Sub(Pos(1, 0)),
		p.Add(Pos(1, 0)), p.Add(Pos(0, 1))}
}

/*
function Dijkstra(Graph, source):
   dist[source]  := 0                     // Distance from source to source is set to 0
   for each vertex v in Graph:            // Initializations
	   if v ≠ source
		   dist[v]  := infinity           // Unknown distance function from source to each node set to infinity
	   add v to Q                         // All nodes initially in Q

  while Q is not empty:                  // The main loop
	  v := vertex in Q with min dist[v]  // In the first run-through, this vertex is the source node
	  remove v from Q

	  for each neighbor u of v:           // where neighbor u has not yet been removed from Q.
		  alt := dist[v] + length(v, u)
		  if alt < dist[u]:               // A shorter path to u has been found
			  dist[u]  := alt            // Update distance of u

  return dist[]
*/
func Dijkstra(cave *cave.Cave, entities Entities, source Position) (map[Position]int, map[Position]Position) {
	q := []Position{source}             // List containing all unvisited positions
	dist := map[Position]int{source: 0} // Position -> cost to move there from source
	prev := make(map[Position]Position) // Position -> previous position

	// Add all unblocked positions to q and set the distance there to "infinity"
	for y := 0; y < cave.Height; y++ {
		for x := 0; x < cave.Width; x++ {
			if !cave.Blocked(x, y) && !entities.Any(At(x, y)) {
				p := Pos(x, y)
				q = append(q, p)
				dist[p] = math.MaxInt32
			}
		}
	}
	for len(q) > 0 {
		// Find the unvisited position with the lowest distance from source
		iu := -1
		for i, pos := range q {
			if iu == -1 || dist[pos] < dist[q[iu]] {
				iu = i
			}
		}
		u := q[iu]
		q = append(q[:iu], q[iu+1:]...)

		// Check all neighbours of u if there is a shorter path there
		for _, v := range neighbours(u) {
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	// Remove unreachable positions
	for k, v := range dist {
		if v == math.MaxInt32 {
			delete(dist, k)
			delete(prev, k)
		}
	}
	return dist, prev
}

func step(c *cave.Cave, entities Entities) bool {
	entities.Sort()
	for _, ent := range entities {
		alive := entities.Filter(Alive)
		var nearbyEnemies Entities
		for _, n := range neighbours(ent.Pos) {
			if !c.Blocked(n.X, n.Y) {
				nearbyEnemies = append(nearbyEnemies, alive.Filter(At(n.X, n.Y), NotOfType(ent.Type), Alive)...)
			}
		}

		// MOVEMENT
		if len(nearbyEnemies) == 0 {
			// find open spots next to enemies
			var potential []Position
			for _, target := range alive.Filter(NotOfType(ent.Type), Alive) {
				for _, n := range neighbours(target.Pos) {
					if !c.Blocked(n.X, n.Y) && len(alive.Filter(At(n.X, n.Y), Alive)) == 0 {
						potential = append(potential, n)
					}
				}
			}
			// find closest open spot using Dijkstra
			dist, prev := Dijkstra(c, alive.Filter(Alive), ent.Pos)
			closest := Pos(-1, -1)
			for _, p := range potential {
				if d, ok := dist[p]; ok {
					if closest == Pos(-1, -1) ||
						d < dist[closest] ||
						(d == dist[closest] && LessThan(p, closest)) {
						closest = p
					}
				}
			}
			// move towards closest open position next to enemy
			if !closest.Eq(Pos(-1, -1)) {
				ent.Move(nextStep(dist, prev, closest))
			}

			for _, n := range neighbours(ent.Pos) {
				if !c.Blocked(n.X, n.Y) {
					nearbyEnemies = append(nearbyEnemies, alive.Filter(At(n.X, n.Y), NotOfType(ent.Type), Alive)...)
				}
			}
		}

		// ATTACKING
		if len(nearbyEnemies) > 0 {
			// Find enemy with lowest HP
			lowest := 0
			for i := range nearbyEnemies {
				if nearbyEnemies[i].HP < nearbyEnemies[lowest].HP {
					lowest = i
				}
			}
			if err := ent.Attack(nearbyEnemies[lowest]); err != nil {
				if e, ok := err.(NotInReachError); ok {
					log.Fatal(e)
				}
			}
		}

		if len(alive.Filter(NotOfType(GOBLIN), Alive)) == 0 || len(alive.Filter(NotOfType(ELF), Alive)) == 0 {
			return false
		}
	}
	return true
}

func nextStep(dist map[Position]int, prev map[Position]Position, goal Position) Position {
	for cur := goal; true; {
		if dist[prev[cur]] == 0 {
			return cur
		}
		cur = prev[cur]
	}
	return goal
}

func runSimulation(c *cave.Cave, entities Entities) (stepCount int) {
	for step(c, entities) {
		stepCount++
	}
	return
}

func printState(cave *cave.Cave, entities Entities, dist map[Position]int) {
	var s string
	for y := 0; y < cave.Height; y++ {
		var ents Entities
		for x := 0; x < cave.Width; x++ {
			if cave.Blocked(x, y) {
				s += "#"
			} else if l := entities.Filter(At(x, y), Alive); len(l) > 0 {
				ents = append(ents, l...)
				s += string(l[0].Type)
			} else if d, ok := dist[Pos(x, y)]; ok {
				s += fmt.Sprintf("%d", d)
			} else {
				s += "."
			}
		}
		if len(ents) > 0 {
			s += "  "
			for i, e := range ents {
				s += fmt.Sprintf("%s(%d)", string(e.Type), e.HP)
				if i < len(ents)-1 {
					s += ", "
				}
			}
		}
		s += "\n"
	}
	fmt.Print(s)
}

func main() {
	f, err := os.Create("day15.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	start := time.Now()
	fileContent, err := ioutil.ReadFile("15_beverage_bandits/day15_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	c, entities := parseInput(string(fileContent))
	steps := runSimulation(c, entities)
	hpPool := 0
	for _, e := range entities.Filter(Alive) {
		hpPool += e.HP
	}
	fmt.Printf("Day 15 part 1 result: %+v\n", steps*hpPool)

	fmt.Printf("Day 15 part 2 result: %+v\n", nil)
	fmt.Println(time.Since(start))
}
