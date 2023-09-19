package day14

import (
	"aoc/utils"
	"strings"
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

type coord struct {
	x int
	y int
}

func getRockPositions(inputFile string) map[coord]bool {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	gridMap := map[coord]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		rockPath := strings.Split(strings.ReplaceAll(line, " ", ""), "->")
		var prev coord

		for i, corner := range rockPath {
			coords := utils.StringsToInts(strings.Split(corner, ","))
			cornerCoord := coord{coords[0], coords[1]}
			gridMap[cornerCoord] = true

			if i == 0 {
				prev = cornerCoord
				continue
			}

			if prev.x == cornerCoord.x {
				largest := max(prev.y, cornerCoord.y)
				smallest := min(prev.y, cornerCoord.y)

				for i := smallest + 1; i < largest; i++ {
					newCoord := coord{prev.x, i}
					gridMap[newCoord] = true
				}
			}

			if prev.y == cornerCoord.y {
				largest := max(prev.x, cornerCoord.x)
				smallest := min(prev.x, cornerCoord.x)

				for i := smallest + 1; i < largest; i++ {
					newCoord := coord{i, prev.y}
					gridMap[newCoord] = true
				}
			}

			prev = cornerCoord
		}
	}

	return gridMap
}

func getMaxY(gridMap map[coord]bool) int {
	max := 0

	for c := range gridMap {
		if c.y > max {
			max = c.y
		}
	}

	return max
}

func canFall(curPos coord, gridMap map[coord]bool) (bool, coord) {
	downCo := coord{curPos.x, curPos.y + 1}
	leftDownCo := coord{curPos.x - 1, curPos.y + 1}
	rightDownCo := coord{curPos.x + 1, curPos.y + 1}

	_, found := gridMap[downCo]
	if !found {
		return true, downCo
	}

	_, found = gridMap[leftDownCo]
	if !found {
		return true, leftDownCo
	}

	_, found = gridMap[rightDownCo]
	if !found {
		return true, rightDownCo
	}

	return false, coord{}
}

func PartOne(inputFile string) int {
	gridMap := getRockPositions(inputFile)
	groundY := getMaxY(gridMap)
	sandAtRest := 0

KeepDropping:
	for true {
		sand := coord{500, 0}

	Drop:
		for true {
			if sand.y > groundY {
				break KeepDropping
			}

			willFall, to := canFall(sand, gridMap)

			if willFall {
				sand = to
				continue
			}

			gridMap[sand] = true
			break Drop
		}

		sandAtRest += 1
	}

	return sandAtRest
}

func PartTwo(inputFile string) int {
	gridMap := getRockPositions(inputFile)
	groundY := getMaxY(gridMap) + 2
	sandAtRest := 0

KeepDropping:
	for true {
		sand := coord{500, 0}

	Drop:
		for true {
			if gridMap[coord{500, 0}] {
				break KeepDropping
			}


			if sand.y+1 == groundY {
				gridMap[sand] = true
				break Drop
			}

			willFall, to := canFall(sand, gridMap)

			if willFall {
				sand = to
				continue
			}

			gridMap[sand] = true
			break Drop
		}

		sandAtRest += 1
	}

	return sandAtRest
}
