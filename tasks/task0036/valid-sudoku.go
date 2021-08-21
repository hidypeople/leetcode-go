package task0036

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
//
// Constraints:
//   board.length == 9
//   board[i].length == 9
//   board[i][j] is a digit or '.'.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Sudoku.
//   Memory Usage: 3.2 MB, less than 45.77% of Go online submissions for Valid Sudoku.
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	// rows and cells
	for i := 0; i < n; i++ {
		buffer1 := make(map[byte]bool, 9)
		buffer2 := make(map[byte]bool, 9)
		for j := 0; j < n; j++ {
			val1 := board[i][j]
			if val1 != '.' {
				if _, ok := buffer1[val1]; ok {
					return false
				}
				buffer1[val1] = true
			}
			val2 := board[j][i]
			if val2 != '.' {
				if _, ok := buffer2[val2]; val2 != '.' && ok {
					return false
				}
				buffer2[val2] = true
			}
		}
	}
	// 3x3 squares
	for i := 0; i < n; i += 3 {
		for j := 0; j < n; j += 3 {
			buffer1 := make(map[byte]bool, 9)
			for k := 0; k < 9; k++ {
				val1 := board[i+k/3][j+k%3]
				if val1 != '.' {
					if _, ok := buffer1[val1]; val1 != '.' && ok {
						return false
					}
					buffer1[val1] = true
				}
			}
		}
	}
	return true
}
