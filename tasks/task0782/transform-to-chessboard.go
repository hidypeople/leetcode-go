package task0782

// You are given an n x n binary grid board. In each move, you can swap any two rows with each other, or any two columns with each other.
// Return the minimum number of moves to transform the board into a chessboard board. If the task is impossible, return -1.
// A chessboard board is a board where no 0's and no 1's are 4-directionally adjacent.
//
// Constraints:
//   n == board.length
//   n == board[i].length
//   2 <= n <= 30
//   board[i][j] is either 0 or 1.
//
// Results:
//   Runtime: 8 ms, faster than 100.00% of Go online submissions for Transform to Chessboard.
//   Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Transform to Chessboard.
func movesToChessboard(board [][]int) int {
	n, rowSum, colSum, rowSwap, colSwap := len(board), 0, 0, 0, 0

	// Find invalid rectangles in the board
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (board[0][0] ^ board[i][0] ^ board[0][j] ^ board[i][j]) == 1 {
				return -1
			}
		}
	}

	// 1. Calculate the sum of values within first row and column
	// 2. Calculate the count of swaps for first row and column:
	//    e.g. current:  [0,1,1,0,1,1,0,0]
	//         expected: [0,1,0,1,0,1,0,1]
	//                   [+,+,-,-,-,+,+,-] -> swap count = 4
	for i := 0; i < n; i++ {
		rowSum += board[i][0]
		colSum += board[0][i]
		if board[i][0] == i%2 {
			rowSwap++
		}
		if board[0][i] == i%2 {
			colSwap++
		}
	}
	// The count of 0 and 1 should be the same or it can differ by 1
	if (rowSum != n/2 && rowSum != (n+1)/2) || (colSum != n/2 && colSum != (n+1)/2) {
		return -1
	}

	if n%2 == 1 {
		// The odd board size
		// if colSwap or rowSwap is odd => change direction
		if colSwap%2 == 1 {
			colSwap = n - colSwap
		}
		if rowSwap%2 == 1 {
			rowSwap = n - rowSwap
		}
	} else {
		// The even board size
		// We need to find the minimum number of swaps to make cells interchange
		colSwap = min(n-colSwap, colSwap)
		rowSwap = min(n-rowSwap, rowSwap)
	}

	// Divide by two: because each swap moves 2 columns or rows
	return (colSwap + rowSwap) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
