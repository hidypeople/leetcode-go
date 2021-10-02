package task0725

import . "leetcode/linkedList"

// Given the head of a singly linked list and an integer k, split the linked list into k consecutive linked list parts.
// The length of each part should be as equal as possible: no two parts should have a size differing by more than one.
// This may lead to some parts being null.
// The parts should be in the order of occurrence in the input list, and parts occurring earlier should always have
// a size greater than or equal to parts occurring later.
// Return an array of the k parts.
//
// Constraints:
//   The number of nodes in the list is in the range [0, 1000].
//   0 <= Node.val <= 1000
//   1 <= k <= 50
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Split Linked List in Parts.
//   Memory Usage: 2.8 MB, less than 34.78% of Go online submissions for Split Linked List in Parts.
func splitListToParts(head *ListNode, k int) []*ListNode {

	// Calculate the count of list items
	n := 0
	curr := head
	for curr != nil {
		n++
		curr = curr.Next
	}

	// We can have 2 group sizes: n/k+1 and n/k, larger groups should be first
	groupSizeMin := n / k
	groupSizeMax := groupSizeMin
	if n%k != 0 {
		groupSizeMax++
	}
	currentGroupSize := groupSizeMax
	countOfMaxSizedGroups := n - groupSizeMin*k

	result := make([]*ListNode, k)
	curr = head
	for i := 0; i < k; i++ {
		if curr == nil {
			break
		}
		result[i] = curr
		prev := curr
		for i := 0; i < currentGroupSize; i++ {
			prev = curr
			curr = curr.Next
		}
		prev.Next = nil

		countOfMaxSizedGroups--
		if countOfMaxSizedGroups == 0 {
			// Update current group size
			currentGroupSize--
		}

	}
	return result
}
