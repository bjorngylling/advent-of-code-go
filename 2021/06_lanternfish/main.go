package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

func sumFish(fish [9]int) int {
	var sum int
	for _, v := range fish {
		sum += v
	}
	return sum
}

func solve(input string) (string, string) {
	var nums [9]int
	for _, n := range strings.Split(input, ",") {
		nums[util.GetInt(n)]++
	}
	part1 := 0
	for i := 0; i < 256; i++ {
		if i == 80 {
			part1 = sumFish(nums)
		}
		var spawning = nums[0]
		nums[7] += nums[0]
		for j := 1; j < len(nums); j++ {
			nums[j-1] = nums[j]
		}
		nums[8] = spawning
	}
	return strconv.Itoa(part1), strconv.Itoa(sumFish(nums))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
