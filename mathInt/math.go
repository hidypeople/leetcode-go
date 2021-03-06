package mathInt

// Math.min
func Min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// Math.max
func Max(val1, val2 int) int {
	if val2 < val1 {
		return val1
	}
	return val2
}

// Math.abs
func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// Math.pow
func Pow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
