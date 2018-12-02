package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func generateChecksum(ids []string) int {
	var twos, threes int
	for _, id := range ids {
		occurrences := make(map[rune]int)
		for _, r := range id {
			occurrences[r] += 1
		}
		var foundTwos, foundThrees bool
		for _, v := range occurrences {
			if !foundTwos && v == 2 {
				twos++
				foundTwos = true
			} else if !foundThrees && v == 3 {
				threes++
				foundThrees = true
			}
		}
	}

	return twos * threes
}

func main() {
	fileContent, err := ioutil.ReadFile("02_inventory_management_system/day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 2 part 1 result: %d\n", generateChecksum(strings.Split(string(fileContent), "\n")))
}
