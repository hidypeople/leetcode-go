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
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return processNode(head)
}

// Will do it recursively
func processNode(head *ListNode) *TreeNode {
	if head.Next == nil {
		return &TreeNode{
			Val:   head.Val,
			Left:  nil,
			Right: nil,
		}
	}

	left := MiddleNodePrev(head)
	middle := left.Next
	right := middle.Next

	left.Next = nil // cut off left tree

	result := &TreeNode{
		Val:   middle.Val,
		Left:  processNode(head),
		Right: nil,
	}
	if right != nil {
		result.Right = processNode(right)
	}
	return result
}

// Find previous of middle node: [1,2,3,4] -> 2, [1,2,3,4,5] -> 2
func MiddleNodePrev(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil && fast.Next.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
