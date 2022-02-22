package task0147

import . "leetcode/linkedList"

// Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head.
// The steps of the insertion sort algorithm:
//   Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
//   At each iteration, insertion sort removes one element from the input data, finds the location it
//   belongs within the sorted list and inserts it there.
//   It repeats until no input elements remain.
// The following is a graphical example of the insertion sort algorithm.
// The partially sorted list (black) initially contains only the first element in the list.
// One element (red) is removed from the input data and inserted in-place into the sorted list with each iteration.
//
// Constraints:
//   The number of nodes in the list is in the range [1, 5000].
//   -5000 <= Node.val <= 5000
//
// Results:
//   Runtime: 4 ms, faster than 98.04% of Go online submissions for Insertion Sort List.
//   Memory Usage: 3.3 MB, less than 100.00% of Go online submissions for Insertion Sort List.
func insertionSortList(head *ListNode) *ListNode {
	newHead, prev, curr := head, head, head.Next
	for curr != nil {
		if curr.Val >= prev.Val {
			// if current is greater that prevous => do not touch it
			curr = curr.Next
			prev = prev.Next
			continue
		}
		// Current is lesser than previous => we need to find the place behind the prevous
		prev.Next = curr.Next
		if curr.Val < newHead.Val {
			// If current is lesser than head => it will become new head
			curr.Next = newHead
			newHead = curr
		} else {
			// Find the p and c where p.Val < curr.Val <= c.Val
			p, c := newHead, newHead.Next
			for c.Val < curr.Val {
				c = c.Next
				p = p.Next
			}
			// Insert current
			p.Next = curr
			curr.Next = c
		}
		curr = prev.Next
	}
	return newHead
}
