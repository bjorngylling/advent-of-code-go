package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	var adapters []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		adapters = append(adapters, util.GetInt(scanner.Text()))
	}
	if adapters == nil {
		panic(errors.New("unable to parse input"))
	}
	sort.Ints(adapters)

	var ones, twos, threes int
	prev := 0
	for _, i := range adapters {
		switch i - prev {
		case 1:
			ones++
			break
		case 2:
			twos++
			break
		case 3:
			threes++
			break
		}
		prev = i
	}
	threes++

	adapters = append([]int{0}, append(adapters, adapters[len(adapters)-1]+3)...)
	stepCount := map[int]int{0: 1}
	for i := range adapters {
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j] > adapters[i]+3 {
				break
			}
			stepCount[j] = stepCount[j] + stepCount[i]
		}
	}

	return strconv.Itoa(ones * threes), strconv.Itoa(stepCount[len(adapters)-1])
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
