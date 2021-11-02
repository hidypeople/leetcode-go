package task0980

// You are given an m x n integer array grid where grid[i][j] could be:
//   1 representing the starting square. There is exactly one starting square.
//   2 representing the ending square. There is exactly one ending square.
//   0 representing empty squares we can walk over.
//   -1 representing obstacles that we cannot walk over.
// Return the number of 4-directional walks from the starting square to the ending square,
// that walk over every non-obstacle square exactly once.
//
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 20
//   1 <= m * n <= 20
//   -1 <= grid[i][j] <= 2
//   There is exactly one starting cell and one ending cell.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths III.
//   Memory Usage: 2.4 MB, less than 24.39% of Go online submissions for Unique Paths III.

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	numberOfSteps := 0
	start := []int{}
	// store the visited cells in the separate array, we can use the grid with specical "visited"
	// value, it will keep some memory, but I decided don't do that
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			cell := grid[i][j]
			if cell != 0 && cell != 2 {
				visited[i][j] = true
			}
			if cell == 0 {
				numberOfSteps++
			} else if cell == 1 {
				start = []int{i, j}
			}
		}
	}

	return dfs(&grid, &visited, m, n, numberOfSteps, start)
}

func dfs(grid *[][]int, visited *[][]bool, m, n, numberOfSteps int, current []int) int {
	// Use dfs to calc all routes from current point
	r, c := current[0], current[1]
	result := 0
	(*visited)[r][c] = true
	for _, direction := range directions {
		nextR, nextC := r+direction[0], c+direction[1]
		if nextR >= 0 && nextR < m && nextC >= 0 && nextC < n && !(*visited)[nextR][nextC] {
			nextVal := (*grid)[nextR][nextC]
			if nextVal == 2 {
				if numberOfSteps == 0 {
					result++
				}
			} else if nextVal == 0 {
				result += dfs(grid, visited, m, n, numberOfSteps-1, []int{nextR, nextC})
			}
		}
	}
	(*visited)[r][c] = false
	return result
}
