package task0029

// Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
// Return the quotient after dividing dividend by divisor.
// The integer division should truncate toward zero, which means losing its fractional part. For example,
// truncate(8.345) = 8 and truncate(-2.7335) = -2.
// Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed
// integer range: [−2^31, 2^31 − 1]. For this problem, assume that your function returns 2^31 − 1 when
// the division result overflows.
//
// Constraints:
//   -2^31 <= dividend, divisor <= 2^31 - 1
//   divisor != 0
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Divide Two Integers.
//   Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Divide Two Integers.
func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	result := 0
	isMinus := (dividend < 0 && divisor >= 0) || (dividend >= 0 && divisor < 0)

	d := abs(divisor)
	rest := abs(dividend)
	for rest >= d {
		curr := d
		res := 1
		for curr < rest>>1 {
			curr <<= 1
			res <<= 1
		}
		result += res
		rest = rest - curr
	}

	if isMinus {
		return -result
	} else {
		MaxUint := ^uint32(0)
		MaxInt := int(MaxUint >> 1)
		if result > MaxInt {
			return MaxInt
		}
		return result
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
