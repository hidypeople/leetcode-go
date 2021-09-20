package task1275

// Tic-tac-toe is played by two players A and B on a 3 x 3 grid.
// Here are the rules of Tic-Tac-Toe:
// Players take turns placing characters into empty squares (" ").
// The first player A always places "X" characters, while the second player B always places "O" characters.
// "X" and "O" characters are always placed into empty squares, never on filled ones.
// The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
// The game also ends if all squares are non-empty.
// No more moves can be played if the game is over.
// Given an array moves where each element is another array of size 2 corresponding to the row and column
// of the grid where they mark their respective character in the order in which A and B play.
// Return the winner of the game if it exists (A or B), in case the game ends in a draw return "Draw", if
// there are still movements to play return "Pending".
// You can assume that moves is valid (It follows the rules of Tic-Tac-Toe), the grid is initially empty
// and A will play first.
//
// Constraints:
//   1 <= moves.length <= 9
//   moves[i].length == 2
//   0 <= moves[i][j] <= 2
//   There are no repeated elements on moves.
//   moves follow the rules of tic tac toe.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Winner on a Tic Tac Toe Game.
//   Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Find Winner on a Tic Tac Toe Game.
func tictactoe(moves [][]int) string {
	n := len(moves)
	if n < 5 {
		return "Pending"
	}

	// "Play" the game

	var matrix [3][3]byte
	isPlayerA := true
	for _, move := range moves {
		if isPlayerA {
			matrix[move[0]][move[1]] = 'A'
		} else {
			matrix[move[0]][move[1]] = 'B'
		}
		isPlayerA = !isPlayerA
	}

	// Check game results

	// I don't want to check rows, columns and diagonals within their own cycles, so I made this
	// array. It contains records with 4 items each: {startR, startC, deltaR, deltaC}
	// It's a bit slower, but no boring code repetitions.
	directions := [][]int{
		{0, 0, 1, 1}, // diagonals
		{2, 0, -1, 1},
		{0, 0, 0, 1}, // rows
		{1, 0, 0, 1},
		{2, 0, 0, 1},
		{0, 0, 1, 0}, // columns
		{0, 1, 1, 0},
		{0, 2, 1, 0},
	}
	for _, direction := range directions {
		r, c := direction[0], direction[1]
		first := matrix[r][c]
		if first == 0 {
			continue
		}
		isWin := true
		for i := 0; i < 2; i++ {
			r += direction[2]
			c += direction[3]
			next := matrix[r][c]
			if next != first {
				isWin = false
				break
			}
		}
		if isWin {
			return string(first)
		}
	}

	if n < 9 {
		return "Pending"
	} else {
		return "Draw"
	}
}
