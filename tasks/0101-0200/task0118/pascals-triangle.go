package task0118

// Given an integer numRows, return the first numRows of Pascal's triangle.
//
// Constraints:
//   1 <= numRows <= 30
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
//   Memory Usage: 2.2 MB, less than 8.85% of Go online submissions for Pascal's Triangle.
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}

	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 0; j < len(result[i-1])-1; j++ {
			result[i][j+1] = result[i-1][j] + result[i-1][j+1]
		}
	}
	return result
}
