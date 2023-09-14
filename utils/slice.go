package utils

import (
	"math"
	"slices"
	"strconv"
)

func Max(slice []int) int {
	max := math.MinInt
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

func Min(slice []int) int {
	min := math.MaxInt
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min
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

func Reverse[T any](slice []T) []T {
	reversed := []T{}

	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

func CollectMap[T, K comparable](input map[T]K) []K {
	slice := []K{}
	for _, v := range input {
		slice = append(slice, v)
	}

	return slice
}

func Filter[T comparable](slice []T, pred func(item T) bool) []T {
	filtered := []T{}
	for _, v := range slice {
		if pred(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func Multiply(slice []int) int {
	product := 1
	for _, v := range slice {
		product *= v
	}

	return product
}

func StringsToInts(slice []string) []int {
	return Map(slice, func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})
}

func TwoDClone[T any](slice [][]T) [][]T {
    cloned := [][]T{}

    for _, row := range slice {
        cloned = append(cloned, slices.Clone(row))
    }

    return cloned
}
