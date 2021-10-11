package task0063

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot is trying to reach
// the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// Now consider if some obstacles are added to the grids. How many unique paths would there be?
// An obstacle and space is marked as 1 and 0 respectively in the grid.
//
// Constraints:
//   m == obstacleGrid.length
//   n == obstacleGrid[i].length
//   1 <= m, n <= 100
//   obstacleGrid[i][j] is 0 or 1.
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths II.
//   Memory Usage: 2.8 MB, less than 10.59% of Go online submissions for Unique Paths II.
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	countRows, countCols := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[countRows-1][countCols-1] == 1 {
		return 0
	}
	// This cache contains the results for grids MxN
	cache := make(map[int]map[int]int)
	return uniquePathsWithObstaclesRecursive(cache, obstacleGrid, 0, 0, countRows, countCols)
}

func uniquePathsWithObstaclesRecursive(cache map[int]map[int]int, obstacleGrid [][]int, r, c, countRows, countCols int) int {
	if r == countRows-1 {
		// last row: directly check for further obstacles
		for i := c + 1; i < countCols; i++ {
			if obstacleGrid[r][i] == 1 {
				return 0
			}
		}
		return 1
	}
	if c == countCols-1 {
		// last column: directly check for further obstacles
		for i := r + 1; i < countRows; i++ {
			if obstacleGrid[i][c] == 1 {
				return 0
			}
		}
		return 1
	}

	// Return from cache
	if cacheResult, ok := cacheGet(cache, r, c); ok {
		return cacheResult
	}

	// Calc result, add to cache and return
	result := 0
	if obstacleGrid[r+1][c] == 0 {
		result += uniquePathsWithObstaclesRecursive(cache, obstacleGrid, r+1, c, countRows, countCols)
	}
	if obstacleGrid[r][c+1] == 0 {
		result += uniquePathsWithObstaclesRecursive(cache, obstacleGrid, r, c+1, countRows, countCols)
	}
	cacheAdd(cache, r, c, result)
	return result
}

func cacheGet(cache map[int]map[int]int, r, c int) (int, bool) {
	if cacheC, ok := cache[r]; ok {
		if count, ok2 := cacheC[c]; ok2 {
			return count, true
		}
	}
	return 0, false
}

func cacheAdd(cache map[int]map[int]int, r, c, value int) {
	if _, ok := cache[r]; !ok {
		cache[r] = make(map[int]int)
	}
	cache[r][c] = value
}
