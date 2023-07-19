package day09_test

import (
	"aoc/day09"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 13
	actual := day09.GetPositions(input, 2)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 36
	actual := day09.GetPositions("input_two_test.txt", 10)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
