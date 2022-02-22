package task0145

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
// Constraints:
//    The number of nodes in the tree is in the range [0, 100].
//    -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
//   Memory Usage: 2.1 MB, less than 30.67% of Go online submissions for Binary Tree Postorder Traversal.
func postorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	result := []int{}
	if node.Left != nil {
		result = append(result, postorderTraversal(node.Left)...)
	}
	if node.Right != nil {
		result = append(result, postorderTraversal(node.Right)...)
	}
	result = append(result, node.Val)
	return result
}
