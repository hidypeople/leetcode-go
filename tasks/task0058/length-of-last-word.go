package task0058

// Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.
//
// Constraints:
//   1 <= s.length <= 104
//   s consists of only English letters and spaces ' '.
//   There will be at least one word in s.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
//   Memory Usage: 2.2 MB, less than 85.21% of Go online submissions for Length of Last Word.
func lengthOfLastWord(s string) int {
	result := 0
	isSpace := true
	for i := len(s) - 1; i >= 0; i-- {
		chr := s[i]
		if chr != ' ' {
			isSpace = false
			result++
		} else if chr == ' ' && !isSpace {
			return result
		}
	}
	return result
}
