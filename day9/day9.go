package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"bufio"
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
	score, garbageCount, depth := 0, 0, 1
	skip, inGarbage := false, false
	for _, r := range input {
		if skip {
			skip = false
			continue
		}
		if !inGarbage {
			switch r {
			case '{':
				score += depth
				depth++
			case '}':
				depth--
			case '<':
				inGarbage = true
			}
		} else if r == '>' {
			inGarbage = false
		} else if r != '!' {
			garbageCount++
		}

		if r == '!' {
			skip = true
		}
	}
	return score, garbageCount
}