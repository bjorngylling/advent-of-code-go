package main

import (
	"fmt"

	"./day6"
	"./day7"
)

func main() {
	fmt.Println("-- Day 6 --")
	cycles, cyclesBetweenDuplicateState := day6.Redistribute(append(make([][]int, 0), day6.CreateBanks(day6.Input)), 0)
	fmt.Printf("cycles:%d, cycles_between_duplicate_state:%d\n", cycles, cyclesBetweenDuplicateState)
	fmt.Println()

	fmt.Println("-- Day 7 --")
	tree := day7.CreateTree(day7.Input)
	fmt.Printf("root_node:%s\n", tree.Root.Name)
	overweight, by := day7.FindOverweightNode(tree.Root)
	fmt.Printf("overweight_node:%s, overweight_by:%d, ideal_weight:%d",
		overweight.Name, by, overweight.Weight - by)
	fmt.Println()
}