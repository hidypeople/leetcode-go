package task0876

import . "leetcode/linkedList"

// Find middle node of linked list: [1,2,3,4] -> 3, [1,2,3,4,5] -> 3
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// Find middle node: [1,2,3,4] -> 2, [1,2,3,4,5] -> 3
func middleNodePrev(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast != nil && fast.Next != nil {
			slow = slow.Next
		}
	}
	return slow
}
