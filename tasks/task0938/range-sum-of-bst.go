package task0938

import . "leetcode/binaryTree"

// Given the root node of a binary search tree and two integers low and high,
// return the sum of values of all nodes with a value in the inclusive range [low, high].
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 2 * 10^4].
//   1 <= Node.val <= 10^5
//   1 <= low <= high <= 10^5
//   All Node.val are unique.
//
// Results:
//   Runtime: 84 ms, faster than 97.42% of Go online submissions for Range Sum of BST.
//   Memory Usage: 9.6 MB, less than 31.73% of Go online submissions for Range Sum of BST.
func rangeSumBST(node *TreeNode, low int, high int) int {
	if node == nil {
		return 0
	}
	result := 0
	if low <= node.Val && node.Val <= high {
		result = node.Val
	}
	result += rangeSumBST(node.Right, low, high)
	result += rangeSumBST(node.Left, low, high)
	return result
}
