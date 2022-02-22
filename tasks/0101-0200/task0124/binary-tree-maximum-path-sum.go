package task0124

import . "leetcode/binaryTree"

var MIN int = -1 << 32

// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes
// in the sequence has an edge connecting them. A node can only appear in the sequence
// at most once. Note that the path does not need to pass through the root.
// The path sum of a path is the sum of the node's values in the path.
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 3 * 10^4].
//   -1000 <= Node.val <= 1000
//
// Results:
//   Runtime: 12 ms, faster than 98.62% of Go online submissions for Binary Tree Maximum Path Sum.
//   Memory Usage: 8 MB, less than 30.00% of Go online submissions for Binary Tree Maximum Path Sum.
func maxPathSum(root *TreeNode) int {
	resultFull, _ := traverse(root)
	return resultFull
}

// Function returns 2 values for given node:
// 1. The maximum path sum for paths that have "Left - Node - Right" forks (we cannot use it for parent nodes)
// 2. The maximum path sum for paths without "Left - Node - Right" forks (we can use it for parent nodes)
func traverse(node *TreeNode) (int, int) {
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val
	}

	leftFull, leftPath, rightFull, rightPath := MIN, MIN, MIN, MIN
	if node.Left != nil {
		leftFull, leftPath = traverse(node.Left)
	}
	if node.Right != nil {
		rightFull, rightPath = traverse(node.Right)
	}

	// Path can contain only current single node (if subtrees have negative sums)
	resultFull, resultPath := node.Val, node.Val
	if leftPath > 0 {
		resultFull += leftPath
		resultPath = leftPath + node.Val
	}
	if rightPath > 0 {
		resultFull += rightPath
		resultPath = max(resultPath, rightPath+node.Val)
	}
	resultFull = max(resultFull, leftFull)
	resultFull = max(resultFull, rightFull)

	return resultFull, resultPath
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
