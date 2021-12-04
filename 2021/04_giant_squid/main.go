package main

import (
	"fmt"
	"github.com/bjorngylling/advent-of-code/util"
	"strconv"
	"strings"
	"time"
)

func checkCard(card [5][5]int, played map[int]struct{}) []int {
	for i := 0; i < 5; i++ {
		var marked []int
		for j := 0; j < 5; j++ {
			if _, ok := played[card[i][j]]; ok {
				marked = append(marked, card[i][j])
			}
			if len(marked) == 5 {
				return marked
			}
		}
		marked = nil
	}
	for j := 0; j < 5; j++ {
		var marked []int
		for i := 0; i < 5; i++ {
			if _, ok := played[card[i][j]]; ok {
				marked = append(marked, card[i][j])
			}
			if len(marked) == 5 {
				return marked
			}
		}
		marked = nil
	}
	return nil
}

func sumUnmarked(card [5][5]int, played map[int]struct{}) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if _, ok := played[card[i][j]]; !ok {
				sum += card[i][j]
			}
		}
	}
	return sum
}

func part1(nums []int, cards [][5][5]int) int {
	played := map[int]struct{}{}
	for _, num := range nums {
		played[num] = struct{}{}
		for _, c := range cards {
			if checkCard(c, played) != nil {
				return sumUnmarked(c, played) * num
			}
		}
	}
	return 0
}

func part2(nums []int, cards [][5][5]int) int {
	played := map[int]struct{}{}
	ignore := map[int]struct{}{}
	for _, num := range nums {
		played[num] = struct{}{}
		for i, c := range cards {
			if _, ok := ignore[i]; ok {
				continue
			}
			if checkCard(c, played) != nil {
				if len(ignore) == len(cards)-1 {
					return sumUnmarked(c, played) * num
				}
				ignore[i] = struct{}{}
			}
		}
	}
	return 0
}

func solve(input string) (string, string) {
	lns := strings.Split(input, "\n")
	var nums []int
	for _, n := range strings.Split(lns[0], ",") {
		nums = append(nums, util.GetInt(n))
	}

	var cards [][5][5]int
	var card [5][5]int
	i := 0
	for _, ln := range lns[2:] {
		if ln == "" {
			cards = append(cards, card)
			card = [5][5]int{}
			i = 0
		} else {
			fmt.Sscan(ln, &card[i][0], &card[i][1], &card[i][2], &card[i][3], &card[i][4])
			i++
		}
	}
	cards = append(cards, card)

	return strconv.Itoa(part1(nums, cards)), strconv.Itoa(part2(nums, cards))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
