package task0142

import . "leetcode/linkedList"

// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again
// by continuously following the next pointer. Internally, pos is used to denote the index of the
// node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
// Notice that you should not modify the linked list.
//
// Constraints:
//   The number of the nodes in the list is in the range [0, 10^4].
//   -10^5 <= Node.val <= 10^5
//   pos is -1 or a valid index in the linked-list.
//
// Results:
//   Runtime: 4 ms, faster than 97.89% of Go online submissions for Linked List Cycle II.
//   Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Linked List Cycle II.
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// slow does 1 step, fast does 2 steps
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
