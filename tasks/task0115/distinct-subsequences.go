package task0115

// Given two strings s and t, return the number of distinct subsequences of s which equals t.
// A string's subsequence is a new string formed from the original string by deleting some
// (can be none) of the characters without disturbing the remaining characters' relative
// positions. (i.e., "ACE" is a subsequence of "ABCDE" while "AEC" is not).
// It is guaranteed the answer fits on a 32-bit signed integer.
//
// Constraints:
//   1 <= s.length, t.length <= 1000
//   s and t consist of English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Distinct Subsequences.
//   Memory Usage: 2 MB, less than 93.81% of Go online submissions for Distinct Subsequences.
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)

	// DP[i] contains the number of substrings for slice t[:i+1]
	dp := make([]int, m)

	// s = "rabbbit", t = "rabbit"
	for j := 0; j < n; j++ {
		// run from first to the last source char: r, a, b, b, b, i, t
		for i := m - 1; i >= 0; i-- {
			// recalculate all DP items for current s[j] char
			//   | r, a, b, b, i, t
			// ---------------------
			// r | 1, 0, 0, 0, 0, 0 // dp[0] += 1
			// a | 1, 1, 0, 0, 0, 0 // dp[1] += dp[3]
			// b | 1, 1, 1, 0, 0, 0 // dp[2] += dp[1]
			// b | 1, 1, 2, 1, 0, 0 // dp[3] += dp[2], dp[2] += dp[1]
			// b | 1, 1, 3, 3, 0, 0 // dp[3] += dp[2], dp[2] += dp[1]
			// i | 1, 1, 3, 3, 3, 0 // dp[4] += dp[3]
			// t | 1, 1, 3, 3, 3, 3 // dp[5] += dp[4]
			//
			// result in dp[5]
			if t[i] == s[j] {
				prev := 1
				if i > 0 {
					prev = dp[i-1]
				}
				dp[i] += prev
			}
		}
	}
	return dp[len(dp)-1]
}
