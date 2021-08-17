package task0008

// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer
// (similar to C/C++'s atoi function).
// The algorithm for myAtoi(string s) is as follows:
// Read in and ignore any leading whitespace.
// Check if the next character (if not already at the end of the string) is '-' or '+'.
// Read this character in if it is either. This determines if the final result is negative
// or positive respectively. Assume the result is positive if neither is present.
// Read in next the characters until the next non-digit charcter or the end of the input is reached.
// The rest of the string is ignored.
// Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read,
// then the integer is 0. Change the sign as necessary (from step 2).
// If the integer is out of the 32-bit signed integer range [-2^31, 2^31 - 1], then clamp the integer
// so that it remains in the range. Specifically, integers less than -2^31 should be clamped to -2^31,
// and integers greater than 2^31 - 1 should be clamped to 2^31 - 1.
// Return the integer as the final result.
// Note:
// Only the space character ' ' is considered a whitespace character.
// Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
//
// Constraints:
//   0 <= s.length <= 200
//   s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
//   Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for String to Integer (atoi).
func myAtoi(s string) int {
	MAX_INT := int32(^uint32(0) >> 1)
	MIN_INT := -MAX_INT - 1
	n := len(s)

	// Trim Left
	isNegative := false
	left := 0
	for left = 0; left < n; left++ {
		x := s[left]
		if x == ' ' {
			continue
		}
		// isWord
		if ('a' <= x && x <= 'z') || ('A' <= x && x <= 'Z') || x == '.' {
			return 0
		}
		if x == '-' || x == '+' {
			if left < n-1 && !isNumber(s[left+1]) {
				return 0
			}
			if x == '-' {
				isNegative = true
			}
			left++
		}
		if isNumber(x) {
			break
		}
		break
	}

	// Trim zeros
	for left < n {
		x := s[left]
		isZero := x == '0'
		if isZero {
			left++
		} else if !isNumber(x) {
			return 0
		} else {
			break
		}
	}

	var result int
	var pow int
	reset := func() {
		result = 0
		if isNegative {
			pow = -1
		} else {
			pow = 1
		}
	}
	reset()
	for i := n - 1; i >= left; i-- {
		x := s[i]
		if x == ' ' || ('a' <= x && x <= 'z') || ('A' <= x && x <= 'Z') || x == '.' || ((x == '-' || x == '+') && i > left) {
			reset()
			continue
		} else if isNumber(x) {
			result += pow * int(x-'0')
			if isNegative && (pow < int(MIN_INT) || result < int(MIN_INT)) {
				return int(MIN_INT)
			}
			if !isNegative && (pow > int(MAX_INT) || result > int(MAX_INT)) {
				return int(MAX_INT)
			}
			pow *= 10
		}
	}
	return result
}

func isNumber(x byte) bool {
	return '0' <= x && x <= '9'
}
