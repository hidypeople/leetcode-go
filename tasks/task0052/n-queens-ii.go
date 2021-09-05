package task0052

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.
//
// Constraints:
//   1 <= n <= 9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for N-Queens II.
//   Memory Usage: 6.3 MB, less than 5.17% of Go online submissions for N-Queens II.
func totalNQueens(n int) int {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	return solveRecursive(board, []int{}, 0, n)
}

// Recursive algorithm
func solveRecursive(board [][]byte, occupiedColumns []int, currentRowIndex, n int) int {
	result := 0

	col := 0
	for col < n {
		// If cell is occupied - skip
		if indexOf(occupiedColumns, col) >= 0 || board[currentRowIndex][col] != 0 {
			col++
			continue
		}

		// If we came to the end: set up the last queen and return the result
		if currentRowIndex == n-1 {
			board[currentRowIndex][col] = 'Q'
			// If we're on the last step: copy result into board
			stringBoard := make([]string, len(board))
			for i, b := range board {
				stringBoard[i] = string(b)
			}
			result++
			col++
			continue
		}

		currBoard := copyBoard(board)
		setQueen(currBoard, currentRowIndex, col, n)
		result += solveRecursive(currBoard, append(occupiedColumns, col), currentRowIndex+1, n)
		col++
	}
	return result
}

func setQueen(occupied [][]byte, row, col, n int) {
	occupied[row][col] = 'Q'
	for i := 0; i < n; i++ {
		if i != col {
			occupied[row][i] = '.'
		}
		if i > row {
			occupied[i][col] = '.'
		}
		if i > 0 && row+i < n && col+i < n {
			occupied[row+i][col+i] = '.'
		}
		if i > 0 && row+i < n && col-i >= 0 {
			occupied[row+i][col-i] = '.'
		}
	}
}

func indexOf(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func copyBoard(src [][]byte) [][]byte {
	n := len(src)
	dst := make([][]byte, n)
	for i, row := range src {
		dst[i] = make([]byte, n)
		copy(dst[i], row)
	}
	return dst
}
