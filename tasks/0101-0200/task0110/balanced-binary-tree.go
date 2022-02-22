package task0110

import . "leetcode/binaryTree"

// Given a binary tree, determine if it is height-balanced.
// For this problem, a height-balanced binary tree is defined as:
// a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 5000].
//   -10^4 <= Node.val <= 10^4
//
// Results:
//   Runtime: 4 ms, faster than 96.66% of Go online submissions for Balanced Binary Tree.
//   Memory Usage: 5.9 MB, less than 99.33% of Go online submissions for Balanced Binary Tree.
func isBalanced(root *TreeNode) bool {
	return traverse(root) != -1
}

// Traverse through the tree node. If left and right subnodes of the current node
// or their subnodes has maximum heights that differs more than 1 - we return -1,
// otherwise we return maximum height
func traverse(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// Calculate left subtree size
	var leftMaxDepth int
	if node.Left != nil {
		leftMaxDepth = traverse(node.Left)
		if leftMaxDepth < 0 {
			// Left subtree is not balanced
			return -1
		}
	}

	// Calculate right subtree size
	var rightMaxDepth int
	if node.Right != nil {
		rightMaxDepth = traverse(node.Right)
		if rightMaxDepth < 0 {
			// Right subtree is not balanced
			return -1
		}
	}

	if abs(leftMaxDepth-rightMaxDepth) > 1 {
		// Current tree node is not balanced
		return -1
	}

	// Return the max depth
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	} else {
		return rightMaxDepth + 1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
