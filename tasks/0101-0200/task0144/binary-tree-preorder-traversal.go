package task0144

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
// Constraints:
//    The number of nodes in the tree is in the range [0, 100].
//    -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
//   Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
func preorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	result := []int{node.Val}
	if node.Left != nil {
		result = append(result, preorderTraversal(node.Left)...)
	}
	if node.Right != nil {
		result = append(result, preorderTraversal(node.Right)...)
	}
	return result
}
