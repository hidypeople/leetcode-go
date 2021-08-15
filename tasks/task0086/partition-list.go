package task0086

import . "leetcode/linkedList"

// Given the head of a linked list and a value x, partition it such that
// all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in
// each of the two partitions.
func Partition(head *ListNode, x int) *ListNode {
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
