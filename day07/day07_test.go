package day07_test

import (
	"aoc/day07"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := 95437
	actual := day07.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 24933642
	actual := day07.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Receieved: %d", expected, actual)
	}
}
