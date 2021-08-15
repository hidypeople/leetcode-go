package task0002

import . "leetcode/linkedList"

// Add two numbers represened as linked lists
// reversed: number 123 => linkedlist: 3->2->1)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
