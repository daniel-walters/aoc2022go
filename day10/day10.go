package day10

import (
	"aoc/utils"
	"strconv"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	x := 1
	completedCycles := 0

	signalStrengths := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		instruction, value := utils.SplitOnce(line, " ")

		currentCycles := []int{completedCycles + 1}
		if instruction != "noop" {
			currentCycles = append(currentCycles, completedCycles+2)
		}

		for _, v := range currentCycles {
			if v == 20 || (v-20)%40 == 0 {
				signalStrengths = append(signalStrengths, v*x)
				break
			}
		}

		if instruction == "noop" {
			completedCycles += 1
			continue
		}

		val, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		completedCycles += 2
		x += val
	}

	return utils.Sum(signalStrengths...)
}

func PartTwo(inputFile string) string {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	x := 1
	completedCycles := 0
	crt := [6][40]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		instruction, value := utils.SplitOnce(line, " ")

		currentCycles := []int{completedCycles + 1}
		if instruction != "noop" {
			currentCycles = append(currentCycles, completedCycles+2)
		}

		for _, v := range currentCycles {
			crtRow := (v - 1) / 40
			crtCol := (v - 1) % 40
			if crtCol == x-1 || crtCol == x || crtCol == x+1 {
				crt[crtRow][crtCol] = true
			}
		}

		if instruction == "noop" {
			completedCycles += 1
			continue
		}

		val, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		completedCycles += 2
		x += val
	}

	string := ""
	for _, row := range crt {
		for i, v := range row {
			if v {
				string += "#"
			} else {
				string += "."
			}
			if i == len(row)-1 {
				string += "\n"
			}
		}
	}

	return string
}
