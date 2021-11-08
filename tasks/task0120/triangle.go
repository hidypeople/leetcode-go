package task0120

// Given a triangle array, return the minimum path sum from top to bottom.
// For each step, you may move to an adjacent number of the row below. More formally,
// if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
//
// Constraints:
//   1 <= triangle.length <= 200
//   triangle[0].length == 1
//   triangle[i].length == triangle[i - 1].length + 1
//   -10^4 <= triangle[i][j] <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Triangle.
//   Memory Usage: 3.3 MB, less than 100.00% of Go online submissions for Triangle.
func minimumTotal(triangle [][]int) int {
	// DP solution:
	// We iterate triangle level by level starting from top
	// and on each level we relace:
	// triangle[i, j] = triangle[i, j] + min(triangle[i-1, j], triangle[i-1, j-1])
	// so the triangle[i, j] will keep the minimum path sum to i,j position

	// Go from the top to the bottom of triangle
	n := len(triangle)
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			minVal := 1<<63 - 1
			if j <= i-1 {
				minVal = triangle[i-1][j]
			}
			if j > 0 {
				minVal = min(minVal, triangle[i-1][j-1])
			}
			triangle[i][j] += minVal
		}
	}

	// Find the minimum value within the last triangle row
	result := 1<<63 - 1
	for i := 0; i < n; i++ {
		if triangle[n-1][i] < result {
			result = triangle[n-1][i]
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
