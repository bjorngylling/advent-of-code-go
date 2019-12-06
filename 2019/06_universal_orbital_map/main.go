package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type stringSet map[string]struct{}

func solve(input string) (string, string) {
	orbits := make(map[string]string)
	planets := make(map[string]stringSet)

	for _, ln := range strings.Split(input, "\n") {
		a, b := strings.Split(ln, ")")[0], strings.Split(ln, ")")[1]
		orbits[b] = a
		if _, found := planets[a]; !found {
			planets[a] = stringSet{}
		}
		if _, found := planets[b]; !found {
			planets[b] = stringSet{}
		}
		planets[a][b] = struct{}{}
		planets[b][a] = struct{}{}
	}

	part1 := 0
	for planet := range planets {
		part1 += orbitCount(orbits, planet)
	}

	_, prev := dijkstra(planets, "YOU", "SAN")
	u := "SAN"

	// Since we are using dijkstra and considering the source and target as nodes in the graph we will get distance
	// including the edges between the source/target and the planets they orbit. So we start with a distance of -2.
	part2 := -2
	for {
		if _, found := prev[u]; !found {
			break
		}
		u = prev[u]
		part2++
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func orbitCount(orbits map[string]string, planet string) int {
	d := 0
	for {
		if _, found := orbits[planet]; !found {
			break
		}
		planet = orbits[planet]
		d++
	}
	return d
}

func dijkstra(graph map[string]stringSet, src string, tar string) (map[string]int, map[string]string) {
	q := stringSet{}
	dist := map[string]int{}
	prev := map[string]string{}

	// Add all nodes in the graph to q and set distance there to "infinity"
	for v := range graph {
		dist[v] = math.MaxInt32
		q[v] = struct{}{}
	}
	// Distance to the source node is 0 since we start there
	dist[src] = 0

	for len(q) > 0 {
		// Find the unvisited position with the lowest distance from source
		u := ""
		for p := range q {
			if dist[p] < dist[u] || u == "" {
				u = p
			}
		}

		// Since we are only looking for a path to tar we can abort when we reach it
		if u == tar {
			break
		}

		// Consider u visited
		delete(q, u)

		// Calculate distance to all neighbours
		for v := range graph[u] {
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist, prev
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
