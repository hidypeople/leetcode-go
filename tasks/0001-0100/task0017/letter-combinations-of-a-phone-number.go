package task0017

// Given a string containing digits from 2-9 inclusive, return all possible
// letter combinations that the number could represent. Return the answer in any order.
// A mapping of digit to letters (just like on the telephone buttons) is given
// below. Note that 1 does not map to any letters.
//
// Constraints:
//   0 <= digits.length <= 4
//   digits[i] is a digit in the range ['2', '9'].
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
//   Memory Usage: 2.2 MB, less than 23.93% of Go online submissions for Letter Combinations of a Phone Number.
func letterCombinations(digits string) []string {
	return combine(digits, make([]string, 0))
}

var letters = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func combine(digits string, current []string) []string {
	if len(digits) == 0 {
		return current
	}
	lenCurrent := len(current)
	digit := digits[0]
	digits = digits[1:]
	if lenCurrent == 0 {
		return combine(digits, letters[digit])
	}
	chars := letters[digit]
	lenChars := len(chars)
	new_n := lenCurrent * lenChars

	newCurrent := make([]string, new_n)
	for i := 0; i < lenCurrent; i++ {
		currItem := current[i]
		for j := 0; j < lenChars; j++ {
			newCurrent[i*lenChars+j] = currItem + chars[j]
		}
	}
	return combine(digits, newCurrent)
}
