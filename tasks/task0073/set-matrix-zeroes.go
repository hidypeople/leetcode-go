package task0073

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's, and return the matrix.
// You must do it in place.
//
// Constraints:
//   m == matrix.length
//   n == matrix[0].length
//   1 <= m, n <= 200
//   -2^31 <= matrix[i][j] <= 2^31 - 1
//
// Results:
//   Runtime: 8 ms, faster than 98.35% of Go online submissions for Set Matrix Zeroes. *Unstable*
//   Memory Usage: 6.4 MB, less than 80.40% of Go online submissions for Set Matrix Zeroes.
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	isZeroRowEmpty, isZeroColEmpty := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			isZeroColEmpty = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			isZeroRowEmpty = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// We use first row and column as indicator that we need to make column/row zero:
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Make zero according to first column values
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// Make zero according to first row values
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if isZeroRowEmpty {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if isZeroColEmpty {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
