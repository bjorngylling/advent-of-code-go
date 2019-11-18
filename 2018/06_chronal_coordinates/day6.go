package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func parseInput(in []string) (r []image.Point) {
	for _, ln := range in {
		s := strings.Split(ln, ", ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		r = append(r, image.Pt(x, y))
	}
	return
}

func findBounds(pts []image.Point) image.Rectangle {
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, 0, 0
	for _, pt := range pts {
		if pt.X < minX {
			minX = pt.X
		} else if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		} else if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	return image.Rect(minX, minY, maxX, maxY)
}

func manhattanDistance(a, b image.Point) int {
	return abs(b.X-a.X) + abs(b.Y-a.Y)
}

func fillGrid(pts []image.Point, bounds image.Rectangle) [][]int {
	grid := make([][]int, bounds.Max.Y+bounds.Min.Y)
	for y := range grid {
		grid[y] = make([]int, bounds.Max.X+bounds.Min.X)
		for x := range grid[y] {
			cur := image.Pt(x, y)
			dist := math.MaxInt32
			for i, pt := range pts {
				newDist := manhattanDistance(cur, pt)
				if newDist < dist {
					dist = newDist
					grid[y][x] = i
				} else if newDist == dist {
					grid[y][x] = -1
				}
			}
		}
	}
	return grid
}

func largestContainedArea(grid [][]int, bounds image.Rectangle) (r int) {
	ignore := make(map[int]struct{})
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		ignore[grid[0][x]] = struct{}{}
		ignore[grid[bounds.Max.Y][x]] = struct{}{}
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		ignore[grid[y][0]] = struct{}{}
		ignore[grid[y][bounds.Max.X]] = struct{}{}
	}
	sums := make(map[int]int)
	for y := range grid {
		for _, i := range grid[y] {
			if _, ok := ignore[i]; !ok {
				sums[i] += 1
				if r < sums[i] {
					r = sums[i]
				}
			}
		}
	}
	return
}

func part2(pts []image.Point, bounds image.Rectangle, maxDist int) (r int) {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			cur := image.Pt(x, y)
			sum := 0
			for _, pt := range pts {
				sum += manhattanDistance(cur, pt)
			}
			if sum < maxDist {
				r += 1
			}
		}
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/06_chronal_coordinates/day6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	points := parseInput(strings.Split(string(fileContent), "\n"))

	bounds := findBounds(points)
	grid := fillGrid(points, bounds)

	fmt.Printf("Day 6 part 1 result: %+v\n", largestContainedArea(grid, bounds))

	fmt.Printf("Day 6 part 2 result: %+v\n", part2(points, bounds, 10000))
}
