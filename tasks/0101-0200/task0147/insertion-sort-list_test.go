package task0147

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertionSortList(t *testing.T) {
	list := LinkedListFromArray([]int{4, 2, 1, 3})
	assert.Equal(t, []int{1, 2, 3, 4}, LinkedListToArray(insertionSortList(list)))
}

func Test_insertionSortList2(t *testing.T) {
	list := LinkedListFromArray([]int{-1, 5, 3, 4, 0})
	assert.Equal(t, []int{-1, 0, 3, 4, 5}, LinkedListToArray(insertionSortList(list)))
}
