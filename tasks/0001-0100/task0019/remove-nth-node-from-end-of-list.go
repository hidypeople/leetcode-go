package task0019

import (
	. "leetcode/linkedList"
)

// Given the head of a linked list, remove the n-th node from the
// end of the list and return its head.
//
// Constraints:
//   The number of nodes in the list is sz.
//   1 <= sz <= 30
//   0 <= Node.val <= 100
//   1 <= n <= sz
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
//   Memory Usage: 2.5 MB, less than 6.36% of Go online submissions for Remove Nth Node From End of List.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	curr := head
	cache := make(map[int]*ListNode)
	i := 0
	for curr != nil {
		cache[i] = curr
		curr = curr.Next
		i++
	}
	if i == n {
		return head.Next
	}
	cache[i-n-1].Next = cache[i-n+1]
	return head
}
