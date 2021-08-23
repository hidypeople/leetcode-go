package task0226

import . "leetcode/binaryTree"

// Given the root of a binary tree, invert the tree, and return its root.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 100].
//   -100 <= Node.val <= 100
//
// Result:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
//   Memory Usage: 2.1 MB, less than 50.25% of Go online submissions for Invert Binary Tree.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
