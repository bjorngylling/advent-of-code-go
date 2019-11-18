package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	PLAYERCOUNTFIELD     = 0
	LASTMARBLEVALUEFIELD = 6
)

func parseLn(ln string) (playerCount, lastMarbleValue int) {
	f := strings.Fields(ln)
	playerCount, _ = strconv.Atoi(f[PLAYERCOUNTFIELD])
	lastMarbleValue, _ = strconv.Atoi(f[LASTMARBLEVALUEFIELD])
	return
}

func score(playerCount, lastMarbleValue int) int {
	scores := make([]int, playerCount)
	l := ring.New(1)
	l.Value = 0
	for turn := 1; turn <= lastMarbleValue; turn++ {
		if turn%23 != 0 {
			n := ring.New(1)
			n.Value = turn
			l = l.Next().Link(n).Prev()
		} else {
			l = l.Move(-8)
			n := l.Unlink(1)
			scores[turn%playerCount] += n.Value.(int) + turn
			l = l.Next()
		}
	}
	highscore := 0
	for _, s := range scores {
		if s > highscore {
			highscore = s
		}
	}
	return highscore
}

func main() {
	fileContent, err := ioutil.ReadFile("2018/09_marble_mania/day9_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	playerCount, lastMarbleValue := parseLn(string(fileContent))
	fmt.Printf("Day 9 part 1 result: %+v\n", score(playerCount, lastMarbleValue))

	fmt.Printf("Day 9 part 2 result: %+v\n", score(playerCount, lastMarbleValue*100))
}
