package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	fmt.Println("-- Day 10 --")
	fmt.Printf("part_1=%d", part1(input))
}

func readInput() string {
	file, err := os.Open("./day10_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(io.Reader(file))
	scanner.Scan()
	return scanner.Text()
}

func part1(data string) int {
	var input []int
	for _, s := range strings.Split(data, ",") {
		i, err := strconv.Atoi(s)

		if err == nil {
			input = append(input, i)
		}
	}

	lst := make([]int, 256)
	for i := range lst {
		lst[i] = i
	}
	var pos, skip int
	for _, length := range input {
		lst = reverse(lst, pos, length)
		pos += length + skip
		if pos > 256 {
			pos -= 256
		}
		skip++
	}

	return lst[0] * lst[1]
}

func rev(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func rotateL(a *[]int, steps int) {
	*a = append((*a)[steps:], (*a)[:steps]...)
}
func rotateR(a *[]int, steps int) {
	rotateL(a, len(*a)-steps)
}

func reverse(a []int, start, length int) []int {
	rotateL(&a, start)
	rev(a[:length])
	rotateR(&a, start)
	return a
}
