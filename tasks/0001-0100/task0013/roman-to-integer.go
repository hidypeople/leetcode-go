package task0013

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
//
// Constraints:
//   1 <= s.length <= 15
//   s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
//   It is guaranteed that s is a valid roman numeral in the range [1, 3999].
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Roman to Integer.
//   Memory Usage: 3 MB, less than 40.50% of Go online submissions for Roman to Integer.
func romanToInt(s string) int {
	result := 0
	lookahead := func(s string, i int) byte {
		if i+1 >= len(s) {
			return 0
		}
		return s[i+1]
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			next := lookahead(s, i)
			if next == 'D' {
				result += 400
				i++
			} else if next == 'M' {
				result += 900
				i++
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			next := lookahead(s, i)
			if next == 'L' {
				result += 40
				i++
			} else if next == 'C' {
				result += 90
				i++
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			next := lookahead(s, i)
			if next == 'V' {
				result += 4
				i++
			} else if next == 'X' {
				result += 9
				i++
			} else {
				result += 1
			}
		}
	}
	return result
}
