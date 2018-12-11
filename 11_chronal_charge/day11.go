package main

import "fmt"

func powerLevel(x, y, gridSerialNo int) int {
	rackId := x + 10
	return (rackId*y+gridSerialNo)*rackId/100%10 - 5
}

func groupPowerLevel(x, y, size, gridSerialNo int) (r int) {
	for i := x; i <= x+size-1; i++ {
		for j := y; j <= y+size-1; j++ {
			r += powerLevel(i, j, gridSerialNo)
		}
	}
	return
}

func findHighestGroupPower(gridSerialNo int) (hX, hY int) {
	var hLevel int
	for x := 0; x < 298; x++ {
		for y := 0; y < 298; y++ {
			cLevel := groupPowerLevel(x, y, 3, gridSerialNo)
			if hLevel < cLevel {
				hLevel, hX, hY = cLevel, x, y
			}
		}
	}
	return
}

func findMaxGroupPower(gridSerialNo int) (hX, hY, hSize int) {
	var hLevel int
	for s := 0; s <= 300; s++ {
		for x := 0; x < 300-s; x++ {
			for y := 0; y < 300-s; y++ {
				cLevel := groupPowerLevel(x, y, s, gridSerialNo)
				if hLevel < cLevel {
					hLevel, hX, hY, hSize = cLevel, x, y, s
				}
			}
		}
	}
	return
}

func main() {
	gridSerialNo := 5468

	hX, hY := findHighestGroupPower(gridSerialNo)
	fmt.Printf("Day 11 part 1 result: %d,%d\n", hX, hY)

	hX, hY, hSize := findMaxGroupPower(gridSerialNo)
	fmt.Printf("Day 11 part 2 result: %d,%d,%d\n", hX, hY, hSize)
}
