package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func generateChecksum(ids []string) int {
	twos, threes := 0, 0

	for _, id := range ids {
		occurrences := make(map[rune]int)
		for _, r := range id {
			if _, ok := occurrences[r]; !ok {
				occurrences[r] = 1
			} else {
				occurrences[r]++
			}
		}
		foundTwos, foundThrees := false, false
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
