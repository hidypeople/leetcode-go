package task0021

import . "leetcode/linkedList"

// Merge two sorted linked lists and return it as a sorted list.
// The list should be made by splicing together the nodes of the first two lists.
//
// Constraints:
//   The number of nodes in both lists is in the range [0, 50].
//   -100 <= Node.val <= 100
//   Both l1 and l2 are sorted in non-decreasing order.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
//   Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.
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
