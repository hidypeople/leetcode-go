package task0234

import (
	. "leetcode/linkedList"
)

// Given the head of a singly linked list, return true if it is a palindrome.
//
// Constraints:
//   The number of nodes in the list is in the range [1, 105].
//   0 <= Node.val <= 9
//
// Results:
//   Runtime: 148 ms, faster than 87.03% of Go online submissions for Palindrome Linked List.
//   Memory Usage: 11.4 MB, less than 19.92% of Go online submissions for Palindrome Linked List.
func isPalindrome(head *ListNode) bool {
	// get the middle point
	middle := middleNodePrev(head)

	// Split list into two chains: left and right
	left, right := middle, middle.Next
	if right == nil {
		return true
	}
	left.Next = nil // break the chain

	// Reverse right chain amd compare left and right
	right = reverseList(right)
	left = head
	for right != nil {
		if left == nil {
			return true
		}
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
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

// Find middle node: [1,2,3,4] -> 2, [1,2,3,4,5] -> 3
func middleNodePrev(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast != nil && fast.Next != nil {
			slow = slow.Next
		}
	}
	return slow
}
