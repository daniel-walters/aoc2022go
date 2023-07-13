package day01_test

import (
	"aoc/day01"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := "input_test.txt"

	expected := 24_000
	actual := day01.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Received: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	input := "input_test.txt"

	expected := 45_000
	actual := day01.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Received: %d", expected, actual)
	}
}
