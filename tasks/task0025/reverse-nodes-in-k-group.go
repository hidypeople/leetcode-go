package task0025

import . "leetcode/linkedList"

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
// k is a positive integer and is less than or equal to the length of the linked list.
// If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
// Constraints:
//   The number of nodes in the list is in the range sz.
//   1 <= sz <= 5000
//   0 <= Node.val <= 1000
//   1 <= k <= sz
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Nodes in k-Group.
//   Memory Usage: 3.7 MB, less than 100.00% of Go online submissions for Reverse Nodes in k-Group.
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}
	currentHead, current, firstInGroup := fakeHead, head, head
	i := 0
	for current != nil {
		if i < k-1 {
			// If haven't come to the end of current group yet: keep moving
			current = current.Next
			i++
		} else {
			// We came to the end of current group: now we need to cut off the tail,
			// revert the whole group and update glue everything back:
			// [... -> currentHead -> ... group... -> tail -> ...]
			//   =>
			// [... -> ... reverse(group) (currentHead points on last group element) ... -> firstInGroup -> ...]
			lastInGroup := current
			tail := current.Next
			lastInGroup.Next = nil
			reverseList(firstInGroup)
			currentHead.Next = lastInGroup
			firstInGroup.Next = tail
			currentHead = firstInGroup
			current = tail
			firstInGroup = tail
			i = 0
		}
	}
	return fakeHead.Next
}

func reverseList(head *ListNode) *ListNode {
	curr, reverse := head, head
	var prev *ListNode = nil
	for curr.Next != nil {
		curr = curr.Next
		reverse.Next = prev
		prev = reverse
		reverse = curr
	}
	reverse.Next = prev
	return reverse
}
