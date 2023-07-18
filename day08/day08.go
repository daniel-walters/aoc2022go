package day08

import (
	"aoc/utils"
	"bufio"
	"strconv"
	"strings"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func isEdge(row, col, numRows, numCols int) bool {
	return row == 0 || row == numRows || col == 0 || col == numCols
}

func isVisible(trees [][]int, row, col int) bool {
	if isEdge(row, col, len(trees)-1, len(trees[0])-1) {
		return true
	}

	tree := trees[row][col]
	numRows := len(trees)
	numCols := len(trees[0])

	//Up
	for i := row - 1; i >= 0; i-- {
		if i == 0 && trees[i][col] < tree {
			return true
		}
		if trees[i][col] >= tree {
			break
		}
	}
	//Right
	for i := col + 1; i < numCols; i++ {
		if i == numCols-1 && trees[row][i] < tree {
			return true
		}
		if trees[row][i] >= tree {
			break
		}
	}
	//Down
	for i := row + 1; i < numRows; i++ {
		if i == numRows-1 && trees[i][col] < tree {
			return true
		}
		if trees[i][col] >= tree {
			break
		}
	}
	//Left
	for i := col - 1; i >= 0; i-- {
		if i == 0 && trees[row][i] < tree {
			return true
		}
		if trees[row][i] >= tree {
			break
		}
	}

	return false
}

func createArray(fileScanner *bufio.Scanner) [][]int {
	arr := [][]int{}

	for fileScanner.Scan() {
		chars := strings.Split(fileScanner.Text(), "")
		inner := []int{}
		for _, c := range chars {
			num, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}

			inner = append(inner, num)
		}

		arr = append(arr, inner)
	}

	return arr
}

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()
	arr := createArray(scanner)

	numVisible := 0

	for rowNum, row := range arr {
		for colNum := range row {
			if isVisible(arr, rowNum, colNum) {
				numVisible++
			}
		}
	}

	return numVisible
}

func getScore(trees [][]int, row, col int) int {
	tree := trees[row][col]
	numRows := len(trees)
	numCols := len(trees[0])
	scores := []int{}

	if isEdge(row, col, numRows-1, numCols-1) {
		return 0
	}

	//Up
	score := 0
UpLoop:
	for i := row - 1; i >= 0; i-- {
		switch {
		case trees[i][col] < tree:
			score++
		case trees[i][col] == tree:
			score++
			break UpLoop
		case trees[i][col] > tree:
			score++
			break UpLoop
		}
	}
	if score > 0 {
		scores = append(scores, score)
		score = 0
	}
	//Right
RightLoop:
	for i := col + 1; i < numCols; i++ {
		switch {
		case trees[row][i] < tree:
			score++
		case trees[row][i] == tree:
			score++
			break RightLoop
		case trees[row][i] > tree:
			score++
			break RightLoop
		}
	}
	if score > 0 {
		scores = append(scores, score)
		score = 0
	}
	//Down
DownLoop:
	for i := row + 1; i < numRows; i++ {
		switch {
		case trees[i][col] < tree:
			score++
		case trees[i][col] == tree:
			score++
			break DownLoop
		case trees[i][col] > tree:
			score++
			break DownLoop
		}
	}
	if score > 0 {
		scores = append(scores, score)
		score = 0
	}
	//Left
LeftLoop:
	for i := col - 1; i >= 0; i-- {
		switch {
		case trees[row][i] < tree:
			score++
		case trees[row][i] == tree:
			score++
			break LeftLoop
		case trees[row][i] > tree:
			score++
			break LeftLoop
		}
	}
	if score > 0 {
		scores = append(scores, score)
	}

	return utils.Multiply(scores)
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()
	arr := createArray(scanner)

	scores := []int{}

	for rowNum, row := range arr {
		for colNum := range row {
			score := getScore(arr, rowNum, colNum)
			scores = append(scores, score)
		}
	}

	return utils.Max(scores)
}
