package task0086

import . "leetcode/linkedList"

// Given the head of a linked list and a value x, partition it such that
// all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in
// each of the two partitions.
//
// Constraints:
//   The number of nodes in the list is in the range [0, 200].
//   -100 <= Node.val <= 100
//   -200 <= x <= 200
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Partition List.
//   Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Partition List.
func partition(head *ListNode, x int) *ListNode {
	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}

	var lessHead, lessCurr, greaterHead, greaterCurr *ListNode = nil, nil, nil, nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = nil
		if curr.Val < x {
			// send to lessList
			if lessHead == nil {
				lessHead, lessCurr = curr, curr
			} else {
				lessCurr.Next = curr
				lessCurr = curr
			}
		} else {
			// send to greater list
			if greaterHead == nil {
				greaterHead, greaterCurr = curr, curr
			} else {
				greaterCurr.Next = curr
				greaterCurr = curr
			}
		}
		curr = next
	}
	if lessHead != nil {
		fakeHead.Next = lessHead
		if greaterHead != nil {
			lessCurr.Next = greaterHead
		}
	} else if greaterHead != nil {
		fakeHead.Next = greaterHead
	}
	return fakeHead.Next
}
