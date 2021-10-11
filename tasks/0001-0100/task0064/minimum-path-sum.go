package task0064

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right,
// which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.
//
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 200
//   0 <= grid[i][j] <= 100
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Minimum Path Sum.
//   Memory Usage: 3.9 MB, less than 75.38% of Go online submissions for Minimum Path Sum.
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				g1, g2 := grid[i-1][j], grid[i][j-1]
				if g1 < g2 {
					grid[i][j] += g1
				} else {
					grid[i][j] += g2
				}
				continue
			}
			if i == 0 && j > 0 {
				grid[0][j] += grid[0][j-1]
				continue
			}
			if i > 0 && j == 0 {
				grid[i][0] += grid[i-1][0]
			}
		}
	}
	return grid[m-1][n-1]
}
