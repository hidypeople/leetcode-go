package tasks

import ll "leetcode/linkedList"

func MergeTwoLists(l1 *ll.ListNodeInt, l2 *ll.ListNodeInt) *ll.ListNodeInt {
	var result *ll.ListNodeInt = nil
	curr1, curr2 := l1, l2
	var curr *ll.ListNodeInt = nil
	for curr1 != nil || curr2 != nil {
		var next *ll.ListNodeInt
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
