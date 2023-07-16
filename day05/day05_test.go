package day05_test

import (
	"aoc/day05"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	expected := "CMZ"
	actual := day05.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %s, Receieved: %s", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := "MCD"
	actual := day05.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %s, Receieved: %s", expected, actual)
	}
}
