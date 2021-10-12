package task0102

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 2000].
//   -1000 <= Node.val <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
//   Memory Usage: 2.8 MB, less than 73.23% of Go online submissions for Binary Tree Level Order Traversal.
func levelOrder(root *TreeNode) [][]int {
	// Complexity: O(N)
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	levelNodes := []*TreeNode{root}
	for len(levelNodes) > 0 {
		resultLevel := []int{}
		nextLevelNodes := []*TreeNode{}
		// Append values from levelNodes into result + make next level nodes array
		for _, node := range levelNodes {
			resultLevel = append(resultLevel, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		result = append(result, resultLevel)
		levelNodes = nextLevelNodes
	}

	return result
}
