package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calculateFrequency(changes []string) (result int) {
	for _, d := range changes {
		d, _ := strconv.Atoi(d)
		result += d
	}
	return
}

func findFirstDuplicateFrequency(changes []string) (result int) {
	freqs := make(map[int]struct{})
	for i := 0; !has(freqs, result); i++ {
		freqs[result] = struct{}{}
		d, _ := strconv.Atoi(changes[i%len(changes)])
		result += d
	}
	return
}
func has(set map[int]struct{}, k int) (present bool) {
	_, present = set[k]
	return
}

func main() {
	fileContent, err := ioutil.ReadFile("01_chronal_calibration/day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 1 part 1 result: %d\n", calculateFrequency(strings.Split(string(fileContent), "\n")))

	fmt.Printf("Day 1 part 2 result: %d\n", findFirstDuplicateFrequency(strings.Split(string(fileContent), "\n")))
}
