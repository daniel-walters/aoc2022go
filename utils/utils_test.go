package utils_test

import (
	"aoc/utils"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		name  string
		input [2]int
		want  int
	}{
		{"should return larger for positive numbers", [2]int{4, 5}, 5},
		{"should return larger for negative numbers", [2]int{-4, -5}, -4},
		{"should return larger for mixed numbers", [2]int{4, -5}, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := utils.Max(test.input[0], test.input[1])
			if res != test.want {
				t.Errorf("got %d, wants %d", res, test.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	wants := 15
	actual := utils.Sum(10, 3, 4, -2)

	if wants != actual {
		t.Errorf("got %d, wants %d", actual, wants)
	}
}
