package task0129

import . "leetcode/binaryTree"

// You are given the root of a binary tree containing digits from 0 to 9 only.
// Each root-to-leaf path in the tree represents a number.
// For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
// Return the total sum of all root-to-leaf numbers. Test cases are generated so
// that the answer will fit in a 32-bit integer.
// A leaf node is a node with no children.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 1000].
//   0 <= Node.val <= 9
//   The depth of the tree will not exceed 10.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum Root to Leaf Numbers.
//   Memory Usage: 2.2 MB, less than 47.50% of Go online submissions for Sum Root to Leaf Numbers.
func sumNumbers(root *TreeNode) int {
	return traverse(0, root)
}

func traverse(currentNumber int, node *TreeNode) int {
	newCurrentNumber := currentNumber*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return newCurrentNumber
	}
	result := 0
	if node.Left != nil {
		result += traverse(newCurrentNumber, node.Left)
	}
	if node.Right != nil {
		result += traverse(newCurrentNumber, node.Right)
	}
	return result
}
