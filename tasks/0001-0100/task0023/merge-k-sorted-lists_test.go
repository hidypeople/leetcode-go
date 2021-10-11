package task0023

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	l1 := LinkedListFromArray([]int{1, 2, 3})
	l2 := LinkedListFromArray([]int{1, 100, 200})
	l3 := LinkedListFromArray([]int{2, 3, 4, 5})
	result := LinkedListToArray(mergeKLists([]*ListNode{l1, l2, l3}))
	assert.Equal(t, result, []int{1, 1, 2, 2, 3, 3, 4, 5, 100, 200})
}
