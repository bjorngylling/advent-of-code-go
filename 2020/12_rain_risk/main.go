package main

import (
	"bufio"
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"image"
	"strconv"
	"strings"
	"time"
)

type direction int

const (
	north direction = iota
	south
	east
	west
	left
	right
	forward
)

var dirToVel = map[direction]image.Point{north: image.Pt(0, -1), south: image.Pt(0, 1), east: image.Pt(1, 0), west: image.Pt(-1, 0)}
var leftTurn = map[direction]direction{north: west, west: south, south: east, east: north}
var rightTurn = map[direction]direction{north: east, east: south, south: west, west: north}

type instr struct {
	dir  direction
	dist int
}

func solve(input string) (string, string) {
	var instructions []instr
	scanner := bufio.NewScanner(strings.NewReader(input))
	dirTable := map[rune]direction{'N': north, 'S': south, 'E': east, 'W': west, 'L': left, 'R': right, 'F': forward}
	for scanner.Scan() {
		i := instr{}
		s := scanner.Text()
		i.dir = dirTable[rune(s[0])]
		i.dist, _ = strconv.Atoi(s[1:])
		instructions = append(instructions, i)
	}
	return strconv.Itoa(util.ManhattanDistance(image.Pt(0, 0), part1(instructions))),
		strconv.Itoa(util.ManhattanDistance(image.Pt(0, 0), part2(instructions)))
}

func part2(instructions []instr) image.Point {
	pos := image.Pt(0, 0)
	way := image.Pt(10, -1)
	for _, i := range instructions {
		switch i.dir {
		case forward:
			pos = pos.Add(way.Mul(i.dist))
			break
		case left:
			for j := 0; j < i.dist/90; j++ {
				way = image.Pt(way.Y, -way.X)
			}
			break
		case right:
			for j := 0; j < i.dist/90; j++ {
				way = image.Pt(-way.Y, way.X)
			}
			break
		default:
			way = way.Add(dirToVel[i.dir].Mul(i.dist))
			break
		}
	}
	return pos
}

func part1(instructions []instr) image.Point {
	pos := image.Pt(0, 0)
	var dir = east
	for _, i := range instructions {
		var d direction
		switch i.dir {
		case forward:
			pos = pos.Add(dirToVel[dir].Mul(i.dist))
			break
		case left:
			for j := 0; j < i.dist/90; j++ {
				dir = leftTurn[dir]
			}
			d = dir
			break
		case right:
			for j := 0; j < i.dist/90; j++ {
				dir = rightTurn[dir]
			}
			d = dir
			break
		default:
			d = i.dir
			pos = pos.Add(dirToVel[d].Mul(i.dist))
			break
		}
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
