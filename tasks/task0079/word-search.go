package task0079

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are
// horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
// Constraints:
//   m == board.length
//   n = board[i].length
//   1 <= m, n <= 6
//   1 <= word.length <= 15
//   board and word consists of only lowercase and uppercase English letters.
//
// Results:
//   Runtime: 44 ms, faster than 99.21% of Go online submissions for Word Search.
//   Memory Usage: 2 MB, less than 92.91% of Go online submissions for Word Search.
func exist(board [][]byte, word string) bool {
	lenRows, lenCols := len(board), len(board[0])
	for i := 0; i < lenRows; i++ {
		for j := 0; j < lenCols; j++ {
			if board[i][j] == word[0] && traverseDfs(board, word, lenRows, lenCols, i, j) {
				return true
			}
		}
	}
	return false
}

func traverseDfs(board [][]byte, word string, lenRows, lenCols, r, c int) bool {
	char := board[r][c]
	if len(word) == 1 {
		// We've just matched the last symbol
		return true
	}
	board[r][c] = 0 // wipe char
	newWord := word[1:]
	if (c > 0 && board[r][c-1] == newWord[0] && traverseDfs(board, newWord, lenRows, lenCols, r, c-1)) ||
		(r > 0 && board[r-1][c] == newWord[0] && traverseDfs(board, newWord, lenRows, lenCols, r-1, c)) ||
		(c < lenCols-1 && board[r][c+1] == newWord[0] && traverseDfs(board, newWord, lenRows, lenCols, r, c+1)) ||
		(r < lenRows-1 && board[r+1][c] == newWord[0] && traverseDfs(board, newWord, lenRows, lenCols, r+1, c)) {
		return true
	}
	board[r][c] = char // restore char
	return false
}
