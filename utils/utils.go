package utils

import (
	"aoc/consts"
	"fmt"
	"os"
)

func HandleFileClose(file *os.File) {
	file.Close()

	if r := recover(); r != nil {
		fmt.Printf("%s[ERROR]: %s%s\n", consts.Red, consts.Reset, r)
		os.Exit(1)
	}
}

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

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
