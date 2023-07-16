package day04_test

import (
	"aoc/day04"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 2
	actual := day04.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 4
	actual := day04.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
