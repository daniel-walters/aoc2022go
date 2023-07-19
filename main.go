package main

import (
	"aoc/consts"
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"aoc/day09"
	"aoc/day10"
	"aoc/generator"
	"aoc/utils"
	"flag"
	"fmt"
	"os"
)

type slnFunc func(input string) (any, any)

var slnMap = map[int]slnFunc{
	1:  day01.Main,
	2:  day02.Main,
	3:  day03.Main,
	4:  day04.Main,
	5:  day05.Main,
	6:  day06.Main,
	7:  day07.Main,
	8:  day08.Main,
	9:  day09.Main,
	10: day10.Main,
}

func main() {
	var runDay, genDay int
	flag.IntVar(&runDay, "day", -1, "The day to run (1 <= day <= 25)")
	flag.IntVar(&genDay, "gen", -1, "The day to generate (1 <= day <= 25)")
	flag.Parse()

	exitWithUsageIf(
		genDay == -1 && runDay == -1,
		"No arguments receieved",
	)
	exitWithUsageIf(
		genDay != -1 && runDay != -1,
		"Cannot combine arguments day and gen",
	)

	if runDay == -1 {
		processGeneratorArg(genDay)
	}

	if genDay == -1 {
		processDayArg(runDay)
	}
}

func processGeneratorArg(genDay int) {
	exitWithUsageIf(genDay < 1 || genDay > 25, "gen must be in 1 <= gen <= 25")

	generator.Generate(genDay)
	fmt.Printf("Solution template generated for day %d", genDay)
	os.Exit(0)
}

func processDayArg(runDay int) {
	exitWithUsageIf(runDay < 1 || runDay > 25, "Invalid argument received")

	fn, exists := slnMap[runDay]
	if !exists {
		utils.ExitWithError(fmt.Sprintf("Day %d may not be implemented", runDay))
	}

	inputFile := fmt.Sprintf("day%02d/input.txt", runDay)

	slnOne, slnTwo := fn(inputFile)
	printSln(slnOne, slnTwo)
}

func exitWithUsageIf(condition bool, msg string) {
	if condition {
		utils.PrintError(msg)
		flag.CommandLine.Usage()
		os.Exit(1)
	}
}

func printSln(slnOne, slnTwo any) {
	fmt.Printf("%sPart 1:\n%s%v\n%sPart 2:\n%s%v\n",
		consts.Green,
		consts.Reset,
		slnOne,
		consts.Green,
		consts.Reset,
		slnTwo,
	)
}
