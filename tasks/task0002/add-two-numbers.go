package task0002

import . "leetcode/linkedList"

// You are given two non-empty linked lists representing two non-negative integers. The digits
// are stored in reverse order, and each of their nodes contains a single digit. Add the two
// numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Constraints:
//   The number of nodes in each linked list is in the range [1, 100].
//   0 <= Node.val <= 9
//   It is guaranteed that the list represents a number that does not have leading zeros.
//
// Results:
//   Runtime: 4 ms, faster than 98.33% of Go online submissions for Add Two Numbers.
//   Memory Usage: 4.8 MB, less than 92.75% of Go online submissions for Add Two Numbers.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode = nil
	var current *ListNode = nil
	curr1, curr2 := l1, l2
	var remain int = 0
	for curr1 != nil || curr2 != nil {
		var sum int = remain
		if curr1 != nil {
			sum += curr1.Val
			curr1 = curr1.Next
		}
		if curr2 != nil {
			sum += curr2.Val
			curr2 = curr2.Next
		}
		remain = sum / 10

		next := &ListNode{Val: sum % 10, Next: nil}
		if root == nil {
			current, root = next, next
		} else {
			current.Next = next
			current = next
		}
	}
	if remain > 0 {
		current.Next = &ListNode{Val: remain, Next: nil}
	}
	return root
}
