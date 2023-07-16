package utils

import "strings"

func SplitOnce(str, sep string) (string, string) {
	split := strings.SplitN(str, sep, 2)

	return split[0], split[1]
}
