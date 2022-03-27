package task1337

import "sort"

// You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians).
// The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of all
// the 0's in each row.
// A row i is weaker than a row j if one of the following is true:
//   The number of soldiers in row i is less than the number of soldiers in row j.
//   Both rows have the same number of soldiers and i < j.
// Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.
//
// Constraints:
//   m == mat.length
//   n == mat[i].length
//   2 <= n, m <= 100
//   1 <= k <= m
//   matrix[i][j] is either 0 or 1.
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for The K Weakest Rows in a Matrix.
//   Memory Usage: 4.8 MB, less than 97.37% of Go online submissions for The K Weakest Rows in a Matrix.
func kWeakestRows(mat [][]int, k int) []int {
	// Put into each matrix row additional values: [stength, index, 1 ... 1, 0 ... 0]
	lenC := len(mat[0])
	for r, row := range mat {
		strength := 0
		for strength < lenC && row[strength] == 1 {
			strength++
		}
		row[0] = strength
		row[1] = r
	}

	// Sort by: [ strength, index ]
	sort.Slice(mat, func(i, j int) bool {
		if mat[i][0] != mat[j][0] {
			return mat[i][0] <= mat[j][0]
		}
		return mat[i][1] <= mat[j][1]
	})

	// Create result
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = mat[i][1]
	}
	return result
}
