package task0109

import (
	. "leetcode/binaryTree"
	"leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedListToBST(t *testing.T) {
	list := BSTFromArray([]*int{I(0), I(-3), I(9), I(-10), nil, I(5)})
	BSTPrint(list)
	result := SortedListToBST(linkedList.LinkedListFromArray([]int{-10, -3, 0, 5, 9}))
	assert.Equal(t, list, result)
}
