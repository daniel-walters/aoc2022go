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

func BatchScannerOutput(size int, scanner *bufio.Scanner) [][]string {
	batches := [][]string{}
	i := -1

	for j := 0; scanner.Scan(); j++ {
		if j%size == 0 {
			i++
			batches = append(batches, []string{})
		}

		line := scanner.Text()
		batches[i] = append(batches[i], line)
	}

	return batches
}
