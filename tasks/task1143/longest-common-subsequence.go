package task1143

// Given two strings text1 and text2, return the length of their longest common subsequence.
// If there is no common subsequence, return 0.
// A subsequence of a string is a new string generated from the original string with some
// characters (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.
//
// Constraints:
//   1 <= text1.length, text2.length <= 1000
//   text1 and text2 consist of only lowercase English characters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Subsequence.
//   Memory Usage: 11.1 MB, less than 67.23% of Go online submissions for Longest Common Subsequence.
func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	// Complexity is O(n1*n2)

	// dp will contain the maximum result for i, j (i from text1, j from text2)
	dp := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i1 := 1; i1 <= n1; i1++ {
		for i2 := 1; i2 <= n2; i2++ {
			if text1[i1-1] == text2[i2-1] {
				// if chars are equal: we take the result for previous chars in both strings
				// and increment it by 1
				// e.g. "....xb...", "....yb...": current char is 'b': we take the result for "....x" and "....y" + 1
				dp[i1][i2] = dp[i1-1][i2-1] + 1
			} else {
				// if chars aren't equal: we take the Maximum
				// e.g. "....ab...", "....xy...": current chars are 'b' and 'y': we take the maximum result for 2 options:
				//    max ("....ab", "....x") and ("....a", "....xy")
				dp[i1][i2] = max(dp[i1][i2-1], dp[i1-1][i2])
			}
		}
	}
	return dp[n1][n2]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
