package main

import (
	"aoc/consts"
	"aoc/day01"
	"aoc/day02"
	"aoc/generator"
	"aoc/utils"
	"flag"
	"fmt"
	"os"
)

var slnMap = map[int]func(input string) (int, int){
	1: day01.Main,
	2: day02.Main,
}

func main() {
	var runDay, genDay int
	flag.IntVar(&runDay, "day", -1, "The day to run (1 <= day <= 25)")
	flag.IntVar(&genDay, "gen", -1, "The day to generate (1 <= day <= 25)")
	flag.Parse()

	exitWithUsage(genDay == -1 && runDay == -1, "No arguments receieved")
	exitWithUsage(genDay != -1 && runDay != -1, "Cannot combine arguments day and gen")

	if runDay == -1 {
		exitWithUsage(genDay < 1 || genDay > 25, "gen must be in 1 <= gen <= 25")

		generator.Generate(genDay)
		fmt.Printf("Solution template generated for day %d", genDay)
		os.Exit(0)
	}

	if genDay == -1 {
		exitWithUsage(runDay < 1 || runDay > 25, "Invalid argument received")

		fn, exists := slnMap[runDay]
		if !exists {
			utils.PrintError(fmt.Sprintf("Day %d may not be implemented", runDay))
			os.Exit(1)
		}

		inputFile := fmt.Sprintf("day%02d/input.txt", runDay)

		slnOne, slnTwo := fn(inputFile)
		printSln(slnOne, slnTwo)
	}
}

func exitWithUsage(condition bool, msg string) {
	if condition {
		utils.PrintError(msg)
		flag.CommandLine.Usage()
		os.Exit(1)
	}
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
