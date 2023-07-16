package utils

import "strings"

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Sum(a ...int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func Map[T, K any](slice []T, pred func(item T) K) []K {
	newSlice := []K{}

	for _, v := range slice {
		mappedValue := pred(v)
		newSlice = append(newSlice, mappedValue)
	}

	return newSlice
}

func SplitOnce(str, sep string) (string, string) {
	split := strings.SplitN(str, sep, 2)

	return split[0], split[1]
}

func Reverse[T any](slice []T) []T {
	reversed := []T{}

	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}
