package task0876

import . "leetcode/linkedList"

// Find middle node of linked list: [1,2,3,4] -> 3, [1,2,3,4,5] -> 3
//
// Constraints:
//   The number of nodes in the list is in the range [1, 100].
//   1 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
//   Memory Usage: 2 MB, less than 11.59% of Go online submissions for Middle of the Linked List.
func middleNode(head *ListNode) *ListNode {
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
