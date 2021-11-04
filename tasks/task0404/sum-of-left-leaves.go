package task0404

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the sum of all left leaves.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 1000].
//   -1000 <= Node.val <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum of Left Leaves.
//   Memory Usage: 2.6 MB, less than 98.21% of Go online submissions for Sum of Left Leaves.
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			} else {
				stack = append(stack, node.Left)
			}
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return sum
}
