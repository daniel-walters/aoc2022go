package day06

import "os"

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func PartOne(inputFile string) int {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return GetPacketMarker(string(input), 4)
}

func PartTwo(inputFile string) int {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return GetPacketMarker(string(input), 14)
}

func GetPacketMarker(input string, distinctChars int) int {
	chars := []rune(input)

	for i := 0; i < len(chars)-distinctChars; i++ {
		charMap := map[rune]bool{}
		for j := 0; j < distinctChars; j++ {
			curChar := chars[i+j]
			if charMap[curChar] {
				break
			}
			charMap[curChar] = true
		}

		if len(charMap) == distinctChars {
			return i + distinctChars
		}
	}

	return -1
}
