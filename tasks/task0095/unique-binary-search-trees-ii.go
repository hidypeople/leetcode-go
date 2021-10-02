package task0095

import . "leetcode/binaryTree"

// Given an integer n, return all the structurally unique BST's (binary search trees),
// which has exactly n nodes of unique values from 1 to n. Return the answer in any order.
//
// Constraints:
//   1 <= n <= 8
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Binary Search Trees II.
//   Memory Usage: 4.6 MB, less than 22.57% of Go online submissions for Unique Binary Search Trees II.
func generateTrees(n int) []*TreeNode {
	result := []*TreeNode{}
	cache := &map[int][]*TreeNode{}
	for i := 1; i <= n; i++ {
		// We task 1,2,...,n as root and go to the deeper levels
		result = append(result, generate(cache, i, 1, n)...)
	}
	return result
}

// Returns the list of all possible trees constist of [min,max] values and starts from root "curr"
func generate(cache *map[int][]*TreeNode, curr, min, max int) []*TreeNode {
	// On each level we check cache that contains key for "curr+min+max" -> list of tree nodes
	cacheKey := min + max<<4 + curr<<8
	if trees, ok := (*cache)[cacheKey]; ok {
		return trees
	}
	result := []*TreeNode{}

	// Collect left trees
	leftTrees := []*TreeNode{}
	if curr == min {
		leftTrees = append(leftTrees, nil)
	} else {
		for i := curr - 1; i >= min; i-- {
			leftTrees = append(leftTrees, generate(cache, i, min, curr-1)...)
		}
	}

	// Collect right trees
	rightTrees := []*TreeNode{}
	if curr == max {
		rightTrees = append(rightTrees, nil)
	} else {
		for i := curr + 1; i <= max; i++ {
			rightTrees = append(rightTrees, generate(cache, i, curr+1, max)...)
		}
	}

	// Append all combinations of left and right trees
	for _, leftTree := range leftTrees {
		for _, rightTree := range rightTrees {
			result = append(result, &TreeNode{
				Val:   curr,
				Left:  leftTree,
				Right: rightTree,
			})
		}
	}
	(*cache)[cacheKey] = result
	return result
}
