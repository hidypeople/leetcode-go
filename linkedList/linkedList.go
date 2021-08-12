package linkedList

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode for int values
type ListNode struct {
	Val  int
	Next *ListNode
}

// Convert array into linked list
func LinkedListFromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	root := &ListNode{arr[0], nil}
	curr := root
	for i := 1; i < len(arr); i++ {
		new := &ListNode{arr[i], nil}
		curr.Next = new
		curr = new
	}
	return root
}

// Convert linked list to array
func LinkedListToArray(list *ListNode) []int {
	visited := make(map[*ListNode]bool)
	result := []int{}
	curr := list
	for curr != nil {
		visited[curr] = true
		result = append(result, curr.Val)
		curr = curr.Next
		if curr != nil && visited[curr] {
			break
		}
	}
	return result
}

// Print single node
func LinkedListPrintNode(node *ListNode) {
	if node == nil {
		fmt.Printf("<nil>\n")
		return
	}
	fmt.Printf("%v\n", node.Val)
}

// Print linked list
func LinkedListPrint(list *ListNode) {
	if list == nil {
		fmt.Printf("<nil>\n")
		return
	}
	builder := new(strings.Builder)
	visited := make(map[*ListNode]bool)
	curr := list
	for curr != nil {
		visited[curr] = true
		builder.WriteString(strconv.Itoa(curr.Val))
		curr = curr.Next
		if curr != nil {
			if visited[curr] {
				builder.WriteString(" (-> ")
				builder.WriteString(strconv.Itoa(curr.Val))
				builder.WriteRune(')')
				break
			}
			builder.WriteRune(' ')
		}
	}
	fmt.Printf("%v\n", builder.String())
}
