package day13_test

import (
	"aoc/day13"
	"testing"
)

const input = "input_test.txt"

func TestPartOne(t *testing.T) {
	t.Skip()
	expected := 13
	actual := day13.PartOne(input)

	if expected != actual {
		t.Errorf("Expected: %d, Received: %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 140
	actual := day13.PartTwo(input)

	if expected != actual {
		t.Errorf("Expected: %d, Received: %d", expected, actual)
	}
}
