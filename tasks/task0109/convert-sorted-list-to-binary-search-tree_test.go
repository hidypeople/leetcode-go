package task0109

import (
	. "leetcode/binaryTree"
	"leetcode/linkedList"
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedListToBST(t *testing.T) {
	list := BSTFromArray([]*int{I(0), I(-10), I(10), I(-50), I(-5), I(5), I(50)})
	result := SortedListToBST(linkedList.LinkedListFromArray([]int{-50, -10, -5, 0, 5, 10, 50}))
	assert.Equal(t, list, result)
}

func TestSortedListToBST2(t *testing.T) {
	list := BSTFromArray([]*int{I(5), I(-10), I(15), I(-100), I(0), I(10), nil})
	result := SortedListToBST(linkedList.LinkedListFromArray([]int{-100, -10, 0, 5, 10, 15}))
	assert.Equal(t, list, result)
}

func TestMiddleNodePrev(t *testing.T) {
	list0 := LinkedListFromArray([]int{1})
	assert.Equal(t, 1, MiddleNodePrev(list0).Val)

	list := LinkedListFromArray([]int{-3, -2, -1, 0, 1, 2, 3})
	assert.Equal(t, -1, MiddleNodePrev(list).Val)

	list2 := LinkedListFromArray([]int{-3, -2, -1, 1, 2, 3})
	assert.Equal(t, -1, MiddleNodePrev(list2).Val)

	list3 := LinkedListFromArray([]int{-3, -2})
	assert.Equal(t, -3, MiddleNodePrev(list3).Val)

	list4 := LinkedListFromArray([]int{-3, -2, -1})
	assert.Equal(t, -3, MiddleNodePrev(list4).Val)
}
