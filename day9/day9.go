package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day9/day9_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(io.Reader(file))
	scanner.Scan()
	input := scanner.Text()
	groupScore, garbageCount := process(input)
	fmt.Println("-- Day 9 --")
	fmt.Printf("group_score=%d,garbage_count=%d\n", groupScore, garbageCount)
}

func process(input string) (int, int) {
	var score, garbageCount, depth int
	var skip, inGarbage bool
	for _, r := range input {
		switch {
		case skip:
			skip = false
			continue
		case !inGarbage && r == '{':
			depth++
			score += depth
		case !inGarbage && r == '}':
			depth--
		case !inGarbage && r == '<':
			inGarbage = true
		case inGarbage && r == '>':
			inGarbage = false
		case inGarbage && r != '!':
			garbageCount++
		}

		skip = r == '!'
	}
	return score, garbageCount
}
