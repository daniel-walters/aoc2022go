package day03_test

import (
	"aoc/day03"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	t.Skip()

	expected := 0
	actual := day03.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	t.Skip()

	expected := 0
	actual := day03.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
