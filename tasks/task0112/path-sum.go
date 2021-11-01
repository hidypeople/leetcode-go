package task0112

import . "leetcode/binaryTree"

// Given the root of a binary tree and an integer targetSum, return true if the tree has a
// root-to-leaf path such that adding up all the values along the path equals targetSum.
// A leaf is a node with no children.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 5000].
//   -1000 <= Node.val <= 1000
//   -1000 <= targetSum <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Path Sum.
//   Memory Usage: 4.7 MB, less than 47.86% of Go online submissions for Path Sum.
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return traverse(root, targetSum)
}

func traverse(root *TreeNode, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	if (root.Left != nil && traverse(root.Left, targetSum-root.Val)) ||
		(root.Right != nil && traverse(root.Right, targetSum-root.Val)) {
		return true
	}
	return false
}
