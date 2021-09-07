package task0059

// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
//
// Constraints:
//   1 <= n <= 20
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix II.
//   Memory Usage: 2.2 MB, less than 95.40% of Go online submissions for Spiral Matrix II.
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	top, left, bottom, right := 0, 0, n-1, n-1
	direction := 0
	counter := 1
	for top <= bottom && left <= right {
		switch direction {
		case 0:
			// top border
			for i := left; i <= right; i++ {
				result[top][i] = counter
				counter++
			}
			top++
		case 1:
			// right border
			for i := top; i <= bottom; i++ {
				result[i][right] = counter
				counter++
			}
			right--
		case 2:
			// bottom border
			for i := right; i >= left; i-- {
				result[bottom][i] = counter
				counter++
			}
			bottom--
		case 3:
			// left border
			for i := bottom; i >= top; i-- {
				result[i][left] = counter
				counter++
			}
			left++
		}
		direction = (direction + 1) % 4
	}

	return result
}
