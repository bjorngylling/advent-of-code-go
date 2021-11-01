package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseMask(m string) (ignore int64, mask int64) {
	if m == "" {
		return 0, 0
	}

	im := strings.ReplaceAll(m, "0", "2")
	im = strings.ReplaceAll(im, "1", "0")
	im = strings.ReplaceAll(im, "X", "0")
	im = strings.ReplaceAll(im, "2", "1")
	ignore, err := strconv.ParseInt(im, 2, 64)
	if err != nil {
		panic(err)
	}

	m = strings.ReplaceAll(m, "X", "0")
	mask, err = strconv.ParseInt(m, 2, 64)
	if err != nil {
		panic(err)
	}

	return ignore, mask
}

func parseMaskPart2(m string) (ignore int64, mask int64) {
	if m == "" {
		return 0, 0
	}

	im := strings.ReplaceAll(m, "1", "0")
	im = strings.ReplaceAll(im, "X", "1")
	ignore, err := strconv.ParseInt(im, 2, 64)
	if err != nil {
		panic(err)
	}

	m = strings.ReplaceAll(m, "X", "1")
	mask, err = strconv.ParseInt(m, 2, 64)
	if err != nil {
		panic(err)
	}

	return ignore, mask
}

func applyMask(mask string, value int64) int64 {
	ignore, m := parseMask(mask)
	value = value | m
	return value &^ ignore
}

func part1(lines []string) map[int64]int64 {
	mask := ""
	mem := map[int64]int64{}
	for _, ln := range lines {
		if strings.HasPrefix(ln, "mask") {
			fmt.Sscanf(ln, "mask = %s", &mask)
		} else {
			var p, i int64
			fmt.Sscanf(ln, "mem[%d] = %d", &p, &i)
			mem[p] = applyMask(mask, i)
		}
	}

	return mem
}

func part2(lines []string) map[int64]int64 {
	mask := ""
	mem := map[int64]int64{}
	for _, ln := range lines {
		if strings.HasPrefix(ln, "mask") {
			fmt.Sscanf(ln, "mask = %s", &mask)
		} else {
			var pntr, value int64
			fmt.Sscanf(ln, "mem[%d] = %d", &pntr, &value)
			ignore, m := parseMaskPart2(mask)

			pntr = pntr | m
			comb := ignore
			floating := []int64{ignore}
			for comb > 0 {
				comb = (comb - 1) & ignore
				floating = append(floating, comb)
			}
			for _, p := range floating {
				mem[pntr &^ p] = value
			}
		}
	}

	return mem
}

func sum(mem map[int64]int64) int64 {
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}

	return sum
}

func solve(input string) (string, string) {
	spl := strings.Split(input, "\n")
    return fmt.Sprintf("%d", sum(part1(spl))), fmt.Sprintf("%d", sum(part2(spl)))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
