package task0166

import "strconv"

// Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.
// If the fractional part is repeating, enclose the repeating part in parentheses.
// If multiple answers are possible, return any of them.
// It is guaranteed that the length of the answer string is less than 10^4 for all the given inputs.
//
// Constraints:
//   -2^31 <= numerator, denominator <= 2^31 - 1
//   denominator != 0
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Fraction to Recurring Decimal.
//   Memory Usage: 2.2 MB, less than 89.66% of Go online submissions for Fraction to Recurring Decimal.
func fractionToDecimal(x int, y int) string {
	if x == x/y*y {
		// Check for non-fractional result
		return strconv.Itoa(x / y)
	}

	// Get rid of "-"
	result := ""
	if x < 0 {
		x = -x
		result = "-"
	}
	if y < 0 {
		y = -y
		if result == "" {
			result = "-"
		} else {
			result = ""
		}
	}
	// Append integer part of result
	result += strconv.Itoa(x / y)
	resultFractional := []byte{}

	repeats := map[int][2]int{}
	isCycle, i, frac := false, 0, x-x/y*y
	var integer int
	for frac > 0 {
		// here we have frac < y. Get the first number of "integer = frac*10/y"
		frac *= 10
		integer = frac / y
		if repeat, ok := repeats[frac]; ok && repeat[0] == integer {
			// Cycle check
			i = repeat[1]
			isCycle = true
			break
		}
		repeats[frac] = [2]int{integer, i}
		i++

		if integer > 0 {
			frac = frac - y*integer
		}
		resultFractional = append(resultFractional, byte('0'+integer))
	}
	if isCycle {
		return result + "." + string(resultFractional[:i]) + "(" + string(resultFractional[i:]) + ")"
	}
	return result + "." + string(resultFractional)
}
