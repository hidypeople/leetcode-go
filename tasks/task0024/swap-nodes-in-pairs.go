package task0024

import . "leetcode/linkedList"

// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)
//
// Constraints:
//   The number of nodes in the list is in the range [0, 100].
//   0 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
//   Memory Usage: 2.1 MB, less than 16.14% of Go online submissions for Swap Nodes in Pairs.
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, curr := head.Next, head
	var prev *ListNode = nil
	for curr != nil && curr.Next != nil {
		next := curr.Next
		nextNext := next.Next
		curr.Next = nextNext
		next.Next = curr
		if prev != nil {
			prev.Next = next
		}
		prev = curr
		curr = curr.Next
	}
	return newHead
}
