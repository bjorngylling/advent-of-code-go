package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type kv = struct {
	key   rune
	value int
}

func solve(input string) (string, string) {
	inputRows := strings.Split(input, "\n")
	var count []map[rune]*kv
	var sortByCount [][]*kv
	for i := 0; i < len(inputRows[0]); i++ {
		count = append(count, map[rune]*kv{})
		sortByCount = append(sortByCount, []*kv{})
	}
	if count == nil {
		panic(fmt.Errorf("count == nil"))
	}
	if sortByCount == nil {
		panic(fmt.Errorf("sortByCount == nil"))
	}
	for _, ln := range inputRows {
		for col, c := range ln {
			if _, found := count[col][c]; found {
				count[col][c].value++
			} else {
				count[col][c] = &kv{c, 1}
				sortByCount[col] = append(sortByCount[col], count[col][c])
			}
			sort.Slice(sortByCount[col], func(i, j int) bool {
				return sortByCount[col][i].value > sortByCount[col][j].value
			})
		}
	}
	return row(sortByCount, 0), row(sortByCount, len(sortByCount[0])-1)
}

func row(columns [][]*kv, row int) string {
	output := ""
	for _, col := range columns {
		output += string(col[row].key)
	}
	return output
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
