package main

import (
	"fmt"

	"./day6"
	)

func main() {
	fmt.Println("-- Day 6 --")
	cycles, cyclesBetweenDuplicateState := day6.Redistribute(append(make([][]int, 0), day6.CreateBanks(day6.Input)), 0)
	fmt.Printf("cycles:%d, cycles_between_duplicate_state:%d\n", cycles, cyclesBetweenDuplicateState)
	fmt.Println()
}