package main

import (
	"aoc/consts"
	"aoc/day01"
	"aoc/utils"
	"flag"
	"fmt"
	"os"
)

var slnMap = map[int]func(input string) (int, int){
	1: day01.Main,
}

func main() {
	var day int
	flag.IntVar(&day, "day", -1, "The day to run (1 <= day <= 25)")
	flag.Parse()

	if day < 1 || day > 25 {
		utils.PrintError("Invalid argument received")
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	fn, exists := slnMap[day]
	if !exists {
		utils.PrintError(fmt.Sprintf("Day %d may not be implemented", day))
		os.Exit(1)
	}

	inputFile := fmt.Sprintf("day%02d/input.txt", day)

	slnOne, slnTwo := fn(inputFile)
	printSln(slnOne, slnTwo)
}

func printSln(slnOne, slnTwo int) {
	fmt.Printf("%sPart 1: %s%d\n%sPart 2: %s%d\n",
		consts.Green,
		consts.Reset,
		slnOne,
		consts.Green,
		consts.Reset,
		slnTwo,
	)
}
