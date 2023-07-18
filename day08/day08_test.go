package day08_test

import (
	"aoc/day08"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 21
	actual := day08.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 8
	actual := day08.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
