package task0100

import . "leetcode/binaryTree"

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//
// Constraints:
//   The number of nodes in both trees is in the range [0, 100].
//   -10^4 <= Node.val <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
//   Memory Usage: 2.1 MB, less than 20.08% of Go online submissions for Same Tree.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	// Мы знаем что p и q не пустые
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
