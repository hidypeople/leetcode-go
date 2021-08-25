package task0633

import "math"

// Given a non-negative integer c, decide whether there're two integers a and b such that a^2 + b^2 = c.
//
// Constraints:
//   0 <= c <= 2^31 - 1
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Square Numbers.
//   Memory Usage: 2 MB, less than 57.14% of Go online submissions for Sum of Square Numbers.
func judgeSquareSum(c int) bool {
	if c < 3 {
		return true
	}
	minSqrt := int(math.Sqrt(float64(c)))
	for i := minSqrt; i >= 1; i-- {
		target := c - i*i
		targetSqrt := int(math.Sqrt(float64(target)))
		if target == targetSqrt*targetSqrt {
			return true
		}
	}
	return false
}
