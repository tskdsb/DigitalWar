package utils

// Greatest common divisor
func GCD(a, b int) int {
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
