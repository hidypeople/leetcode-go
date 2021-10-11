package task0037

// Write a program to solve a Sudoku puzzle by filling the empty cells.
// A sudoku solution must satisfy all of the following rules:
//   Each of the digits 1-9 must occur exactly once in each row.
//   Each of the digits 1-9 must occur exactly once in each column.
//   Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
//   The '.' character indicates empty cells.
//
// Constraints:
//   board.length == 9
//   board[i].length == 9
//   board[i][j] is a digit or '.'.
// It is guaranteed that the input board has only one solution.
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Sudoku Solver.
//   Memory Usage: 2.4 MB, less than 34.65% of Go online submissions for Sudoku Solver.
func solveSudoku(board [][]byte) {
	dfs(board, 0)
}

func dfs(board [][]byte, idx int) bool {
	nextIdx := idx
	// skip occupied cells: we need to find free cell
	for nextIdx < 81 {
		x1 := nextIdx / 9
		y1 := nextIdx % 9
		if board[x1][y1] == '.' {
			break
		}
		nextIdx++
	}
	if nextIdx > 80 {
		return true
	}
	// Get available values for current cell
	x := nextIdx / 9
	y := nextIdx % 9
	availables := availableValues(board, x, y)
	if len(availables) == 0 {
		// No available value: break the chain (somewhere before we did wrong choise)
		return false
	}
	for _, availableValue := range availables {
		// Set up available value and go deeper
		board[x][y] = availableValue
		if dfs(board, nextIdx+1) {
			return true
		}
		// Restore if we're coming from incorrect chain
		board[x][y] = '.'
	}
	return false
}

// Get array of '1'-'9' that can be allied to current cell
func availableValues(board [][]byte, x, y int) []byte {
	var done int16 = 0
	for i := 0; i < 9; i++ {
		// column
		if board[i][y] != '.' {
			done = done | 1<<(board[i][y]-'0')
		}
		// cell
		if board[x][i] != '.' {
			done = done | 1<<(board[x][i]-'0')
		}
		// 3x3 square
		if board[3*(x/3)+i/3][3*(y/3)+i%3] != '.' {
			done = done | 1<<(board[3*(x/3)+i/3][3*(y/3)+i%3]-'0')
		}
	}
	result := []byte{}
	var b byte = '1'
	for b <= '9' {
		if (1<<(b-'0'))&done == 0 {
			result = append(result, b)
		}
		b++
	}
	return result
}
