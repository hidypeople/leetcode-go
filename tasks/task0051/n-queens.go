package task0051

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both
// indicate a queen and an empty space, respectively.
//
// Constraints:
//   1 <= n <= 9
//
// Results:
//   Runtime: 3 ms, faster than 87.50% of Go online submissions for N-Queens.
//   Memory Usage: 6.6 MB, less than 13.33% of Go online submissions for N-Queens.
func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	return solveRecursive(board, []int{}, 0, n)
}

// Recursive algorithm
func solveRecursive(board [][]byte, occupiedColumns []int, currentRowIndex, n int) [][]string {
	result := [][]string{}

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
			result = append(result, stringBoard)
			col++
			continue
		}

		currBoard := copyBoard(board)
		setQueen(currBoard, currentRowIndex, col, n)
		currResult := solveRecursive(currBoard, append(occupiedColumns, col), currentRowIndex+1, n)
		result = append(result, currResult...)
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
