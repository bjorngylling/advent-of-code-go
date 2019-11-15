package main

import "fmt"

type Grid struct {
	w, h     int
	t        []int
	serialNo int
}

func newGrid(w, h, serialNo int) *Grid {
	return &Grid{w: w, h: h, t: make([]int, w*h), serialNo: serialNo}
}

func (g *Grid) at(x, y int) int {
	if x < 0 || y < 0 {
		return powerLevel(x, y, g.serialNo)
	}
	return g.t[x+y*g.w]
}

func (g *Grid) set(x, y, val int) {
	g.t[x+y*g.w] = val
}

func powerLevel(x, y, gridSerialNo int) int {
	rackId := x + 10
	return (rackId*y+gridSerialNo)*rackId/100%10 - 5
}

func groupPowerLevel(x, y, size int, g *Grid) int {
	x, y = x-1, y-1
	return g.at(x+size, y+size) + g.at(x, y) - g.at(x+size, y) - g.at(x, y+size)
}

func findHighestGroupPower(s, gridSerialNo int) (hX, hY int) {
	g := summedAreaTable(300, 300, gridSerialNo)
	var hLevel int
	for x := 0; x < 300-s; x++ {
		for y := 0; y < 300-s; y++ {
			cLevel := groupPowerLevel(x, y, s, g)
			if hLevel < cLevel {
				hLevel, hX, hY = cLevel, x, y
			}
		}
	}
	return
}

func findMaxGroupPower(gridSerialNo int) (hX, hY, hSize int) {
	g := summedAreaTable(300, 300, gridSerialNo)
	var hLevel int
	for s := 1; s < 300; s++ {
		for x := 0; x < 300-s; x++ {
			for y := 0; y < 300-s; y++ {
				cLevel := groupPowerLevel(x, y, s, g)
				if hLevel < cLevel {
					hLevel, hX, hY, hSize = cLevel, x, y, s
				}
			}
		}
	}
	return
}

func summedAreaTable(w, h, gridSerialNo int) *Grid {
	g := newGrid(w, h, gridSerialNo)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			sum := powerLevel(x, y, gridSerialNo) + g.at(x, y-1) + g.at(x-1, y) - g.at(x-1, y-1)

			g.set(x, y, sum)
		}
	}
	return g
}

func main() {
	gridSerialNo := 5468

	hX, hY := findHighestGroupPower(3, gridSerialNo)
	fmt.Printf("Day 11 part 1 result: %d,%d\n", hX, hY)

	hX, hY, hSize := findMaxGroupPower(gridSerialNo)
	fmt.Printf("Day 11 part 2 result: %d,%d,%d\n", hX, hY, hSize)
}
