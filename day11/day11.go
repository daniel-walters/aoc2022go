package day11

import (
	"aoc/dsa"
	"aoc/utils"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func getItems(instructions []string) ([]dsa.Queue[int], []int, int) {
	getNumsRegex := regexp.MustCompile(`[0-9]+`)
	items := []dsa.Queue[int]{}
	inspectCount := []int{}
	modulo := 1

	for _, v := range instructions {
		lines := strings.Split(v, "\n")
		startingItems := getNumsRegex.FindAllString(lines[1], -1)
		items = append(items, dsa.NewQueueFrom(utils.StringsToInts(startingItems)))
		inspectCount = append(inspectCount, 0)
		testWords := strings.Split(lines[3], " ")
		num, err := strconv.Atoi(testWords[len(testWords)-1])
		if err != nil {
			panic(err)
		}
		modulo *= num
	}

	return items, inspectCount, modulo

}

func parseOperation(opLine string, old int) int {
	words := strings.Split(opLine, " ")
	factorString := words[len(words)-1]
	factor := 0

	if factorString == "old" {
		factor = old
	} else {
		num, err := strconv.Atoi(factorString)
		if err != nil {
			panic(err)
		}
		factor = num
	}

	operation := words[len(words)-2]

	if operation == "*" {
		return old * factor
	}
	if operation == "+" {
		return old + factor
	}

	panic("Could not parse operation")
}

func parseTest(test string, worry int) bool {
	words := strings.Split(test, " ")
	testNum, err := strconv.Atoi(words[len(words)-1])
	if err != nil {
		panic(err)
	}

	return worry%testNum == 0
}

func PartOne(inputFile string) int {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	monkeyInstructions := strings.Split(string(file), "\n\n")
	items, inspectCount, _ := getItems(monkeyInstructions)

	for j := 0; j < 20; j++ {
		for i, instructions := range monkeyInstructions {
			lines := strings.Split(instructions, "\n")

			for !items[i].IsEmpty() {
				inspectCount[i]++
				worry, err := items[i].Dequeue()
				if err != nil {
					panic(err)
				}

				newWorry := parseOperation(strings.Trim(lines[2], " "), worry)
				newWorry = newWorry / 3

				result := parseTest(strings.Trim(lines[3], " "), newWorry)
				monkey := -1

				if result {
					trueWords := strings.Split(lines[4], " ")
					num, err := strconv.Atoi(trueWords[len(trueWords)-1])
					if err != nil {
						panic(err)
					}
					monkey = num
				} else {
					falseWords := strings.Split(lines[5], " ")
					num, err := strconv.Atoi(falseWords[len(falseWords)-1])
					if err != nil {
						panic(err)
					}
					monkey = num
				}

				items[monkey].Enqueue(newWorry)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectCount)))

	return inspectCount[0] * inspectCount[1]
}

func PartTwo(inputFile string) int {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	monkeyInstructions := strings.Split(string(file), "\n\n")
	items, inspectCount, modulo := getItems(monkeyInstructions)

	for j := 0; j < 10000; j++ {
		for i, instructions := range monkeyInstructions {
			lines := strings.Split(instructions, "\n")

			for !items[i].IsEmpty() {
				inspectCount[i]++
				worry, err := items[i].Dequeue()
				if err != nil {
					panic(err)
				}

				newWorry := parseOperation(strings.Trim(lines[2], " "), worry)
				newWorry = newWorry % modulo

				result := parseTest(strings.Trim(lines[3], " "), newWorry)
				monkey := -1

				if result {
					trueWords := strings.Split(lines[4], " ")
					num, err := strconv.Atoi(trueWords[len(trueWords)-1])
					if err != nil {
						panic(err)
					}
					monkey = num
				} else {
					falseWords := strings.Split(lines[5], " ")
					num, err := strconv.Atoi(falseWords[len(falseWords)-1])
					if err != nil {
						panic(err)
					}
					monkey = num
				}

				items[monkey].Enqueue(newWorry)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectCount)))

	return inspectCount[0] * inspectCount[1]
}
