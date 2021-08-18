package task0206

import . "leetcode/linkedList"

// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// Constraints:
//   The number of nodes in the list is the range [0, 5000].
//   -5000 <= Node.val <= 5000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
//   Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Reverse Linked List.
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, reverse := head, head
	var prev *ListNode = nil
	for curr.Next != nil {
		curr = curr.Next
		reverse.Next = prev
		prev = reverse
		reverse = curr
	}
	reverse.Next = prev
	return reverse
}
