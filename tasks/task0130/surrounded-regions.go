package task0130

// Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
//
// Constraints:
//   m == board.length
//   n == board[i].length
//   1 <= m, n <= 200
//   board[i][j] is 'X' or 'O'.
//
// Results:
//   Runtime: 16 ms, faster than 99.49% of Go online submissions for Surrounded Regions.
//   Memory Usage: 10.4 MB, less than 32.99% of Go online submissions for Surrounded Regions.
func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	// If point has already been visited pointsVisited[i][j] will be true, false otherwise
	pointsVisited := make([][]bool, m)
	for i := 0; i < m; i++ {
		pointsVisited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' || pointsVisited[i][j] {
				// Skip 'X' and visited points
				continue
			}

			// found 'O' => run DFS and collect 'O'-points into the pointsDone
			// DFS is preferred because it's better chanse to reach the border and
			// stop collecting pointsDone
			pointsDone := [][2]int{}
			stack := [][2]int{{i, j}}
			isSurrounded := true
			for len(stack) > 0 {
				currentPoint := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				r, c := currentPoint[0], currentPoint[1]
				if isSurrounded {
					pointsDone = append(pointsDone, [2]int{r, c})
				}
				pointsVisited[r][c] = true
				if r == 0 || r == m-1 || c == 0 || c == n-1 {
					// If current point is on the border => the whole area is not surrounded
					isSurrounded = false
				}
				if r > 0 && board[r-1][c] == 'O' && !pointsVisited[r-1][c] {
					stack = append(stack, [2]int{r - 1, c})
				}
				if r < m-1 && board[r+1][c] == 'O' && !pointsVisited[r+1][c] {
					stack = append(stack, [2]int{r + 1, c})
				}
				if c > 0 && board[r][c-1] == 'O' && !pointsVisited[r][c-1] {
					stack = append(stack, [2]int{r, c - 1})
				}
				if c < n-1 && board[r][c+1] == 'O' && !pointsVisited[r][c+1] {
					stack = append(stack, [2]int{r, c + 1})
				}
			}
			if isSurrounded {
				// If current region is surrounded => change the points 'O' -> 'X'
				for _, pointDone := range pointsDone {
					board[pointDone[0]][pointDone[1]] = 'X'
				}
			}
		}
	}
}
