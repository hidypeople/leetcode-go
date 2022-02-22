package task0151

// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters.
// The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
// Note that s may contain leading or trailing spaces or multiple spaces between
// two words. The returned string should only have a single space separating the
// words. Do not include any extra spaces.
//
// Constraints:
// 1 <= s.length <= 10^4
// s contains English letters (upper-case and lower-case), digits, and spaces ' '.
// There is at least one word in s.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Words in a String.
//   Memory Usage: 7.9 MB, less than 14.47% of Go online submissions for Reverse Words in a String.
func reverseWords(s string) string {
	words := []string{}
	wordStartIdx := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if wordStartIdx >= 0 {
				words = append(words, s[wordStartIdx:i])
				wordStartIdx = -1
			}
		} else {
			if wordStartIdx < 0 {
				wordStartIdx = i
			}
		}
	}
	if wordStartIdx >= 0 {
		words = append(words, s[wordStartIdx:])
	}
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i] + " "
	}
	return result[:len(result)-1]
}
