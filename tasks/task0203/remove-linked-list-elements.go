package task0203

import . "leetcode/linkedList"

// Given the head of a linked list and an integer val, remove all
// the nodes of the linked list that has Node.val == val, and return
// the new head.
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
