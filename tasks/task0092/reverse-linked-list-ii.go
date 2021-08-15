package task0092

import . "leetcode/linkedList"

//Given the head of a singly linked list and two integers left and right where left <= right,
//reverse the nodes of the list from position left to position right, and return the reversed list.
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}

	// skip beyond left
	prevGlobal := fakeHead
	for i := 1; i < left; i++ {
		prevGlobal = prevGlobal.Next
	}

	curr, reverse := prevGlobal.Next, prevGlobal.Next
	var prev *ListNode = nil
	last := curr
	for i := left; curr != nil && i <= right; i++ {
		curr = curr.Next
		reverse.Next = prev
		prev = reverse
		reverse = curr
	}
	prevGlobal.Next = prev
	if last != nil {
		last.Next = curr
	}

	return fakeHead.Next
}
