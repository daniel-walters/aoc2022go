package utils

import (
	"aoc/consts"
	"fmt"
	"os"
)

func PrintError(err any) {
	fmt.Printf("%s[ERROR]: %s%s\n", consts.Red, consts.Reset, err)
}

func HandleFileClose(file *os.File) {
	file.Close()

	if r := recover(); r != nil {
		PrintError(r)
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
