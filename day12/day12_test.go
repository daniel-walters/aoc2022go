package day12_test

import (
	"aoc/day12"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 31
	actual := day12.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Received: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 29
	actual := day12.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Received: %d", expected, actual)
	}
}
