package task0538

import . "leetcode/binaryTree"

// Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of
// the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.
// As a reminder, a binary search tree is a tree that satisfies these constraints:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 10^4].
//   -10^4 <= Node.val <= 10^4
//   All the values in the tree are unique.
//   root is guaranteed to be a valid binary search tree.
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Convert BST to Greater Tree.
//   Memory Usage: 6.7 MB, less than 80.60% of Go online submissions for Convert BST to Greater Tree.
func convertBST(root *TreeNode) *TreeNode {
	traverse(root, 0)
	return root
}

func traverse(node *TreeNode, currentSum int) int {
	if node == nil {
		return currentSum
	}
	node.Val += traverse(node.Right, currentSum)
	currentSum = traverse(node.Left, node.Val)
	return currentSum
}
