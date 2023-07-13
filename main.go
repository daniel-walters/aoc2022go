package main

import (
	"aoc/day01"
	"aoc/utils"
	"fmt"
)

func main() {
	slnOne, slnTwo := day01.Main()
	fmt.Printf("%sPart 1: %s%d\n%sPart 2: %s%d\n", utils.Green, utils.Reset, slnOne, utils.Green, utils.Reset, slnTwo)
}
