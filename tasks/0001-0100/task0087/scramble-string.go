package task0087

// We can scramble a string s to get a string t using the following algorithm:
// If the length of the string is 1, stop.
// If the length of the string is > 1, do the following:
// Split the string into two non-empty substrings at a random index, i.e., if the string is s, divide it to x and y where s = x + y.
// Randomly decide to swap the two substrings or to keep them in the same order. i.e., after this step, s may become s = x + y or s = y + x.
// Apply step 1 recursively on each of the two substrings x and y.
// Given two strings s1 and s2 of the same length, return true if s2 is a scrambled string of s1, otherwise, return false.
//
// Constraints:
//   s1.length == s2.length
//   1 <= s1.length <= 30
//   s1 and s2 consist of lower-case English letters.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Scramble String.
//   Memory Usage: 2.4 MB, less than 93.55% of Go online submissions for Scramble String.
func isScramble(s1input string, s2input string) bool {
	// Recursive solution.
	n := len(s1input)
	s1, s2 := &s1input, &s2input

	// dp meaning:
	//   1. it is cache of previously computed results
	//   2. it contains values 0, 1 or 2 (0 means unset, 1 means false, 2 means true)
	//   3. p[i][j][k] means the [i,j] interval k means length of the interval
	var dp [31][31][31]int
	return helper(&dp, s1, s2, 0, 0, n)
}

func helper(dp *[31][31][31]int, s1, s2 *string, i1, i2, n int) bool {
	if dp[i1][i2][n] != 0 {
		return dp[i1][i2][n] == 2
	}
	match := true
	for i := 0; i < n; i++ {
		if (*s1)[i1+i] != (*s2)[i2+i] {
			match = false
		}
	}
	if match {
		dp[i1][i2][n] = 2
		return true
	}
	count := make([]int, 26)
	for i := 0; i < n; i++ {
		count[(*s1)[i1+i]-'a']++
		count[(*s2)[i2+i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			dp[i1][i2][n] = 1
			return false
		}
	}
	for i := 1; i < n; i++ {
		if helper(dp, s1, s2, i1, i2, i) && helper(dp, s1, s2, i1+i, i2+i, n-i) {
			dp[i1][i2][n] = 2
			return true
		}
		if helper(dp, s1, s2, i1, i2+n-i, i) && helper(dp, s1, s2, i1+i, i2, n-i) {
			dp[i1][i2][n] = 2
			return true
		}
	}
	dp[i1][i2][n] = 1
	return false
}
