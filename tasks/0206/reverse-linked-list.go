package tasks

import . "leetcode/linkedList"

// Reverse linked list
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

func Tests() {
	var list *ListNode

	list = nil
	LinkedListPrint(list)
	LinkedListPrint(reverseList(list))

	list = LinkedListFromArray([]int{1})
	LinkedListPrint(list)
	LinkedListPrint(reverseList(list))

	list = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	LinkedListPrint(list)
	LinkedListPrint(reverseList(list))

	list = LinkedListFromArray([]int{1, 2, 3, 4, 5, 6})
	LinkedListPrint(list)
	LinkedListPrint(reverseList(list))
}
