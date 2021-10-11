package task0096

// Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly
// n nodes of unique values from 1 to n
//
// Constraints:
//   1 <= n <= 19
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Binary Search Trees.
//   Memory Usage: 1.9 MB, less than 78.53% of Go online submissions for Unique Binary Search Trees.
func numTrees(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	cache := map[int]int{}
	return traverse(&cache, n)
}

// Recursive check for left and right number of options for each node int interval [0, n-1]
func traverse(cache *map[int]int, n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	if cacheResult, ok := (*cache)[n]; ok {
		return cacheResult
	}
	result := 0
	for i := 0; i < n; i++ {
		leftResult := 1
		rightResult := 1
		if i > 0 {
			leftResult = traverse(cache, i)
		}
		if i < n-1 {
			rightResult = traverse(cache, n-1-i)
		}
		result += leftResult * rightResult
	}
	(*cache)[n] = result
	return result
}
