package main

import (
	"fmt"
	"github.com/justintout/aoc2019/day2"
)

// type instruction struct {
// 	address int
// 	opcode     int
// 	parameters []int
// }

// func (i instruction) nextAddress() int {
// 	if (i.opcode == 99) {
// 		return i.address
// 	}
// 	return i.address + len(i.parameters)
// }

// "Each of the two input values will be between 0 and 99, inclusive."
var max = 99

func main() {
	var program string
	_, err := fmt.Scanln(&program)
	if err != nil {
		panic("failed to read program")
	}
	noun, verb := find(19690720, program)
	fmt.Printf("%d", 100*noun+verb)
}

func find(target int, program string) (int, int) {
	resultChan := make(chan [2]int)
	for noun := 0; noun < max; noun++ {
		for verb := 0; verb < max; verb++ {
			go func(noun, verb int) {
				initial := day2.Parse(program)
				output := day2.Run(noun, verb, initial)
				if output == target {
					resultChan <- [2]int{noun, verb}
					close(resultChan)
				}
			}(noun, verb)
		}
	}
	result := <-resultChan
	return result[0], result[1]
}
