package task0069

// Given a non-negative integer x, compute and return the square root of x.
// Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
// Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.
//
// Constraints:
//   0 <= n <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Sqrt(x).
//   Memory Usage: 2.3 MB, less than 16.40% of Go online submissions for Sqrt(x).
func mySqrt(n int) int {
	// Newton method for f(x) = x^2 - n = 0
	if n < 2 {
		return n
	}
	// small = 2 * mySqrt(n / 4),        e.g. n = 34 -> small = 2 * mySqrt(8)     = 4
	// large = 2 * mySqrt(n / 4) + 1     e.g. n = 34 -> large = 2 * mySqrt(8) + 1 = 5 this is isqrt(34)
	smallCandidate := mySqrt(n>>2) << 1
	largeCandidate := smallCandidate + 1

	println(n, smallCandidate)
	if largeCandidate*largeCandidate > n {
		return smallCandidate
	} else {
		return largeCandidate
	}
}

// The same as mySqrt but not recursive
func mySqrt_nonRecursive(n int) int {
	shift := 2
	nShifted := n >> shift
	for nShifted != 0 && nShifted != n {
		shift = shift + 2
		nShifted = n >> shift
	}
	shift = shift - 2
	result := 0
	for shift >= 0 {
		result = result << 1
		candidateResult := result + 1
		if candidateResult*candidateResult <= n>>shift {
			result = candidateResult
		}
		shift = shift - 2
	}
	return result
}
