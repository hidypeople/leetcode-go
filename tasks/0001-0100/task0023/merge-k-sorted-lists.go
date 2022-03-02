package task0023

import . "leetcode/linkedList"

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Constraints:
//   k == lists.length
//   0 <= k <= 10^4
//   0 <= lists[i].length <= 500
//   -10^4 <= lists[i][j] <= 10^4
//   lists[i] is sorted in ascending order.
//   The sum of lists[i].length won't exceed 10^4.
//
// Results:
//   Runtime: 8 ms, faster than 92.45% of Go online submissions for Merge k Sorted Lists.
//   Memory Usage: 5.2 MB, less than 100.00% of Go online submissions for Merge k Sorted Lists.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 2 {
		return mergeList(lists[0], lists[1])
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 0 {
		return nil
	}
	median := len(lists) / 2
	return mergeList(mergeKLists(lists[:median]), mergeKLists(lists[median:]))
}

// Merge 2 lists
func mergeList(list1, list2 *ListNode) *ListNode {
	beforeFirstNode := &ListNode{}
	current := beforeFirstNode
	for list1 != nil && list2 != nil {
		if list2.Val < list1.Val {
			current.Next = list2
			list2 = list2.Next
			current = current.Next
		} else {
			current.Next = list1
			list1 = list1.Next
			current = current.Next
		}
	}
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return beforeFirstNode.Next
}
