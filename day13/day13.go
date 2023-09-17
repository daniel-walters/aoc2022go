package day13

import (
	"aoc/utils"
	"bytes"
	"encoding/json"
	"os"
	"reflect"
)

type Outcome int

const (
	ORDERED Outcome = iota
	NOT_ORDERED
	UNKNOWN
)

func Main(inputFile string) (any, any) {
	slnOne := PartOne(inputFile)
	slnTwo := PartTwo(inputFile)

	return slnOne, slnTwo
}

func comparePair(left, right any) Outcome {
	leftArr, leftIsArr := left.([]any)
	rightArr, rightIsArr := right.([]any)
	if !leftIsArr || !rightIsArr {
		panic("Could not convert pair to slice")
	}

	for i := range leftArr {
		if i >= len(rightArr) {
			return NOT_ORDERED
		}

		leftVal := leftArr[i]
		rightVal := rightArr[i]

		leftNum, leftIsNum := leftVal.(float64)
		rightNum, rightIsNum := rightVal.(float64)

		if leftIsNum && rightIsNum {
			if leftNum > rightNum {
				return NOT_ORDERED
			} else if leftNum < rightNum {
				return ORDERED
			}

			continue
		}

		leftArr, leftIsArr := leftVal.([]any)
		rightArr, rightIsArr := rightVal.([]any)

		if leftIsArr && rightIsArr {
			outcome := comparePair(leftArr, rightArr)

			if outcome == UNKNOWN {
				continue
			}

			return outcome
		}

		if leftIsArr && rightIsNum {
			outcome := comparePair(leftArr, []any{rightNum})

			if outcome == UNKNOWN {
				continue
			}

			return outcome
		}

		if rightIsArr && leftIsNum {
			outcome := comparePair([]any{leftNum}, rightArr)

			if outcome == UNKNOWN {
				continue
			}

			return outcome
		}
	}

	if len(rightArr) > len(leftArr) {
		return ORDERED
	}

	return UNKNOWN
}

func PartOne(inputFile string) int {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	pairs := bytes.Split(file, []byte("\n\n"))
	sumOfOrderedIndicies := 0

	for i, pair := range pairs {
		var left, right any

		split := bytes.Split(pair, []byte("\n"))

		json.Unmarshal(split[0], &left)
		json.Unmarshal(split[1], &right)

		if comparePair(left, right) == ORDERED {
			sumOfOrderedIndicies += (i + 1)
		}
	}

	return sumOfOrderedIndicies
}

func PartTwo(inputFile string) int {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(file, []byte("\n"))

	arrays := utils.Map(lines, func(line []byte) any {
		var arr any
		json.Unmarshal(line, &arr)
		return arr
	})

	arrays = utils.Filter(arrays, func(arr any) bool {
		return arr != nil
	})

	var firstAdd any
	var secondAdd any

	json.Unmarshal([]byte("[[2]]"), &firstAdd)
	json.Unmarshal([]byte("[[6]]"), &secondAdd)

	arrays = append(arrays, firstAdd)
	arrays = append(arrays, secondAdd)

	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(arrays)-1-i; j++ {
			if comparePair(arrays[j], arrays[j+1]) == NOT_ORDERED {
				temp := arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = temp
			}
		}
	}

	decoderOne := 0
	decoderTwo := 0

	for i, arr := range arrays {
		if reflect.DeepEqual(arr, firstAdd) {
			decoderOne = i + 1
		}

		if reflect.DeepEqual(arr, secondAdd) {
			decoderTwo = i + 1
		}
	}

	return decoderOne * decoderTwo
}
