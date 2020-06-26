package utils

// Greatest common divisor
func GCD(a, b uint) uint {
	if a == b {
		return 1
	}
	for {
		if a == 0 {
			return b
		}
		if b == 0 {
			return a
		}

		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
