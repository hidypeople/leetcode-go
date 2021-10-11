package task0543

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
// This path may or may not pass through the root.
// The length of a path between two nodes is represented by the number of edges between them.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 10^4].
//   -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Diameter of Binary Tree.
//   Memory Usage: 4.6 MB, less than 31.10% of Go online submissions for Diameter of Binary Tree.
func diameterOfBinaryTree(root *TreeNode) int {
	_, result := traverseTree(root)
	return result
}

// This function returns 2 values for the given node: maximum depth and diameter
func traverseTree(node *TreeNode) (int, int) {
	if node.Left == nil && node.Right == nil {
		return 1, 0
	} else if node.Right == nil {
		leftDepth, leftDiameter := traverseTree(node.Left)
		return 1 + leftDepth, max(leftDiameter, leftDepth)
	} else if node.Left == nil {
		rightDepth, rightDiameter := traverseTree(node.Right)
		return 1 + rightDepth, max(rightDiameter, rightDepth)
	} else {
		leftDepth, leftDiameter := traverseTree(node.Left)
		rightDepth, rightDiameter := traverseTree(node.Right)
		return 1 + max(leftDepth, rightDepth), max(max(leftDiameter, rightDiameter), leftDepth+rightDepth)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
