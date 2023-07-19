package day10_test

import (
	"aoc/day10"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	t.Skip()
	expected := 13140
	actual := day10.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	day10.PartTwo(input)
}
