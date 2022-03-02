package task0392

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by
// deleting some (can be none) of the characters without disturbing the relative positions
// of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
// Constraints:
//   0 <= s.length <= 100
//   0 <= t.length <= 10^4
//   s and t consist only of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Is Subsequence.
//   Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Is Subsequence.
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
