package task0098

import . "leetcode/binaryTree"

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 10^4].
//   -2^31 <= Node.val <= 2^31 - 1
//
// Results:
//   Runtime: 3 ms, faster than 95.38% of Go online submissions for Validate Binary Search Tree.
//   Memory Usage: 5.3 MB, less than 47.76% of Go online submissions for Validate Binary Search Tree.
func isValidBST(root *TreeNode) bool {
	result := true
	if root.Left != nil {
		result = isValidNode(root.Left, nil, &root.Val)
	}
	if result && root.Right != nil {
		result = isValidNode(root.Right, &root.Val, nil)
	}
	return result
}

func isValidNode(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		return false
	}
	result := true
	if node.Left != nil {
		result = isValidNode(node.Left, min, &node.Val)
	}
	if result && node.Right != nil {
		result = isValidNode(node.Right, &node.Val, max)
	}
	return result
}
