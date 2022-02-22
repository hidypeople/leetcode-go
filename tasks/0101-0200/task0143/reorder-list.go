package task0143

import . "leetcode/linkedList"

// You are given the head of a singly linked-list. The list can be represented as:
//   L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:
//   L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.
//
// Constraints:
//   The number of nodes in the list is in the range [1, 5 * 10^4].
//   1 <= Node.val <= 1000
//
// Results:
//   Runtime: 4 ms, faster than 100.00% of Go online submissions for Reorder List.
//   Memory Usage: 5.4 MB, less than 100.00% of Go online submissions for Reorder List.
func reorderList(head *ListNode) {

	// Find the middle point of LL
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// Reverse right part
	curr := slow
	var prev *ListNode = nil
	for curr != nil {
		nextNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNext
	}

	// Go through the both lists and merge them
	l, r := head, prev
	for r != nil && l != r && l.Next != r {
		l_next, r_next := l.Next, r.Next
		l.Next = r
		r.Next = l_next
		l = l_next
		r = r_next
	}
}
