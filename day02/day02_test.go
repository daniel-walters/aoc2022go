package day02_test

import (
	"aoc/day02"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 15
	actual := day02.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 12
	actual := day02.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
