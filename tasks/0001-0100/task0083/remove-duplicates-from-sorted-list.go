package task0083

import . "leetcode/linkedList"

// Given the head of a sorted linked list, delete all duplicates
// such that each element appears only once. Return the linked list sorted as well.
//
// Constraints:
//   The number of nodes in the list is in the range [0, 300].
//   -100 <= Node.val <= 100
//   The list is guaranteed to be sorted in ascending order.
//
// Results:
//   Runtime: 4 ms, faster than 80.59% of Go online submissions for Remove Duplicates from Sorted List.
//   Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicates(head *ListNode) *ListNode {
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
