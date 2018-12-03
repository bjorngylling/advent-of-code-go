package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("03_inventory_management_system/day3_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 3 part 1 result: %d\n", part1(strings.Split(string(fileContent), "\n")))

	fmt.Printf("Day 3 part 2 result: %s\n", part2(strings.Split(string(fileContent), "\n")))
}
