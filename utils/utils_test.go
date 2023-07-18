package utils_test

import (
	"aoc/utils"
	"bufio"
	"reflect"
	"strings"
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
			res := utils.Max(test.input[:])
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

func TestBatchScannerOutput(t *testing.T) {
	scanner := bufio.NewScanner(
		strings.NewReader("foo bar baz bar foo bar baz bar foo"),
	)
	scanner.Split(bufio.ScanWords)

	wants := [][]string{
		{"foo", "bar", "baz"},
		{"bar", "foo", "bar"},
		{"baz", "bar", "foo"},
	}
	actual := utils.BatchScannerOutput(3, scanner)

	if !reflect.DeepEqual(wants, actual) {
		t.Errorf("got %v, wants %v", actual, wants)
	}
}

func TestMap(t *testing.T) {
	slice := []int{2, 4, 6, 8}
	mappingFn := func(num int) int {
		return num + 2
	}
	wants := []int{4, 6, 8, 10}
	actual := utils.Map(slice, mappingFn)

	if !reflect.DeepEqual(wants, actual) {
		t.Errorf("got %v, wants %v", actual, wants)
	}
}
