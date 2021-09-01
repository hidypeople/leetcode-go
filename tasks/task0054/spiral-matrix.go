package task0054

// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Constraints:
//   m == matrix.length
//   n == matrix[i].length
//   1 <= m, n <= 10
//   -100 <= matrix[i][j] <= 100
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix.
//   Memory Usage: 2 MB, less than 100.00% of Go online submissions for Spiral Matrix.
func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])

	result := make([]int, n*m)

	// indexes of borders
	top, left, bottom, right := 0, 0, n-1, m-1
	direction := 0
	resultIdx := 0
	for left <= right && top <= bottom {
		switch direction {
		case 0:
			for i := left; i <= right; i++ {
				result[resultIdx] = matrix[top][i]
				resultIdx++
			}
			top++
		case 1:
			for i := top; i <= bottom; i++ {
				result[resultIdx] = matrix[i][right]
				resultIdx++
			}
			right--
		case 2:
			for i := right; i >= left; i-- {
				result[resultIdx] = matrix[bottom][i]
				resultIdx++
			}
			bottom--
		case 3:
			for i := bottom; i >= top; i-- {
				result[resultIdx] = matrix[i][left]
				resultIdx++
			}
			left++
		}
		direction = (direction + 1) % 4
	}

	return result
}
