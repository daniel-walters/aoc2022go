package utils

import (
	"aoc/consts"
	"fmt"
	"os"
)

func PrintError(err any) {
	fmt.Printf("%s[ERROR]: %s%s\n", consts.Red, consts.Reset, err)
}

func ExitWithError(err any) {
	PrintError(err)
	os.Exit(1)
}
