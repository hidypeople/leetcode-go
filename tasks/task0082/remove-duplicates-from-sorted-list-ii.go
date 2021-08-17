package task0082

import . "leetcode/linkedList"

// Given the head of a sorted linked list, delete all nodes that have duplicate
// numbers, leaving only distinct numbers from the original list. Return the
// linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	// This node will never go into the result.
	// It is used for easier nil handling
	result := &ListNode{
		Val:  0,
		Next: head,
	}

	curr, prev := result, result
	for curr.Next != nil {
		isSkippingNext := false
		for curr.Next.Next != nil && curr.Next.Val == curr.Next.Next.Val {
			curr = curr.Next
			isSkippingNext = true
		}
		if isSkippingNext {
			prev.Next = curr.Next.Next
			curr = prev
		} else {
			prev = prev.Next
			curr = curr.Next
		}
	}

	return result.Next
}
