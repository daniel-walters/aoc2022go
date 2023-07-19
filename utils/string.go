package utils

import "strings"

func SplitOnce(str, sep string) (string, string) {
	split := strings.SplitN(str, sep, 2)

	if len(split) == 2 {
		return split[0], split[1]
	}

	if len(split) == 1 {
		return split[0], ""
	}

	return "", ""
}
