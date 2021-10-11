package task0014

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
//
// Constraints:
//   1 <= strs.length <= 200
//   0 <= strs[i].length <= 200
//   strs[i] consists of only lower-case English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
//   Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Longest Common Prefix.
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	result := strs[0]
	maxLength := len(result)
	for i := 1; i < n; i++ {
		currStr := strs[i]
		currLen := len(currStr)
		if currLen < maxLength {
			maxLength = currLen
		}
		for j := 0; j < len(currStr) && j < maxLength; j++ {
			if currStr[j] != result[j] {
				if j == 0 {
					return ""
				}
				maxLength = j
			}
		}
	}
	return result[:maxLength]
}
