package tasks

import (
	. "leetcode/linkedList"
)

// Given the head of a linked list, remove the n-th node from the
// end of the list and return its head.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
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
