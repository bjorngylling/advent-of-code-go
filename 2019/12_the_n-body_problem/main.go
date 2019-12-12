package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

type vec3 struct {
	x, y, z int
}

func (lhs vec3) add(rhs vec3) vec3 {
	return vec3{x: lhs.x + rhs.x, y: lhs.y + rhs.y, z: lhs.z + rhs.z}
}

type moonBody struct {
	pos vec3
	vel vec3
}

func (m moonBody) energy() int {
	return (util.Abs(m.pos.x) + util.Abs(m.pos.y) + util.Abs(m.pos.z)) *
		(util.Abs(m.vel.x) + util.Abs(m.vel.y) + util.Abs(m.vel.z))
}

func (m moonBody) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>",
		m.pos.x, m.pos.y, m.pos.z, m.vel.x, m.vel.y, m.vel.z)
}

func pairs(lst []*moonBody) [][2]*moonBody {
	var pairs [][2]*moonBody
	for i, a := range lst {
		for _, b := range lst[i+1:] {
			pairs = append(pairs, [2]*moonBody{a, b})
		}
	}
	return pairs
}

func solve(input string, steps int) (string, string) {
	var moons []*moonBody
	for _, ln := range strings.Split(input, "\n") {
		m := &moonBody{}
		fmt.Sscanf(ln, "<x=%d, y=%d, z=%d>", &m.pos.x, &m.pos.y, &m.pos.z)
		moons = append(moons, m)
	}

	for i := 0; i < steps; i++ {
		for _, p := range pairs(moons) {
			a, b := p[0], p[1]
			switch {
			case a.pos.x < b.pos.x:
				a.vel.x++
				b.vel.x--
			case a.pos.x > b.pos.x:
				a.vel.x--
				b.vel.x++
			}
			switch {
			case a.pos.y < b.pos.y:
				a.vel.y++
				b.vel.y--
			case a.pos.y > b.pos.y:
				a.vel.y--
				b.vel.y++
			}
			switch {
			case a.pos.z < b.pos.z:
				a.vel.z++
				b.vel.z--
			case a.pos.z > b.pos.z:
				a.vel.z--
				b.vel.z++
			}
		}
		for _, m := range moons {
			m.pos = m.pos.add(m.vel)
		}
	}

	totalEnergy := 0
	for _, m := range moons {
		totalEnergy += m.energy()
	}

	return strconv.Itoa(totalEnergy), ""
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle, 1000)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
