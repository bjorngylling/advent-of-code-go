package main

import (
	"fmt"
	"image"
	"math"
	"strings"
	"time"
)

func neighbours(p image.Point) [4]image.Point {
	return [4]image.Point{p.Sub(image.Pt(0, 1)), p.Sub(image.Pt(1, 0)),
		p.Add(image.Pt(1, 0)), p.Add(image.Pt(0, 1))}
}

func dijkstra(world map[image.Point]bool, source image.Point) (map[image.Point]int, map[image.Point]image.Point) {
	q := []image.Point{source}                // List containing all unvisited positions
	dist := make(map[image.Point]int)         // Position -> cost to move there from source
	prev := make(map[image.Point]image.Point) // Position -> previous position

	// Add all unblocked positions to q and set the distance there to "infinity"
	for p, blocked := range world {
		if !blocked {
			q = append(q, p)
			dist[p] = math.MaxInt32
		}
	}
	dist[source] = 0
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

func isDoor(c rune) bool {
	return 'A' <= c && c <= 'Z'
}

func isKey(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func doorFor(key rune) rune {
	return key - 32
}

func keyFor(door rune) rune {
	return door + 32
}

func parse(input string) (map[image.Point]rune, image.Point) {
	world := map[image.Point]rune{}
	entrance := image.Point{}
	for y, ln := range strings.Split(input, "\n") {
		for x, c := range ln {
			world[image.Pt(x, y)] = c
			if c == '@' {
				entrance = image.Pt(x, y)
			}
		}
	}
	return world, entrance
}

func solve(input string) (string, string) {
	world, entrance := parse(input)
	return "", ""
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
