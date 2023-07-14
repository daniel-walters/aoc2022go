package main

import (
	"aoc/consts"
	"aoc/day01"
	"fmt"
)

func main() {
	slnOne, slnTwo := day01.Main()
	fmt.Printf("%sPart 1: %s%d\n%sPart 2: %s%d\n",
		consts.Green,
		consts.Reset,
		slnOne,
		consts.Green,
		consts.Reset,
		slnTwo,
	)
}
