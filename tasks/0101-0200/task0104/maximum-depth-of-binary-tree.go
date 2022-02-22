package task0104

import . "leetcode/binaryTree"

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 10^4].
//   -100 <= Node.val <= 100
//
// Results:
//   Runtime: 4 ms, faster than 89.81% of Go online submissions for Maximum Depth of Binary Tree.
//   Memory Usage: 4.3 MB, less than 54.33% of Go online submissions for Maximum Depth of Binary Tree.
func maxDepth(root *TreeNode) int {
	// I don't know how to make performarmance 100%, because it's already optimal
	if root == nil {
		return 0
	}
	var leftMaxDepth int
	var rightMaxDepth int
	if root.Left != nil {
		leftMaxDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightMaxDepth = maxDepth(root.Right)
	}
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	} else {
		return rightMaxDepth + 1
	}
}
