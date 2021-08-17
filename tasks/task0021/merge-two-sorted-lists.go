package task0021

import . "leetcode/linkedList"

// Merge two ordered sorted lists into one ordered sorted list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil
	curr1, curr2 := l1, l2
	var curr *ListNode = nil
	for curr1 != nil || curr2 != nil {
		var next *ListNode
		if curr1 == nil || (curr2 != nil && curr2.Val < curr1.Val) {
			next = curr2
			curr2 = curr2.Next
		} else {
			next = curr1
			curr1 = curr1.Next
		}
		if result == nil {
			result, curr = next, next
		} else {
			curr.Next = next
			curr = next
		}
	}
	return result
}
