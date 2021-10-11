package task0065

// A valid number can be split up into these components (in order):
//   1. A decimal number or an integer.
//   2. (Optional) An 'e' or 'E', followed by an integer.
// A decimal number can be split up into these components (in order):
//   1. (Optional) A sign character (either '+' or '-').
//   2. One of the following formats:
//     1. One or more digits, followed by a dot '.'.
//     2. One or more digits, followed by a dot '.', followed by one or more digits.
//     3. A dot '.', followed by one or more digits.
// An integer can be split up into these components (in order):
//   1. (Optional) A sign character (either '+' or '-').
//   2. One or more digits.
// For example, all the following are valid numbers:
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"],
// while the following are not valid numbers:\
// ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
// Given a string s, return true if s is a valid number.
//
// Constraints:
//   1 <= s.length <= 20
//   s consists of only English letters (both uppercase and lowercase), digits (0-9), plus '+', minus '-', or dot '.'.
//
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Number.
//   Memory Usage: 2.5 MB, less than 21.05% of Go online submissions for Valid Number.
func isNumber(s string) bool {
	bytes := []byte(s)
	dotCount := 0
	digitCount := 0
	for i, c := range bytes {
		is_digit, is_sign, is_e, is_dot := isDigit(c), isSign(c), isE(c), c == '.'

		// Filter out illegal symbols
		if !is_sign && !is_digit && !is_e && !is_dot {
			return false
		}
		// Sign not on the first position => fail
		if is_sign && i > 0 {
			return false
		}
		if is_dot {
			// Not allowed more than 1 dot
			if dotCount > 0 {
				return false
			}
			dotCount++
		}
		if is_digit {
			digitCount++
		}

		// We met 'e' sign: validate the rest part is integer
		if is_e {
			// we need at least 1 digit before 'e'
			if digitCount == 0 {
				return false
			}
			return validateInteger(bytes[i+1:])
		}
	}
	return digitCount > 0
}

func validateInteger(bytes []byte) bool {
	if len(bytes) == 0 {
		return false
	}
	digitCount := 0
	for i, c := range bytes {
		digit, sign := isDigit(c), isSign(c)

		// Filter out illegal symbols
		if !digit && !sign {
			return false
		}

		// Sign can be only on 1st place
		if i > 0 && sign {
			return false
		}
		if digit {
			digitCount++
		}
	}
	// We need at least 1 digit
	return digitCount > 0
}

func isE(c byte) bool {
	return c == 'e' || c == 'E'
}

func isSign(c byte) bool {
	return c == '+' || c == '-'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
