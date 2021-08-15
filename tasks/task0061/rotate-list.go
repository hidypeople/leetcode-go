package task0061

import (
	. "leetcode/linkedList"
)

//Given the head of a linked list, rotate the list to the right by k places.
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	n := 1
	for curr.Next != nil {
		n++
		curr = curr.Next
	}

	n = n - k%n - 1

	curr.Next = head // make cycle
	curr = head
	i := 0
	for i < n {
		curr = curr.Next
		i++
	}
	result := curr.Next
	curr.Next = nil
	return result
}
