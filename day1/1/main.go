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
		total += (mass / 3) - 2
	}
	fmt.Printf("%d", total)
}
