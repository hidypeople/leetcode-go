package task0542

// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.
//
// Constraints:
//   m == mat.length
//   n == mat[i].length
//   1 <= m, n <= 10^4
//   1 <= m * n <= 10^4
//   mat[i][j] is either 0 or 1.
//   There is at least one 0 in mat.
//
// Results:
//   Runtime: 48 ms, faster than 99.57% of Go online submissions for 01 Matrix.
//   Memory Usage: 7.4 MB, less than 78.21% of Go online submissions for 01 Matrix.
func updateMatrix(mat [][]int) [][]int {
	INF := 1 << 32
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}

	// Go from top-left -> bottom-right
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := mat[i][j]
			if val == 0 {
				dp[i][j] = 0
			}
			if i > 0 {
				dp[i-1][j] = min(dp[i-1][j], dp[i][j]+1)
			}
			if i < m-1 {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			}
			if j > 0 {
				dp[i][j-1] = min(dp[i][j-1], dp[i][j]+1)
			}
			if j < n-1 {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+1)
			}
		}
	}
	// Go from bottom-right -> top-left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i > 0 {
				dp[i-1][j] = min(dp[i-1][j], dp[i][j]+1)
			}
			if i < m-1 {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			}
			if j > 0 {
				dp[i][j-1] = min(dp[i][j-1], dp[i][j]+1)
			}
			if j < n-1 {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+1)
			}
		}
	}
	return dp
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
