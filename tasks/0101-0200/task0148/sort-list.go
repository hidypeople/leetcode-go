package task0148

import . "leetcode/linkedList"

// Given the head of a linked list, return the list after sorting it in ascending order.
//
// Constraints:
//   The number of nodes in the list is in the range [0, 5 * 104].
//   -10^5 <= Node.val <= 10^5
//
// Results:
//   Runtime: 24 ms, faster than 98.22% of Go online submissions for Sort List.
//   Memory Usage: 7 MB, less than 100.00% of Go online submissions for Sort List.
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Merge sort:
	//   1. Split list into 2 halves
	//   2. Sort each half
	//   3. Merge it into the one list
	// Complexity: O(N*lnN)
	headRight := split(head)
	leftResult, rightResult := sortList(head), sortList(headRight)
	return merge(leftResult, rightResult)
}

func split(head *ListNode) *ListNode {
	slow, fast := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func merge(left, right *ListNode) *ListNode {
	preHead := &ListNode{}
	curr := preHead
	l, r := left, right
	for l != nil && r != nil {
		if l.Val < r.Val {
			curr.Next = l
			l = l.Next
		} else {
			curr.Next = r
			r = r.Next
		}
		curr = curr.Next
	}
	if l != nil {
		curr.Next = l
	}
	if r != nil {
		curr.Next = r
	}

	return preHead.Next
}
