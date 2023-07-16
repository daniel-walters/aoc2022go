package day03

import (
	"aoc/utils"
	"errors"
	"unicode"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func mapFromChars(chars []rune) map[rune]bool {
	boolMap := make(map[rune]bool, len(chars))

	for _, v := range chars {
		boolMap[v] = true
	}

	return boolMap
}

func charIsInMap(chars []rune, charMap map[rune]bool) (rune, bool) {
	for _, v := range chars {
		if charMap[v] {
			return v, true
		}
	}

	return 0, false
}

func getAllCharsInMap(chars []rune, charMap map[rune]bool) []rune {
	matches := []rune{}

	for _, v := range chars {
		if charMap[v] {
			matches = append(matches, v)
		}
	}

	return matches
}

func charToPriority(char rune) int {
	charValue := unicode.ToLower(char) - 96

	if unicode.IsUpper(char) {
		charValue += 26
	}

	return int(charValue)
}

var ErrNotFound = errors.New("No matching char in map")

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune(line)
		mid := len(chars) / 2
		left := chars[:mid]
		right := chars[mid:]

		leftMap := mapFromChars(left)

		matchingChar, found := charIsInMap(right, leftMap)
		if !found {
			panic(ErrNotFound)
		}

		charValue := unicode.ToLower(matchingChar) - 96

		if unicode.IsUpper(matchingChar) {
			charValue += 26
		}

		sum += charToPriority(matchingChar)
	}

	return sum
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	batched := utils.BatchScannerOutput(3, scanner)
	sum := 0

	for _, v := range batched {
		mapCharsInOne := mapFromChars([]rune(v[0]))
		charsInOneAndTwo := getAllCharsInMap([]rune(v[1]), mapCharsInOne)
		mapCharsInOneAndTwo := mapFromChars(charsInOneAndTwo)

		charInAll, found := charIsInMap([]rune(v[2]), mapCharsInOneAndTwo)
		if !found {
			panic(ErrNotFound)
		}

		sum += charToPriority(charInAll)
	}

	return sum
}
