package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type passport struct {
	ecl string
	pid string
	eyr string
	hcl string
	byr string
	iyr string
	cid string
	hgt string
}

func solve(input string) (string, string) {
	split := strings.Split(input, "\n")
	var passports []passport
	t := passport{}
	for i, ln := range split {
		if i > 0 && split[i] == "" {
			passports = append(passports, t)
			t = passport{}
			continue
		}
		pairs := strings.Split(ln, " ")
		for _, p := range pairs {
			if len(p) < 1 {
				continue
			}
			s := strings.Split(p, ":")
			field, val := s[0], s[1]
			switch field {
			case "ecl":
				t.ecl = val
				break
			case "pid":
				t.pid = val
				break
			case "eyr":
				t.eyr = val
				break
			case "hcl":
				t.hcl = val
				break
			case "byr":
				t.byr = val
				break
			case "iyr":
				t.iyr = val
				break
			case "cid":
				t.cid = val
				break
			case "hgt":
				t.hgt = val
				break
			}
		}
	}
	passports = append(passports, t)

	solution1 := 0
	for _, p := range passports {
		if p.ecl != "" && p.pid != "" && p.eyr != "" && p.hcl != "" && p.byr != "" && p.iyr != "" && p.hgt != "" {
			solution1++
		}
	}

	solution2 := 0
	for _, p := range passports {
		if p.ecl != "" && p.pid != "" && p.eyr != "" && p.hcl != "" && p.byr != "" && p.iyr != "" && p.hgt != "" {
			if !strings.Contains("amb blu brn gry grn hzl oth", p.ecl) {
				continue
			}
			if len(p.pid) != 9 {
				continue
			}
			if expYear, err := strconv.Atoi(p.eyr); err != nil || expYear < 2020 || expYear > 2030 {
				continue
			}
			if p.hcl[0] != '#' {
				continue
			}
			if _, err := hex.DecodeString(p.hcl[1:]); err != nil {
				continue
			}
			if birthYear, err := strconv.Atoi(p.byr); err != nil || birthYear < 1920 || birthYear > 2002 {
				continue
			}
			if issuingYear, err := strconv.Atoi(p.iyr); err != nil || issuingYear < 2010 || issuingYear > 2020 {
				continue
			}
			if strings.HasSuffix(p.hgt, "cm") {
				cm, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm"))
				if err != nil || cm < 150 || cm > 193 {
					continue
				}
			} else if strings.HasSuffix(p.hgt, "in") {
				in, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in"))
				if err != nil || in < 59 || in > 76 {
					continue
				}
			} else {
				continue
			}
			solution2++
		}
	}

	return fmt.Sprint(solution1), fmt.Sprint(solution2)
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
