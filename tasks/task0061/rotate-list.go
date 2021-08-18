package task0061

import (
	. "leetcode/linkedList"
)

// Given the head of a linked list, rotate the list to the right by k places.
//
// Constraints:
//   The number of nodes in the list is in the range [0, 500].
//   -100 <= Node.val <= 100
//   0 <= k <= 2 * 10^9
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate List.
//   Memory Usage: 2.6 MB, less than 29.00% of Go online submissions for Rotate List.
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	n := 1
	for curr.Next != nil {
		n++
		curr = curr.Next
	}

	n = n - k%n - 1

	curr.Next = head // make cycle
	curr = head
	i := 0
	for i < n {
		curr = curr.Next
		i++
	}
	result := curr.Next
	curr.Next = nil
	return result
}
