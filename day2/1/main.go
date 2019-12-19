package main

import (
	"fmt"
	"github.com/justintout/aoc2019/day2"
)

func main() {
	// assuming program comes in as single line, comma separated, no spaces
	var program string
	_, err := fmt.Scanln(&program)
	if err != nil {
		panic("failed to read program")
	}
	tokens := day2.Parse(program)
	// "...before running the program, replace position 1 with
	// the value 12 and replace position 2 with the value 2."
	result := day2.Run(12, 2, tokens)
	fmt.Printf("%d", result)
}


