package task1446

// The power of the string is the maximum length of a non-empty substring that
// contains only one unique character.
// Given a string s, return the power of s.
//
// Constraints:
//   1 <= s.length <= 500
//   s consists of only lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Consecutive Characters.
//   Memory Usage: 2.3 MB, less than 38.71% of Go online submissions for Consecutive Characters.
func maxPower(s string) int {
	maxLen := 1
	prevChr := s[0]
	currLen := 1
	for i := 1; i < len(s); i++ {
		if s[i] == prevChr {
			currLen++
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			currLen = 1
			prevChr = s[i]
		}
	}
	return maxLen
}
