package task0328

import . "leetcode/linkedList"

// Given the head of a singly linked list, group all the nodes with odd indices together
// followed by the nodes with even indices, and return the reordered list.
// The first node is considered odd, and the second node is even, and so on.
// Note that the relative order inside both the even and odd groups should remain as it was in the input.
// You must solve the problem in O(1) extra space complexity and O(n) time complexity.
//
// Constraints:
//   n == number of nodes in the linked list
//   0 <= n <= 10^4
//   -10^6 <= Node.val <= 10^6
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Odd Even Linked List.
//   Memory Usage: 3.3 MB, less than 43.65% of Go online submissions for Odd Even Linked List.
func oddEvenList(head *ListNode) *ListNode {
	var evenRoot *ListNode = nil
	var evenCurrent *ListNode = nil
	if head != nil {
		evenRoot, evenCurrent = head.Next, head.Next
	}

	odd := head
	for odd != nil {
		even := odd.Next
		if even == nil {
			// only odd left
			odd.Next = evenRoot
			break
		} else {
			odd.Next = odd.Next.Next
			if evenCurrent != even {
				evenCurrent.Next = even
				evenCurrent = evenCurrent.Next
			}
			if odd.Next == nil {
				odd.Next = evenRoot
				break
			}
			odd = odd.Next
		}
	}

	if evenCurrent != nil {
		// Cut off last even tail
		evenCurrent.Next = nil
	}
	return head
}
