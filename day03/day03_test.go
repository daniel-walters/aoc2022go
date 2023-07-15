package day03_test

import (
	"aoc/day03"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 157
	actual := day03.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 70
	actual := day03.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
