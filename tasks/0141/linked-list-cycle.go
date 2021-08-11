package tasks

import (
	. "leetcode/linkedList"
)

// Is list node has cycle
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// slow does 1 step, fast does 2 steps
		slow = slow.Next
		fast = fast.Next.Next

		// if fast made a cycle and met with slow => cycle
		if fast == slow {
			return true
		}
	}
	return false
}
