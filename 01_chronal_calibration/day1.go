package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calculateFrequency(changes []string) int {
	result := 0
	for _, change := range changes {
		change, _ := strconv.Atoi(change)
		result = result + change
	}
	return result
}

func findFirstDuplicateFrequency(changes []string) int {
	freqs := make(map[int]bool)
	curFreq := 0

	for i := 0; !freqs[curFreq]; i++ {
		freqs[curFreq] = true

		change, _ := strconv.Atoi(changes[i%len(changes)])
		curFreq = curFreq + change
		if freqs[curFreq] {
			break
		}
	}

	return curFreq
}

func main() {
	fileContent, err := ioutil.ReadFile("day1/day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 1 part 1 result: %d\n", calculateFrequency(strings.Split(string(fileContent), "\n")))

	fmt.Printf("Day 1 part 2 result: %d\n", findFirstDuplicateFrequency(strings.Split(string(fileContent), "\n")))
}
