package task0132

// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return the minimum cuts needed for a palindrome partitioning of s.
//
// Constraints:
//   1 <= s.length <= 2000
//   s consists of lower-case English letters only.
//
// Results:
//   Runtime: 4 ms, faster than 80.00% of Go online submissions for Palindrome Partitioning II.
//   Memory Usage: 7.6 MB, less than 10.00% of Go online submissions for Palindrome Partitioning II.
func minCut(s string) int {
	n := len(s)

	// Minimum number of cuts required for substring
	cuts := make([]int, n)

	// DP will contain the "true" for [i,l] if the s[i:i+l+1] is palindrome
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		minCuts := length - 1
		for i := 0; i <= length-1; i++ {
			if s[i] == s[length-1] && (i+1 > length-2 || dp[i+1][length-2]) {
				// if the edges are the same and inner part also is palindrom => it is palindrom
				dp[i][length-1] = true
				if i == 0 {
					// The whole substring is palindrom
					minCuts = 0
				} else {
					// take best result for previous step
					newCuts := cuts[i-1] + 1
					if minCuts > newCuts {
						minCuts = newCuts
					}
				}
			}
		}
		cuts[length-1] = minCuts
	}

	return cuts[n-1]
}
