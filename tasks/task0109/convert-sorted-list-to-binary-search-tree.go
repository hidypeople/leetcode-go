package task0109

import (
	. "leetcode/binaryTree"
	. "leetcode/linkedList"
)

// Given the head of a singly linked list where elements are sorted in ascending order,
// convert it to a height balanced BST.
// For this problem, a height-balanced binary tree is defined as a binary tree in
// which the depth of the two subtrees of every node never differ by more than 1.
//
// Constraints:
//   The number of nodes in head is in the range [0, 2 * 104].
//   -105 <= Node.val <= 105
func SortedListToBST(head *ListNode) *TreeNode {
	return &TreeNode{
		Val:   head.Val,
		Left:  nil,
		Right: nil,
	}
}
