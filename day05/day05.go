package day05

import (
	"aoc/dsa"
	"aoc/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func initialiseStacks(input string) []dsa.Stack[string] {
	lines := strings.Split(input, "\n")
	stacksLines := lines[:len(lines)-1]
	lastLine := lines[len(lines)-1]

	stackNums := []rune(strings.ReplaceAll(lastLine, " ", ""))
	stacks := []dsa.Stack[string]{}

	stacksLines = utils.Reverse(stacksLines)

	for _, v := range stackNums {
		stack := dsa.Stack[string]{}
		idx := strings.Index(lastLine, string(v))

		for _, line := range stacksLines {
			letter := []rune(line)[idx]
			if unicode.IsLetter(letter) {
				stack.Push(string(letter))
			}
		}

		stacks = append(stacks, stack)
	}

	return stacks
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func parseInstruction(instruction string) (int, int, int) {
	reg := regexp.MustCompile(`[0-9]+`)
	nums := utils.Map(reg.FindAllString(instruction, -1), stringToInt)

	return nums[0], nums[1], nums[2]
}

func PartOne(inputFile string) string {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	initStacks, instructions := utils.SplitOnce(string(file), "\n\n")
	stacks := initialiseStacks(initStacks)

	for _, inst := range strings.Split(instructions, "\n") {
		if inst == "" {
			continue
		}

		amm, start, end := parseInstruction(inst)

		for i := 0; i < amm; i++ {
			s, err := stacks[start-1].Pop()
			if err != nil {
				panic(err)
			}
			stacks[end-1].Push(s)
		}
	}

	return getAnswer(stacks)
}

func PartTwo(inputFile string) string {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	initStacks, instructions := utils.SplitOnce(string(file), "\n\n")
	stacks := initialiseStacks(initStacks)

	for _, inst := range strings.Split(instructions, "\n") {
		if inst == "" {
			continue
		}

		amm, start, end := parseInstruction(inst)

		toMove := []string{}

		for i := 0; i < amm; i++ {
			val, err := stacks[start-1].Pop()
			if err != nil {
				panic(err)
			}

			toMove = append(toMove, val)
		}

		for _, v := range utils.Reverse(toMove) {
			stacks[end-1].Push(v)
		}
	}

	return getAnswer(stacks)
}

func getAnswer(stacks []dsa.Stack[string]) string {
	answer := ""

	for _, stack := range stacks {
		val, err := stack.Peek()
		if err != nil {
			panic(err)
		}
		answer += val
	}

	return answer
}
