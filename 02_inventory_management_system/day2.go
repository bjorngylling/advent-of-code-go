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

func hammingDistance(s string, t string) (dist int) {
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			dist += 1
		}
	}
	return
}

func findSimilarIds(ids []string) (result []string) {
	for i, s := range ids {
		for _, t := range ids[i+1:] {
			if hammingDistance(s, t) == 1 {
				result = append(result, s, t)
			}
		}
	}
	return
}

func findCommonId(ids []string) (result string) {
	ids = findSimilarIds(ids)
	s, t := ids[0], ids[1]
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			result = s[:i] + s[i+1:]
		}
	}
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("02_inventory_management_system/day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 2 part 1 result: %d\n", generateChecksum(strings.Split(string(fileContent), "\n")))

	fmt.Printf("Day 2 part 2 result: %s\n", findCommonId(strings.Split(string(fileContent), "\n")))
}
