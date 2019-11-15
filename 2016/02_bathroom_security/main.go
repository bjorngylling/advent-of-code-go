package main

import (
	"fmt"
	"image"
	"strings"
	"time"
)

var delta = map[rune]image.Point{
	'U': image.Pt(0, 1),
	'R': image.Pt(1, 0),
	'D': image.Pt(0, -1),
	'L': image.Pt(-1, 0),
}

type keypad = map[image.Point]string

var keypadPart1 = keypad{
	image.Pt(0, 2): "1",
	image.Pt(1, 2): "2",
	image.Pt(2, 2): "3",
	image.Pt(0, 1): "4",
	image.Pt(1, 1): "5",
	image.Pt(2, 1): "6",
	image.Pt(0, 0): "7",
	image.Pt(1, 0): "8",
	image.Pt(2, 0): "9",
}
var keypadPart2 = keypad{
	image.Pt(2, 4): "1",
	image.Pt(1, 3): "2",
	image.Pt(2, 3): "3",
	image.Pt(3, 3): "4",
	image.Pt(0, 2): "5",
	image.Pt(1, 2): "6",
	image.Pt(2, 2): "7",
	image.Pt(3, 2): "8",
	image.Pt(4, 2): "9",
	image.Pt(1, 1): "A",
	image.Pt(2, 1): "B",
	image.Pt(3, 1): "C",
	image.Pt(2, 0): "D",
}

func solve(input string) (string, string) {
	pos1 := image.Pt(1, 1)
	pos2 := image.Pt(0, 2)
	part1 := ""
	part2 := ""
	for _, command := range strings.Split(input, "\n") {
		for _, dir := range command {
			pos1 = move(pos1, dir, keypadPart1)
			pos2 = move(pos2, dir, keypadPart2)
		}
		part1 += keypadPart1[pos1]
		part2 += keypadPart2[pos2]
	}
	return part1, part2
}

func move(pos image.Point, dir rune, kp keypad) image.Point {
	newPos := pos.Add(delta[dir])
	if _, found := kp[newPos]; found {
		pos = newPos
	}
	return pos
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
