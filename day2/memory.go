package day2

import (
	"strconv"
	"strings"
)

// Parse takes an Intcode program as a single line of
// comma-separated integers. Returns a slice of ints
// from parsing this program.
func Parse(program string) []int {
	tokens := []int{}
	program = strings.TrimSuffix(program, ",")
	ts := strings.Split(program, ",")
	for _, t := range ts {
		value, err := strconv.Atoi(t)
		if err != nil {
			panic("scanned non-integer: " + t)
		}
		tokens = append(tokens, value)
	}
	return tokens
}

// Run accepts an Intcode program as a slice of ints and
// returns the value at address 0 after replacing address 1
// with noun and address 2 with verb
func Run(noun, verb int, tokens []int) int {
	tokens[1] = noun
	tokens[2] = verb
	for i := 0; i < len(tokens); i += 4 {
		if tokens[i] == 99 {
			return tokens[0]
		}
		first := tokens[i+1]
		second := tokens[i+2]
		location := tokens[i+3]
		switch tokens[i] {
		case 1:
			tokens[location] = tokens[first] + tokens[second]
		case 2:
			tokens[location] = tokens[first] * tokens[second]
		}
	}
	return tokens[0]
}

// Initialize accepts an Intcode progam as a slice of ints.
// It initializes program memory by running the program.
func Initialize(tokens []int) []int {
	for i := 0; i < len(tokens); i += 4 {
		if tokens[i] == 99 {
			return tokens
		}
		first := tokens[i+1]
		second := tokens[i+2]
		location := tokens[i+3]
		switch tokens[i] {
		case 1:
			tokens[location] = tokens[first] + tokens[second]
		case 2:
			tokens[location] = tokens[first] * tokens[second]
		}
	}
	return tokens
}
