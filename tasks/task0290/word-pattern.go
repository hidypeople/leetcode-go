package task0290

// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
//
// Constraints:
//   1 <= pattern.length <= 300
//   pattern contains only lower-case English letters.
//   1 <= s.length <= 3000
//   s contains only lowercase English letters and spaces ' '.
//   s does not contain any leading or trailing spaces.
//   All the words in s are separated by a single space.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Pattern.
//   Memory Usage: 2.1 MB, less than 10.38% of Go online submissions for Word Pattern.
func wordPattern(pattern string, s string) bool {
	occupiedWords := map[string]bool{}
	occupiedPatterns := make([]string, 26)
	patternItem, patternIdx, isPatternExprired := pattern[0], 0, false
	startIdx := 0
	for i, chr := range s {
		if isPatternExprired {
			if patternIdx == len(pattern)-1 {
				return false
			}
			patternIdx++
			patternItem = pattern[patternIdx]
			isPatternExprired = false
		}
		if chr == ' ' || i+1 == len(s) {
			endIdx := i
			if i+1 == len(s) {
				endIdx = i + 1
				if patternIdx+1 < len(pattern) {
					return false
				}
			}
			word := s[startIdx:endIdx]
			expectedWord := occupiedPatterns[patternItem-'a']
			if len(expectedWord) == 0 {
				if occupiedWords[word] {
					return false
				}
				// new word
				occupiedPatterns[patternItem-'a'] = word
				occupiedWords[word] = true
			} else if expectedWord != word {
				return false
			}
			isPatternExprired = true
			startIdx = i + 1
		}
	}
	return true
}
