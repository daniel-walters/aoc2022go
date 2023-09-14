package day12

import (
	"aoc/dsa"
	"aoc/utils"
	"math"
	"slices"
)

type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type gridSpace struct {
	row int
	col int
}

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

const VISITED_MARKER = rune(int('z') + 1)

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	grid := [][]rune{}
	dists := [][]int{}

	row := 0
	var start gridSpace
	var end gridSpace

	for scanner.Scan() {
		line := scanner.Text()
		gridRow := []rune{}
		distRow := []int{}

		for col, char := range []rune(line) {
			gridRow = append(gridRow, char)
			distRow = append(distRow, math.MaxInt)

			if char == 'S' {
				start = gridSpace{row, col}
			}

			if char == 'E' {
				end = gridSpace{row, col}
			}
		}

		grid = append(grid, gridRow)
		dists = append(dists, distRow)

		row += 1
	}

	dists[start.row][start.col] = 0

	steps := getShortestPath(start, end, grid, dists)

	return steps
}

func getNeighbors(space gridSpace, rowSize, colSize int) []gridSpace {
	row := space.row
	col := space.col

	neighbors := []gridSpace{
		{row + 1, col},
		{row - 1, col},
		{row, col + 1},
		{row, col - 1},
	}

	neighbors = utils.Filter(neighbors, func(s gridSpace) bool {
		if s.row == -1 || s.col == -1 || s.row > rowSize-1 || s.col > colSize-1 {
			return false
		}

		return true
	})

	return neighbors
}

func getShortestPath(start, end gridSpace, grid [][]rune, dists [][]int) int {
	grid[start.row][start.col] = 'a'
	grid[end.row][end.col] = 'z'

	queue := dsa.NewQueue[gridSpace]()
	queue.Enqueue(start)

	rowLen := len(grid)
	colLen := len(grid[0])

	for !queue.IsEmpty() {
		curr, err := queue.Dequeue()
		if err != nil {
			panic(err)
		}

		if grid[curr.row][curr.col] == VISITED_MARKER {
			continue
		}

		curHeight := getHeight(grid[curr.row][curr.col])
		neighbors := getNeighbors(curr, rowLen, colLen)

		for _, space := range neighbors {
			if grid[space.row][space.col] == VISITED_MARKER {
				continue
			}

			height := getHeight(grid[space.row][space.col])
			canTraverse := height-curHeight <= 1
			dist := dists[curr.row][curr.col] + 1

			if dists[space.row][space.col] > dist && canTraverse {
				dists[space.row][space.col] = dist
			}

			if canTraverse {
				queue.Enqueue(space)
			}
		}
		grid[curr.row][curr.col] = VISITED_MARKER
	}

	return dists[end.row][end.col]
}

func getHeight(char rune) int {
	switch char {
	case 'S':
		return 1
	case 'E':
		return 26
	default:
		return int(char) - 96
	}
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	grid := [][]rune{}
	dists := [][]int{}

	row := 0
	startSquares := []gridSpace{}
	var end gridSpace
	var ogStart gridSpace

	for scanner.Scan() {
		line := scanner.Text()
		gridRow := []rune{}
		distRow := []int{}

		for col, char := range []rune(line) {
			gridRow = append(gridRow, char)
			distRow = append(distRow, math.MaxInt)

			if char == 'S' {
				startSquares = append(startSquares, gridSpace{row, col})
				ogStart = gridSpace{row, col}
			}
			if char == 'a' {
				startSquares = append(startSquares, gridSpace{row, col})
			}

			if char == 'E' {
				end = gridSpace{row, col}
			}
		}

		grid = append(grid, gridRow)
		dists = append(dists, distRow)

		row += 1
	}

	grid[ogStart.row][ogStart.col] = 'a'

    pathLens := []int{}

	for _, start := range startSquares {
		distClone := utils.TwoDClone(dists)
		gridClone := utils.TwoDClone(grid)

		distClone[start.row][start.col] = 0

		steps := getShortestPath(start, end, gridClone, distClone)

        pathLens = append(pathLens, steps)
	}

	return slices.Min(pathLens)
}
