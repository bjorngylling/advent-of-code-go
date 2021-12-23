package main

import (
	"fmt"
	"image"
	"strconv"
	"time"
)

func solve(input string) (string, string) {
	var x1, x2, y1, y2 int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	// part 1
	y := 0
	for velY := (y1 * -1) - 1; velY != 0; velY-- {
		y += velY
	}

	// part2
	validCount := 0
	pos := image.Pt(0, 0)
	tar := image.Rect(x1, y1, x2+1, y2+1)
	fmt.Println(tar.Max.X + 1)
	for vel := image.Pt(0, -1000); vel.X < tar.Max.X+1; vel.X++ {
		for vel.Y = -1000; vel.Y < 1000; vel.Y++ {
			if shoot(pos, vel, tar) {
				validCount++
			}
		}
	}

	return strconv.Itoa(y), strconv.Itoa(validCount)
}

func shoot(pos, vel image.Point, tar image.Rectangle) bool {
	for {
		pos = pos.Add(vel)

		if vel.X > 0 {
			vel.X--
		} else if vel.X < 0 {
			vel.X++
		}

		vel.Y--

		if pos.In(tar) {
			return true
		}

		if (pos.X > tar.Max.X) || (pos.X < tar.Min.X && vel.X == 0) {
			return false
		}
		if pos.Y < tar.Min.Y {
			return false
		}
	}
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
