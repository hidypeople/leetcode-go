package task0203

import . "leetcode/linkedList"

// Given the head of a linked list and an integer val, remove all
// the nodes of the linked list that has Node.val == val, and return
// the new head.
//
// Constraints:
//   The number of nodes in the list is in the range [0, 10^4].
//   1 <= Node.val <= 50
//   0 <= val <= 50
//
// Results:
//   Runtime: 8 ms, faster than 86.15% of Go online submissions for Remove Linked List Elements.
//   Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.
func removeElements(head *ListNode, val int) *ListNode {
	// First fake element needed to handle nil results
	result := &ListNode{
		Val:  0,
		Next: head,
	}

	curr := result
	for curr.Next != nil {
		if curr.Next.Val == val {
			// skip nodes with "val"
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return result.Next
}
