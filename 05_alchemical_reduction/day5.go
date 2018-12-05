package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func reaction(p string) string {
	l := []byte(p)
	for i := 0; i < len(l)-1; i++ {
		if diffCase(l[i], l[i+1]) {
			l = append(l[:i], l[i+2:]...)
		}
	}
	return string(l)
}

func allReactions(p string) string {
	prev := ""
	for {
		p = reaction(p)
		if p == prev {
			break
		}
		prev = p
	}
	return p
}

// This optimized solution replaces reaction and allReactions.
//
// Instead of iterating over the polymer until it stabilizes it backs off to find secondary reactions
func allReactions2(p string) string {
	l := []byte(p)
	for i := 0; i < len(l)-1; i++ {
		if diffCase(l[i], l[i+1]) {
			l = append(l[:i], l[i+2:]...)
			if i == 0 {
				i -= 1
			} else {
				i -= 2
			}
		}
	}
	return string(l)
}

// Further optimized, replacing allReactions2.
//
// Instead of operating on a list "in place" and removing from it we can create a new list of polymers and only move the polymers only when
// there is no reaction.
func allReactions3(p string) string {
	l := []byte{0}
	for i := 0; i < len(p); i++ {
		v := l[len(l)-1]
		if diffCase(v, p[i]) {
			l = l[:len(l)-1]
		} else {
			l = append(l, p[i])
		}
	}
	return string(l[1:])
}

// True if a and b are the same letter with different cases
func diffCase(a, b byte) bool {
	return abs(int(a)-int(b)) == 32
}

func removePolymer(p string, del string) string {
	p = strings.Replace(p, del, "", -1)
	p = strings.Replace(p, strings.ToUpper(del), "", -1)
	return p
}

func shortestPolymer(p string) string {
	shortest := p
	for i := int('a'); i <= int('z'); i++ {
		r := allReactions3(removePolymer(p, string(i)))
		if len(r) < len(shortest) {
			shortest = r
		}
	}
	return shortest
}

func main() {
	fileContent, err := ioutil.ReadFile("05_alchemical_reduction/day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 3 part 1 result: %+v\n", len(allReactions3(string(fileContent))))

	fmt.Printf("Day 3 part 2 result: %+v\n", len(shortestPolymer(allReactions3(string(fileContent)))))
}
