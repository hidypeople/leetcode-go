package task0994

// You are given an m x n grid where each cell can have one of three values:
//   0 representing an empty cell,
//   1 representing a fresh orange, or
//   2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
//
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 10
//   grid[i][j] is 0, 1, or 2.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotting Oranges.
//   Memory Usage: 3 MB, less than 50.90% of Go online submissions for Rotting Oranges.
func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	// Calculate the number of fresh oranges and put rotten oranges to the stack
	numberFresh := 0
	stack := [][]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 2 {
				stack = append(stack, []int{i, j})
			} else if cell == 1 {
				numberFresh++
			}
		}
	}
	// If there is no fresh oranges - return 0
	if numberFresh == 0 {
		return 0
	}

	minutes := 0
	for len(stack) > 0 {
		nextStack := [][]int{}
		// Here we process all stack items within one batch (it is 1 minute)
		for len(stack) > 0 {
			rottenCoord := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			r, c := rottenCoord[0], rottenCoord[1]
			if r > 0 && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				nextStack = append(nextStack, []int{r - 1, c})
				numberFresh--
			}
			if r < n-1 && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				nextStack = append(nextStack, []int{r + 1, c})
				numberFresh--
			}
			if c > 0 && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				nextStack = append(nextStack, []int{r, c - 1})
				numberFresh--
			}
			if c < m-1 && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				nextStack = append(nextStack, []int{r, c + 1})
				numberFresh--
			}
		}
		minutes++
		if numberFresh <= 0 {
			break
		}
		stack = nextStack
	}
	if numberFresh > 0 {
		return -1
	}
	return minutes
}
