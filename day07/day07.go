package day07

import (
	"aoc/dsa"
	"aoc/utils"
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func getPath(dirStack dsa.Stack[string]) string {
	path := ""
	for _, v := range dirStack {
		path += v
	}

	return path
}

func generateSizeMap(scanner *bufio.Scanner) map[string]int {
	dirSizes := map[string]int{}
	dirStack := dsa.NewStack[string]()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			parts := strings.Split(line, " ")
			dirName := parts[len(parts)-1]
			if dirName == ".." {
				dirStack.Pop()
			} else {
				path := getPath(dirStack) + dirName
				dirSizes[path] = 0
				dirStack.Push(path)
			}
		}

		if unicode.IsDigit(rune(line[0])) {
			size, err := strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				panic(err)
			}

			for _, v := range dirStack {
				dirSizes[v] += size
			}
		}
	}

	return dirSizes
}

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	dirSizes := generateSizeMap(scanner)
	sizes := utils.CollectMap(dirSizes)
	filtered := utils.Filter(sizes, func(n int) bool {
		return n <= 100_000
	})

	return utils.Sum(filtered...)
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	dirSizes := generateSizeMap(scanner)
	const totalSpace = 7_000_000_0
	const spaceNeeded = 3_000_000_0
	freeSpace := totalSpace - dirSizes["/"]
	minSizeToDelete := spaceNeeded - freeSpace

	sizes := utils.CollectMap(dirSizes)
	filtered := utils.Filter(sizes, func(n int) bool {
		return n >= minSizeToDelete
	})

	return utils.Min(filtered)
}
