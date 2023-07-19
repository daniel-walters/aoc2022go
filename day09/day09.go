package day09

import (
	"aoc/utils"
	"errors"
	"fmt"
	"strconv"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

var errNotADirection = errors.New("Received an invalid direction")

func strToDir(str string) direction {
	switch str {
	case "U":
		return up
	case "R":
		return right
	case "D":
		return down
	case "L":
		return left
	}

	panic(errNotADirection)
}

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func newPoint(x, y int) point {
	return point{x, y}
}

func (p1 *point) add(p2 point) {
	p1.x += p2.x
	p1.y += p2.y
}

func (p *point) move(dir direction) {
	switch dir {
	case up:
		p.add(point{0, 1})
	case right:
		p.add(point{1, 0})
	case down:
		p.add(point{0, -1})
	case left:
		p.add(point{-1, 0})
	}
}

func (p1 *point) follow(p2 point) {
	if p1.isAdjacentTo(p2) {
		return
	}

	x := utils.Clamp(p2.x-p1.x, -1, 1)
	y := utils.Clamp(p2.y-p1.y, -1, 1)

	movementPoint := newPoint(x, y)

	p1.add(movementPoint)
}

func (p1 point) isAdjacentTo(p2 point) bool {
	return utils.Abs(p1.x-p2.x) <= 1 && utils.Abs(p1.y-p2.y) <= 1
}

type rope struct {
	knots       []point
	tailHistory map[point]bool
}

func newRope(numKnots int) rope {
	knots := []point{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots, newPoint(0, 0))
	}

	return rope{
		knots: knots,
		tailHistory: map[point]bool{
			newPoint(0, 0): true,
		},
	}
}

func (r *rope) move(dir direction) {
	r.knots[0].move(dir)

	for i := 1; i < len(r.knots); i++ {
		r.knots[i].follow(r.knots[i-1])

		if i == len(r.knots)-1 {
			r.tailHistory[r.knots[i]] = true
		}
	}
}

func Main(inputFile string) (any, any) {
	slnOne := GetPositions(inputFile, 2)
	slnTwo := GetPositions(inputFile, 10)

	return slnOne, slnTwo
}

func GetPositions(inputFile string, numKnots int) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	rope := newRope(numKnots)

	for scanner.Scan() {
		line := scanner.Text()
		dirStr, countStr := utils.SplitOnce(line, " ")

		dir := strToDir(dirStr)
		count, err := strconv.Atoi(countStr)
		if err != nil {
			panic(err)
		}

		for i := 0; i < count; i++ {
			rope.move(dir)
		}
	}

	return len(rope.tailHistory)
}
