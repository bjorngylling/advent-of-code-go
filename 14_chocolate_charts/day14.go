package main

import (
	"fmt"
	"strings"
)

func arrayToString(l []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(l), " ", delim, -1), "[]")
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func scores(iterations int) []int {
	scores := []int{3, 7}
	e1, e2 := 0, 1
	for i := 0; i < iterations+10; i++ {
		sum := scores[e1] + scores[e2]
		if sum >= 10 {
			scores = append(scores, sum/10, sum%10)
		} else {
			scores = append(scores, sum)
		}
		e1 = (e1 + 1 + scores[e1]) % len(scores)
		e2 = (e2 + 1 + scores[e2]) % len(scores)
	}
	return scores[iterations : iterations+10]
}

func rounds(requestedScores []int) int {
	scores := []int{3, 7}
	e1, e2 := 0, 1
	for true {
		sum := scores[e1] + scores[e2]
		if sum >= 10 {
			scores = append(scores, sum/10, sum%10)
			if len(scores) > len(requestedScores) {
				if Equal(scores[len(scores)-len(requestedScores):], requestedScores) {
					return len(scores) - len(requestedScores)
				} else if Equal(scores[len(scores)-len(requestedScores)-1:len(scores)-1], requestedScores) {
					return len(scores) - len(requestedScores) - 1
				}
			}
		} else {
			scores = append(scores, sum)
			if len(scores) > len(requestedScores) && Equal(scores[len(scores)-len(requestedScores):], requestedScores) {
				return len(scores) - len(requestedScores)
			}
		}
		e1 = (e1 + 1 + scores[e1]) % len(scores)
		e2 = (e2 + 1 + scores[e2]) % len(scores)
	}
	return -1
}

func main() {
	fmt.Printf("Day 14 part 1 result: %+v\n", arrayToString(scores(236021), ""))

	fmt.Printf("Day 14 part 2 result: %+v\n", rounds([]int{2, 3, 6, 0, 2, 1}))
}
