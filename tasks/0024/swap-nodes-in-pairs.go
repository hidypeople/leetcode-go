package tasks

import . "leetcode/linkedList"

// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)
func SwapPairs(head *ListNode) *ListNode {
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
