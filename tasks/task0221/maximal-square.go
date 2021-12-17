package task0221

// Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
//
// Constraints:
//   m == matrix.length
//   n == matrix[i].length
//   1 <= m, n <= 300
//   matrix[i][j] is '0' or '1'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximal Square.
//   Memory Usage: 4.2 MB, less than 68.59% of Go online submissions for Maximal Square.
func maximalSquare(matrix [][]byte) int {
	// DP solution
	m, n := len(matrix), len(matrix[0])
	prevLine := make([]int, n)
	result := 0
	for r := 0; r < m; r++ {
		// currLine and prevLine contains the maximum square for cells matrix[r][i] and matrix[r-1][i]
		currLine := make([]int, n)
		for c := 0; c < n; c++ {
			if matrix[r][c] == '0' {
				continue
			}
			if r == 0 || c == 0 {
				// Set max square on the first line to 1
				currLine[c] = 1
			} else {
				// Max Square for cell [r,c] is minumum of left,top,left-top + 1
				currLine[c] = min(currLine[c-1], min(prevLine[c], prevLine[c-1])) + 1
			}
			if result < currLine[c] {
				result = currLine[c]
			}
		}
		prevLine = currLine
	}
	return result * result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
