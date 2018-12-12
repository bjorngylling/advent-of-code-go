package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContent, err := ioutil.ReadFile("12_subterranean_sustainability/day12_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 12 part 1 result: %+v\n", string(fileContent))

	fmt.Printf("Day 12 part 2 result: %+v\n", nil)
}
