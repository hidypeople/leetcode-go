package tasks

import . "leetcode/linkedList"

// Find middle node of linked list
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func Tests() {
	var list *ListNode

	list = nil
	LinkedListPrint(list)
	LinkedListPrintNode(middleNode(list))

	list = LinkedListFromArray([]int{1})
	LinkedListPrint(list)
	LinkedListPrintNode(middleNode(list))

	list = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	LinkedListPrint(list)
	LinkedListPrintNode(middleNode(list))

	list = LinkedListFromArray([]int{1, 2, 3, 4, 5, 6})
	LinkedListPrint(list)
	LinkedListPrintNode(middleNode(list))
}
