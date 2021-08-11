package tasks

import ll "leetcode/linkedList"

func AddTwoNumbers(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	var root *ll.ListNode = nil
	var current *ll.ListNode = nil
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

		next := &ll.ListNode{Val: sum % 10, Next: nil}
		if root == nil {
			current, root = next, next
		} else {
			current.Next = next
			current = next
		}
	}
	if remain > 0 {
		current.Next = &ll.ListNode{Val: remain, Next: nil}
	}
	return root
}
