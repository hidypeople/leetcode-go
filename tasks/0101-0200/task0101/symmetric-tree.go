package task0101

import . "leetcode/binaryTree"

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 1000].
//   -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
//   Memory Usage: 2.9 MB, less than 38.33% of Go online submissions for Symmetric Tree.
func isSymmetric(node *TreeNode) bool {
	if node == nil {
		return true
	}
	return isSymmetricEqual(node.Left, node.Right)
}

// Is one tree equals to reflected another tree
func isSymmetricEqual(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	return tree1.Val == tree2.Val &&
		isSymmetricEqual(tree1.Left, tree2.Right) &&
		isSymmetricEqual(tree1.Right, tree2.Left)
}
