package task0062

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot is trying to reach
// the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?
//
// Constraints:
//   1 <= m, n <= 100
//   It's guaranteed that the answer will be less than or equal to 2 * 109.
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
//   Memory Usage: 2.5 MB, less than 17.68% of Go online submissions for Unique Paths.
func uniquePaths(m int, n int) int {

	// This cache contains the results for grids MxN
	cache := make(map[int]map[int]int)
	return uniquePathsRecursive(cache, m, n)
}

func uniquePathsRecursive(cache map[int]map[int]int, m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	var result int
	if cacheN, ok := cache[m]; ok {
		if count, ok2 := cacheN[n]; ok2 {
			// Hit cache - return result
			return count
		}
		// Calculate result and put it into the cache
		result = uniquePathsRecursive(cache, m-1, n) + uniquePathsRecursive(cache, m, n-1)
		cacheN[n] = result
		return result
	} else {
		// Calculate result and put it into the cache
		cache[m] = make(map[int]int)
		result = uniquePathsRecursive(cache, m-1, n) + uniquePathsRecursive(cache, m, n-1)
		cache[m][n] = result
		return result
	}
}
