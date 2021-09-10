package task0885

// You start at the cell (rStart, cStart) of an rows x cols grid facing east. The northwest
// corner is at the first row and column in the grid, and the southeast corner is at the last row and column.
// You will walk in a clockwise spiral shape to visit every position in this grid. Whenever
// you move outside the grid's boundary, we continue our walk outside the grid (but may return to the
// grid boundary later.). Eventually, we reach all rows * cols spaces of the grid.
// Return an array of coordinates representing the positions of the grid in the order you visited them.
//
// Constraints:
//   1 <= rows, cols <= 100
//   0 <= rStart < rows
//   0 <= cStart < cols
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Spiral Matrix III.
//   Memory Usage: 6.9 MB, less than 84.21% of Go online submissions for Spiral Matrix III.
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	result := make([][]int, rows*cols)
	result[0] = []int{rStart, cStart}
	iResult := 1

	spins := max(max(rStart, rows-rStart)+1, max(cStart, cols-cStart))

	direction := 0
	currentSteps := 0
	currRow, currCol := rStart, cStart
	for k := 0; k < spins*4; k++ {
		if iResult == len(result) {
			return result
		}
		if direction == 0 {
			// top border
			currentSteps++
			if currRow >= 0 && currRow < rows {
				for i := 1; i <= currentSteps && currCol+i < cols; i++ {
					if currCol+i >= 0 {
						result[iResult] = []int{currRow, currCol + i}
						iResult++
					}
				}
			}
			currCol += currentSteps
			direction = 1
			continue
		}
		if direction == 1 {
			// right border
			if currCol >= 0 && currCol < cols {
				for i := 1; i <= currentSteps && currRow+i < rows; i++ {
					if currRow+i >= 0 {
						result[iResult] = []int{currRow + i, currCol}
						iResult++
					}
				}
			}
			currRow += currentSteps
			direction = 2
			continue
		}
		if direction == 2 {
			// bottom border
			currentSteps++
			if currRow >= 0 && currRow < rows {
				for i := 1; i <= currentSteps && currCol-i >= 0; i++ {
					if currCol-i < cols {
						result[iResult] = []int{currRow, currCol - i}
						iResult++
					}
				}
			}
			currCol -= currentSteps
			direction = 3
			continue
		}

		// left border
		if currCol >= 0 && currCol < cols {
			for i := 1; i <= currentSteps && currRow-i >= 0; i++ {
				if currRow-i < rows {
					result[iResult] = []int{currRow - i, currCol}
					iResult++
				}
			}
		}
		currRow -= currentSteps
		direction = 0
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
