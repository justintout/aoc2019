package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		total += moduleFuel(mass)
	}
	fmt.Printf("%d", total)
}

func moduleFuel(mass int) int {
	if massFuel(mass) == 0 {
		return 0
	}
	return massFuel(mass) + moduleFuel(massFuel(mass))
}

func massFuel(mass int) int {
	return floor((mass/3)-2, 0)
}

func floor(i, floor int) int {
	if i < floor {
		return floor
	}
	return i
}
