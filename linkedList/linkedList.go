package linkedList

import "fmt"

// ListNode for int values
type ListNodeInt struct {
	Val  int
	Next *ListNodeInt
}

// Convert array into linked list
func LinkedListFromArray(arr []int) *ListNodeInt {
	if len(arr) == 0 {
		return nil
	}
	root := &ListNodeInt{arr[0], nil}
	curr := root
	for i := 1; i < len(arr); i++ {
		new := &ListNodeInt{arr[i], nil}
		curr.Next = new
		curr = new
	}
	return root
}

// Print linked list
func LinkedListPrint(list *ListNodeInt) {
	curr := list
	for curr != nil {
		fmt.Printf("%v", curr.Val)
		curr = curr.Next
		if curr != nil {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
