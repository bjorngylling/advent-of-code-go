package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/2019/intcode"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"time"
)

const (
	left  = 0
	right = 1
)

var directions = []image.Point{image.Pt(0, 1), image.Pt(1, 0), image.Pt(0, -1), image.Pt(-1, 0)}

func turn(dir image.Point, turn int) image.Point {
	dirId := 0
	for i, d := range directions {
		if dir.Eq(d) {
			dirId = i
		}
	}
	switch turn {
	case left:
		if dirId == 0 {
			dirId = 3
		} else {
			dirId--
		}
	case right:
		if dirId == 3 {
			dirId = 0
		} else {
			dirId++
		}
	}
	return directions[dirId]
}

func paint(input string, firstSquare int) map[image.Point]int {
	in := make(chan int, 1)
	out := make(chan int)
	c, err := intcode.Init(input, in, out)
	if err != nil {
		panic(err)
	}

	hull := map[image.Point]int{image.Pt(0, 0): firstSquare}

	// The robot starts at 0,0 facing north/up.
	pos := image.Pt(0, 0)
	dir := image.Pt(0, 1)

	// First thing the robot will do is to read the starting sq color so prefill the input.
	in <- getColor(hull, pos)

	go func() {
		c.Run()
		close(out)
	}()

	for c := range out {
		// Paint current position
		hull[pos] = c

		// Move
		t := <-out
		dir = turn(dir, t)
		pos = pos.Add(dir)

		// Input the color of the new position
		in <- getColor(hull, pos)
	}

	return hull
}

func solve(input string) (string, string) {
	part2Hull := paint(input, 1)
	img := image.NewGray(image.Rect(-0, -1, 41, 7))
	for pos, c := range part2Hull {
		if c == 0 {
			img.Set(pos.X, -pos.Y, color.Black)
		} else {
			img.Set(pos.X, -pos.Y, color.White)
		}
	}
	f, _ := os.Create("part2.png")
	defer f.Close()
	png.Encode(f, img)

	return strconv.Itoa(len(paint(input, 0))), "See part2.png"
}

func getColor(hull map[image.Point]int, pos image.Point) int {
	color := 0
	if c, found := hull[pos]; found {
		color = c
	}
	return color
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
