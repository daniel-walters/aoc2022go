package day04

import (
	"aoc/utils"
	"regexp"
	"strconv"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func rangesOverlap(a, b []int) bool {
	if b[0] >= a[0] && b[1] <= a[1] {
		return true
	}

	if a[0] >= b[0] && a[1] <= b[1] {
		return true
	}

	return false
}

func rangesInclude(a, b []int) bool {
	if b[0] <= a[1] && b[1] >= a[0] {
		return true
	}

	return false
}

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	overlaps := 0
	matcher := regexp.MustCompile(`[-,]`)

	for scanner.Scan() {
		line := scanner.Text()
		sections := utils.Map(matcher.Split(line, -1), stringToInt)

		if rangesOverlap(sections[:2], sections[2:]) {
			overlaps++
		}
	}

	return overlaps
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	overlaps := 0
	matcher := regexp.MustCompile(`[-,]`)

	for scanner.Scan() {
		line := scanner.Text()
		sections := utils.Map(matcher.Split(line, -1), stringToInt)

		if rangesInclude(sections[:2], sections[2:]) {
			overlaps++
		}
	}

	return overlaps
}
