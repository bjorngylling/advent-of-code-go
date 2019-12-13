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

func part1(input string, steps int) string {
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
	return strconv.Itoa(totalEnergy)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

type posAndVel1d struct {
	pos, vel int
}

func findPeriod(bodies [4]posAndVel1d) int {
	initialState := bodies
	states := map[[4]posAndVel1d]struct{}{initialState: {}}
	steps := 0
	for ; ; steps++ {
		for i := 0; i < len(bodies); i++ {
			for j := i + 1; j < len(bodies); j++ {
				switch {
				case bodies[i].pos < bodies[j].pos:
					bodies[i].vel++
					bodies[j].vel--
				case bodies[i].pos > bodies[j].pos:
					bodies[i].vel--
					bodies[j].vel++
				}
			}
		}
		for i := range bodies {
			bodies[i].pos += bodies[i].vel
		}
		if _, found := states[bodies]; found && steps > 1 {
			break
		}
		newState := bodies
		states[newState] = struct{}{}
	}
	return steps + 1
}

func part2(input string) string {
	var moons []*moonBody
	for _, ln := range strings.Split(input, "\n") {
		m := &moonBody{}
		fmt.Sscanf(ln, "<x=%d, y=%d, z=%d>", &m.pos.x, &m.pos.y, &m.pos.z)
		moons = append(moons, m)
	}

	var xPlane [4]posAndVel1d
	var yPlane [4]posAndVel1d
	var zPlane [4]posAndVel1d
	for i, m := range moons {
		xPlane[i] = posAndVel1d{pos: m.pos.x, vel: m.vel.x}
		yPlane[i] = posAndVel1d{pos: m.pos.y, vel: m.vel.y}
		zPlane[i] = posAndVel1d{pos: m.pos.z, vel: m.vel.z}
	}

	pX := findPeriod(xPlane)
	pY := findPeriod(yPlane)
	pZ := findPeriod(zPlane)

	return strconv.Itoa(lcm(pX, pY, pZ))
}

func solve(input string, steps int) (string, string) {
	return part1(input, steps), part2(input)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle, 1000)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
