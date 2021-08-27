package task0522

// Given an array of strings strs, return the length of the longest uncommon
// subsequence between them. If the longest uncommon subsequence does not exist, return -1.
// An uncommon subsequence between an array of strings is a string that is a subsequence of one string but not the others.
// A subsequence of a string s is a string that can be obtained after deleting any number of characters from s.
// For example, "abc" is a subsequence of "aebdc" because you can delete the underlined
// characters in "aebdc" to get "abc". Other subsequences of "aebdc" include "aebdc", "aeb", and "" (empty string).
//
// Constraints:
//   1 <= strs.length <= 50
//   1 <= strs[i].length <= 10
//   strs[i] consists of lowercase English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Uncommon Subsequence II.
//   Memory Usage: 2.3 MB, less than 14.29% of Go online submissions for Longest Uncommon Subsequence II.
func findLUSlength(strs []string) int {
	result, n := -1, len(strs)

	// Run through the strings
	for i := 0; i < n; i++ {
		// If current string is less than uncommon result - don't need to go further
		if len(strs[i]) < result {
			continue
		}

		// Compare string with others
		j := 0
		for j < n {
			if j != i && isSubSequence(strs[j], strs[i]) {
				break
			}
			j++
		}
		// if there is no string that is substring of current string - current string is good
		if j == n && result < len(strs[i]) {
			result = len(strs[i])
		}
	}
	return result
}

func isSubSequence(s string, sub string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i] == sub[j] {
			j++
		}
		i++
	}
	return j == len(sub)
}
