package task0138

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const NULL int = 1<<63 - 1

func LinkedListRandomToArray(head *Node) [][]int {
	result := [][]int{}
	curr := head
	nodes := map[*Node]int{}
	i := 0
	for curr != nil {
		nodes[curr] = i
		result = append(result, []int{curr.Val, NULL})
		curr = curr.Next
		i++
	}

	// Set up random
	curr = head
	for i = 0; i < len(result); i++ {
		if curr.Random != nil {
			result[i][1] = nodes[curr.Random]
		}
		curr = curr.Next
	}

	return result
}

func LinkedListRandomFromArray(arr [][]int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{
		Val: arr[0][0],
	}
	nodes := map[int]*Node{0: head}

	// Build the linked list
	curr := head
	for i := 1; i < len(arr); i++ {
		next := &Node{Val: arr[i][0]}
		nodes[i] = next
		curr.Next = next
		curr = next
	}

	// set up randoms
	curr = head
	for i := 0; i < len(arr); i++ {
		if arr[i][1] != NULL {
			if randNode, ok := nodes[arr[i][1]]; ok {
				curr.Random = randNode
			}
		}
		curr = curr.Next
	}
	return head
}

func Test_copyRandomList(t *testing.T) {
	assert.Nil(t, copyRandomList(nil))
}

func Test_copyRandomList2(t *testing.T) {
	randomList := LinkedListRandomFromArray(
		[][]int{{7, NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	copy := copyRandomList(randomList)
	assert.Equal(t,
		[][]int{{7, NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		LinkedListRandomToArray(copy))
}

func Test_copyRandomList3(t *testing.T) {
	randomList := LinkedListRandomFromArray(
		[][]int{{1, 1}, {2, 1}})
	assert.Equal(t,
		[][]int{{1, 1}, {2, 1}},
		LinkedListRandomToArray(copyRandomList(randomList)))
}

func Test_copyRandomList4(t *testing.T) {
	randomList := LinkedListRandomFromArray(
		[][]int{{3, NULL}, {3, 0}, {3, NULL}})
	assert.Equal(t,
		[][]int{{3, NULL}, {3, 0}, {3, NULL}},
		LinkedListRandomToArray(copyRandomList(randomList)))
}
