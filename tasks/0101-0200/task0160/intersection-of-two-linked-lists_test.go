package task0160

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	listA := LinkedListFromArray([]int{4, 1, 8, 4, 5})
	listB := LinkedListFromArray([]int{5, 6, 1, 8, 4, 5})
	listA.Next.Next = listB.Next.Next.Next
	assert.Equal(t, []int{8, 4, 5}, LinkedListToArray(getIntersectionNode(listA, listB)))
}

func Test_getIntersectionNode2(t *testing.T) {
	listA := LinkedListFromArray([]int{1, 9, 1, 2, 4})
	listB := LinkedListFromArray([]int{3, 2, 4})
	listA.Next.Next.Next = listB.Next
	assert.Equal(t, []int{2, 4}, LinkedListToArray(getIntersectionNode(listA, listB)))
}

func Test_getIntersectionNode3(t *testing.T) {
	listA := LinkedListFromArray([]int{2, 6, 4})
	listB := LinkedListFromArray([]int{1, 5})
	assert.Nil(t, getIntersectionNode(listA, listB))
}
