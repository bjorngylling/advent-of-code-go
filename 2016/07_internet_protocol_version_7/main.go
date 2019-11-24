package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve(input string) (string, string) {
	tlsIpCount := 0
	for _, ln := range strings.Split(input, "\n") {
		if tlsSupport(ln) {
			tlsIpCount++
		}
	}
	sslIpCount := 0
	for _, ln := range strings.Split(input, "\n") {
		if sslSupport(ln) {
			sslIpCount++
		}
	}
	return strconv.Itoa(tlsIpCount), strconv.Itoa(sslIpCount)
}

func sslSupport(addr string) bool {
	supernet, hypernet := parse(addr)

	var abaList []string
	for _, seq := range supernet {
		abaList = append(abaList, aba(seq)...)
	}
	for _, aba := range abaList {
		bab := abaToBab(aba)
		for _, h := range hypernet {
			if strings.Contains(h, bab) {
				return true
			}
		}
	}

	return false
}

func tlsSupport(addr string) bool {
	supernet, hypernet := parse(addr)
	abbaCountInSupernet := 0
	for _, seq := range supernet {
		if abba(seq) {
			abbaCountInSupernet++
		}
	}
	abbaCountInHypernet := 0
	for _, seq := range hypernet {
		if abba(seq) {
			abbaCountInHypernet++
		}
	}
	return abbaCountInSupernet > 0 && abbaCountInHypernet == 0
}

func abba(seq string) bool {
	for i := 0; i < len(seq)-3; i++ {
		if seq[i] == seq[i+3] && seq[i+1] == seq[i+2] && seq[i] != seq[i+1] {
			return true
		}
	}
	return false
}

func aba(seq string) []string {
	var res []string
	for i := 0; i < len(seq)-2; i++ {
		if seq[i] == seq[i+2] && seq[i] != seq[i+1] {
			res = append(res, seq[i:i+3])
		}
	}
	return res
}

func abaToBab(aba string) string {
	return string([]uint8{aba[1], aba[0], aba[1]})
}

func parse(addr string) ([]string, []string) {
	var supernet, hypernet []string
	inHypernet := false
	var buf []rune
	for _, c := range addr {
		switch c {
		case '[':
			if len(buf) > 0 {
				if inHypernet {
					hypernet = append(hypernet, string(buf))
				} else {
					supernet = append(supernet, string(buf))
				}
			}
			buf = []rune{}
			inHypernet = true
		case ']':
			if len(buf) > 0 {
				hypernet = append(hypernet, string(buf))
			}
			buf = []rune{}
			inHypernet = false
		default:
			buf = append(buf, c)
		}
	}
	if len(buf) > 0 {
		supernet = append(supernet, string(buf))
	}
	return supernet, hypernet
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
