package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	part1 := 0
	part2 := ""
	for _, ln := range strings.Split(input, "\n") {
		sectorIdx := strings.LastIndex(ln, "-") + 1
		s := strings.Split(ln[sectorIdx:], "[")
		sector, expectedChecksum := s[0], strings.Trim(s[1], "]")
		name := ln[:sectorIdx-1]
		if expectedChecksum == checksum(name) {
			i, err := strconv.Atoi(sector)
			if err != nil {
				panic(err)
			}
			part1 += i
			decrypted := decrypt(name, i)
			if strings.Contains(decrypted, "north") {
				part2 = sector
			}
		}
	}
	return strconv.Itoa(part1), part2
}

func checksum(s string) string {
	letters := map[rune]int{}
	for _, r := range s {
		if r != '-' {
			if _, found := letters[r]; found {
				letters[r]++
			} else {
				letters[r] = 1
			}
		}
	}

	type kv = struct {
		key   rune
		value int
	}
	var pairs []kv
	for k, v := range letters {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].value == pairs[j].value {
			return pairs[i].key < pairs[j].key
		}
		return pairs[i].value > pairs[j].value
	})
	res := ""
	for i, p := range pairs {
		if i >= 5 {
			break
		}
		res += string(p.key)
	}
	return res
}

func decrypt(name string, sectorId int) string {
	res := ""
	shift := (sectorId%26 + 26) % 26 // 'compress' to between 0,25
	for _, r := range name {
		if r == '-' {
			res += " "
		} else {
			if int(r)+shift > 'z' {
				res += string('a' + int(r) + shift - 'z' - 1)
			} else {
				res += string(int(r) + shift)
			}
		}
	}
	return res
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
