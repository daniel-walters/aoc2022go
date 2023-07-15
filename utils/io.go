package utils

import (
	"bufio"
	"os"
)

func GetLineScanner(fileName string) (*bufio.Scanner, func() error) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file), file.Close
}
