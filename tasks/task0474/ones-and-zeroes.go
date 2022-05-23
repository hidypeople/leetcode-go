package task0474

// You are given an array of binary strings strs and two integers m and n.
// Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
// A set x is a subset of a set y if all elements of x are also elements of y.
//
// Constraints:
//   1 <= strs.length <= 600
//   1 <= strs[i].length <= 100
//   strs[i] consists only of digits '0' and '1'.
//   1 <= m, n <= 100
//
// Results:
//   Runtime: 25 ms, faster than 97.56% of Go online submissions for Ones and Zeroes.
//   Memory Usage: 3.6 MB, less than 100.00% of Go online submissions for Ones and Zeroes.
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	var i int
	var j int
	var cnt0 int
	var cnt1 int
	var chr rune
	var val int
	var newVal int
	for i = 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		cnt0, cnt1 = 0, 0
		for _, chr = range str {
			if chr == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}

		for i = m; i >= cnt0; i-- {
			for j = n; j >= cnt1; j-- {
				val = dp[i][j]
				newVal = dp[i-cnt0][j-cnt1] + 1
				if val < newVal {
					dp[i][j] = newVal
				}
			}
		}
	}
	return dp[m][n]
}
