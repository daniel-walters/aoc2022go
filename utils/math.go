package utils

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

func Clamp(num, min, max int) int {
	switch {
	case num < min:
		return min
	case num > max:
		return max
	default:
		return num
	}
}
