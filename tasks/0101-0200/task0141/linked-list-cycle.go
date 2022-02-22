package task0141

import (
	. "leetcode/linkedList"
)

// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer. Internally, pos is used to denote the
// index of the node that tail's next pointer is connected to. Note that pos is not passed
// as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
// Constraints:
//   The number of the nodes in the list is in the range [0, 10^4].
//   -10^5 <= Node.val <= 10^5
//   pos is -1 or a valid index in the linked-list.
//
// Results:
//   Runtime: 8 ms, faster than 88.40% of Go online submissions for Linked List Cycle.
//   Memory Usage: 4.4 MB, less than 100.00% of Go online submissions for Linked List Cycle.
func hasCycle(head *ListNode) bool {
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
