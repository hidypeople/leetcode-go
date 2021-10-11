package task0074

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.
//
// Constraints:
//   m == matrix.length
//   n == matrix[i].length
//   1 <= m, n <= 100
//   -10^4 <= matrix[i][j], target <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Search a 2D Matrix.
//   Memory Usage: 2.7 MB, less than 41.69% of Go online submissions for Search a 2D Matrix.
func searchMatrix(matrix [][]int, target int) bool {
	// Binary search. Compexity is O(ln(m*n))
	m, n := len(matrix), len(matrix[0])
	// Range check
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	// We found a row. Now we do binary search again and try to find the result within the row
	var middle, middleValue int
	left, right := 0, m*n-1
	for left < right-1 {
		middle = left + (right-left)/2
		middleValue = matrix[middle/n][middle%n]
		if target == middleValue {
			return true
		} else if target < middleValue {
			right = middle - 1
		} else {
			left = middle
		}
	}
	return target == matrix[left/n][left%n] || target == matrix[right/n][right%n]
}

// The same as previous, but here we find the row first, and then find column (it is a bit faster)
func searchMatrix2(matrix [][]int, target int) bool {
	// Binary search. Compexity is O(ln(m*n))

	m, n := len(matrix), len(matrix[0])

	// Range check
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	// Find row using binary search
	var middleRow, middleRowVal int
	topRow, bottomRow := 0, m-1
	for topRow < bottomRow-1 {
		middleRow = topRow + (bottomRow-topRow)/2
		middleRowVal = matrix[middleRow][0]
		if target < middleRowVal {
			bottomRow = middleRow - 1
		} else if target > middleRowVal {
			topRow = middleRow
		} else {
			return true
		}
	}
	var row []int
	if target < matrix[bottomRow][0] {
		row = matrix[topRow]
	} else {
		row = matrix[bottomRow]
	}

	// We found a row. Now we do binary search again and try to find the result within the row
	var middle, middleValue int
	left, right := 0, n-1
	for left < right-1 {
		middle = left + (right-left)/2
		middleValue = row[middle]
		if target < middleValue {
			right = middle - 1
		} else if target > middleValue {
			left = middle
		} else {
			return true
		}
	}
	return target == row[left] || target == row[right]
}
