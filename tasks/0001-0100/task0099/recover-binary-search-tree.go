package task0099

import . "leetcode/binaryTree"

var first *TreeNode
var second *TreeNode
var prev *TreeNode

// You are given the root of a binary search tree (BST), where the values of exactly two nodes
// of the tree were swapped by mistake. Recover the tree without changing its structure.
//
// Constraints:
//   The number of nodes in the tree is in the range [2, 1000].
//   -2^31 <= Node.val <= 2^31 - 1
//
// Results:
//   Runtime: 4 ms, faster than 97.90% of Go online submissions for Recover Binary Search Tree.
//   Memory Usage: 6 MB, less than 61.54% of Go online submissions for Recover Binary Search Tree.
func recoverTree(root *TreeNode) {
	first = nil
	second = nil
	prev = &TreeNode{Val: -1 << 63}
	traverse(root)
	first.Val, second.Val = second.Val, first.Val
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)

	if first == nil && prev.Val >= root.Val {
		first = prev
	}
	if first != nil && prev.Val >= root.Val {
		second = root
	}
	prev = root

	traverse(root.Right)
}
