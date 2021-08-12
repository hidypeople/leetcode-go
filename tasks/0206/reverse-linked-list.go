package tasks

import . "leetcode/linkedList"

// Reverse linked list
func ReverseList(head *ListNode) *ListNode {
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
