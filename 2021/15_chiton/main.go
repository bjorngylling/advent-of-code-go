package main

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/bjorngylling/advent-of-code/util"
)

func neighbours(p image.Point) [4]image.Point {
	return [4]image.Point{p.Sub(image.Pt(0, 1)), p.Sub(image.Pt(1, 0)),
		p.Add(image.Pt(1, 0)), p.Add(image.Pt(0, 1))}
}

func dijkstra(w []int, width, height int, source, target image.Point) int {
	var q []image.Point                       // List containing all unvisited positions
	dist := make(map[image.Point]int)         // Position -> cost to move there from source
	prev := make(map[image.Point]image.Point) // Position -> previous position

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := image.Pt(x, y)
			dist[p] = math.MaxInt
			q = append(q, p)
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
			if u.X >= 0 && u.X < width && u.Y >= 0 && u.Y < height {
				alt := dist[u] + w[u.X+u.Y*width]
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}

	path := map[image.Point]struct{}{}
	u := target
	if _, ok := prev[u]; ok || u == source {
		for ok {
			path[u] = struct{}{}
			u, ok = prev[u]
		}
	}

	return dist[target] - w[source.X+source.Y*width] + w[target.X+target.Y*width]
}

func extendMap(w []int, width, height int) []int {
	m := make([]int, width*height*5*5)
	newWidth := width * 5
	for row := 0; row < 5; row++ {
		for y := 0; y < height; y++ {
			for col := 0; col < 5; col++ {
				for x := 0; x < width; x++ {
					v := w[x+y*width] + row + col
					for v > 9 {
						v -= 9
					}
					px, py := width*col+x, height*row+y
					m[px+py*newWidth] = v
				}
			}
		}
	}
	return m
}

func solve(input string) (string, string) {
	spl := strings.Split(input, "\n")
	height, width := len(spl), len(spl[0])
	w := make([]int, width*height)
	for y, ln := range spl {
		for x, c := range ln {
			w[x+y*width] = util.GetInt(string(c))
		}
	}

	p2map := extendMap(w, width, height)

	return strconv.Itoa(dijkstra(w, width, height, image.Pt(0, 0), image.Pt(width-1, height-1))), strconv.Itoa(dijkstra(p2map, width*5, height*5, image.Pt(0, 0), image.Pt(width*5-1, height*5-1)))
}

func main() {
	go http.ListenAndServe("localhost:8080", nil)

	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
