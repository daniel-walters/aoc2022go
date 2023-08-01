package day11_test

import (
	"aoc/day11"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 10605
	actual := day11.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 2713310158
	actual := day11.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
