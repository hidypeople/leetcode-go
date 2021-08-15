package task0083

import . "leetcode/linkedList"

// Given the head of a sorted linked list, delete all duplicates
// such that each element appears only once. Return the linked list sorted as well.
func DeleteDuplicates(head *ListNode) *ListNode {
	// This node will never go into the result.
	// It is used for easier nil handling
	result := &ListNode{
		Val:  0,
		Next: head,
	}

	curr := result
	for curr.Next != nil {
		if curr.Next.Next != nil && curr.Next.Next.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return result.Next
}
