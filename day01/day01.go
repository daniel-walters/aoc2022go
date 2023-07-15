package day01

import (
	"aoc/utils"
	"strconv"
)

func Main(input string) (int, int) {
	slnOne := PartOne(input)
	slnTwo := PartTwo(input)

	return slnOne, slnTwo
}

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	maxCalories := 0
	curCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			maxCalories = utils.Max(maxCalories, curCalories)
			curCalories = 0
			continue
		}

		lineNum, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		curCalories += lineNum
	}

	maxCalories = utils.Max(maxCalories, curCalories)

	return maxCalories
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	maxCalories := [3]int{0, 0, 0}
	curCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			bubbleUpNumber(&maxCalories, curCalories)
			curCalories = 0

			continue
		}

		lineNum, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		curCalories += lineNum
	}

	bubbleUpNumber(&maxCalories, curCalories)

	return utils.Sum(maxCalories[:]...)
}

func bubbleUpNumber(biggestThree *[3]int, numToAdd int) {
	for i, curVal := range biggestThree {
		if numToAdd < curVal {
			break
		}

		if i == 0 {
			biggestThree[0] = numToAdd
		} else {
			biggestThree[i-1] = curVal
			biggestThree[i] = numToAdd
		}
	}
}
