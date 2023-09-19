package day14_test

import (
	"aoc/day14"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 24
	actual := day14.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 93
	actual := day14.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
