package main

import (
	"crypto/md5"
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"strconv"
	"strings"
	"time"
)

type password [8]rune

func solve(input string) (string, string) {
	isTerm := isatty.IsTerminal(os.Stdout.Fd())
	part1 := ""
	part2 := password{'_', '_', '_', '_', '_', '_', '_', '_'}
	for i := 0; len(part1) < 8 || !part2.done(); i++ {
		output := false
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if strings.HasPrefix(hash, "00000") {
			if len(part1) < 8 {
				part1 += string(hash[5])
				output = true
			}
			pos, err := strconv.Atoi(string(hash[5]))
			if err == nil && pos < 8 && part2[pos] == '_' {
				part2[pos] = rune(hash[6])
				output = true
			}
		}
		if !output && i%200 == 0 {
			output = true
		}
		if isTerm && output {
			p2 := part2.stringify()
			for i, c := range p2 {
				if c == '_' {
					p2 = p2[:i] + string(hash[i]) + p2[i+1:]
				}
			}
			printProgress(part1+hash[:8-len(part1)], p2)
		}
	}
	return part1, part2.stringify()
}

func (p password) stringify() string {
	res := ""
	for i := 0; i < 8; i++ {
		res += string(p[i])
	}
	return res
}
func (p password) done() bool {
	for i := 0; i < 8; i++ {
		if p[i] == '_' {
			return false
		}
	}
	return true
}

func printProgress(part1 string, part2 string) {
	fmt.Printf("\rPart 1: %s%sPart 2: %s", part1, strings.Repeat(" ", 14-len(part1)), part2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	if !isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	} else {
		fmt.Println()
	}
	fmt.Printf("Program took %s", elapsed)
}
