package task1290

import . "leetcode/linkedList"

// Given head which is a reference node to a singly-linked list.
// The value of each node in the linked list is either 0 or 1. The linked list holds the
// binary representation of a number.
// Return the decimal value of the number in the linked list.
//
// Constraints:
//   The Linked List is not empty.
//   Number of nodes will not exceed 30.
//   Each node's value is either 0 or 1.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Convert Binary Number in a Linked List to Integer.
//   Memory Usage: 2 MB, less than 59.38% of Go online submissions for Convert Binary Number in a Linked List to Integer.
func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = (result << 1) | head.Val
		head = head.Next
	}
	return result
}
