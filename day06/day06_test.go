package day06_test

import (
	"aoc/day06"
	"testing"
)

type inputs struct {
	inputs   []string
	expected []int
}

var testsOne = inputs{
	inputs: []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	},
	expected: []int{7, 5, 6, 10, 11},
}

func TestPartOne(t *testing.T) {
	for i := 0; i < 5; i++ {
		input := testsOne.inputs[i]
		expected := testsOne.expected[i]
		actual := day06.GetPacketMarker(input, 4)
		if expected != actual {
			t.Errorf("Expected: %d, Receieved: %d", expected, actual)
		}
	}

}

var testsTwo = inputs{
	inputs: []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	},
	expected: []int{19, 23, 23, 29, 26},
}

func TestPartTwo(t *testing.T) {
	for i := 0; i < 5; i++ {
		input := testsTwo.inputs[i]
		expected := testsTwo.expected[i]
		actual := day06.GetPacketMarker(input, 14)
		if expected != actual {
			t.Errorf("Expected: %d, Receieved: %d", expected, actual)
		}
	}
}
