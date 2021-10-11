package task0085

// Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
//
// Constraints:
//   rows == matrix.length
//   cols == matrix[i].length
//   0 <= row, cols <= 200
//   matrix[i][j] is '0' or '1'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximal Rectangle.
//   Memory Usage: 4.9 MB, less than 44.26% of Go online submissions for Maximal Rectangle.
func maximalRectangle(matrix [][]byte) int {
	// Complexity: O(N*M)
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	maxRectSquare := 0
	heights := make([]int, n)
	for r := 0; r < m; r++ {
		emptyCount := 0
		// Complexity: O(N)
		for c := 0; c < n; c++ {
			val := matrix[r][c]
			if val == '0' {
				emptyCount++
				heights[c] = 0
			} else {
				heights[c]++
			}
		}
		if emptyCount == n {
			continue
		}
		// task0084: max rectangle in histogram
		// Complexity: O(N)
		leftBoundaryArray, rightBoundaryArray := make([]int, n), make([]int, n)
		leftBoundaryArray[0] = -1
		rightBoundaryArray[n-1] = n
		for i := 1; i < n; i++ {
			p := i - 1
			for p >= 0 && heights[p] >= heights[i] {
				p = leftBoundaryArray[p]
			}
			leftBoundaryArray[i] = p
		}
		for i := n - 2; i >= 0; i-- {
			p := i + 1
			for p < n && heights[p] >= heights[i] {
				p = rightBoundaryArray[p]
			}
			rightBoundaryArray[i] = p
		}
		maxHeight := 0
		for i := 0; i < n; i++ {
			maxHeight = max(maxHeight, heights[i]*(rightBoundaryArray[i]-leftBoundaryArray[i]-1))
		}
		maxRectSquare = max(maxRectSquare, maxHeight)
	}
	return maxRectSquare
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
