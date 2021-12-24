package task0148

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {
	list := LinkedListFromArray([]int{4, 2, 1, 3})
	sortedList := LinkedListToArray(sortList(list))
	assert.Equal(t, []int{1, 2, 3, 4}, sortedList)
}

func Test_sortList2(t *testing.T) {
	list := LinkedListFromArray([]int{-1, 5, 3, 4, 0})
	sortedList := LinkedListToArray(sortList(list))
	assert.Equal(t, []int{-1, 0, 3, 4, 5}, sortedList)
}

func Test_sortList3(t *testing.T) {
	list := LinkedListFromArray([]int{})
	sortedList := LinkedListToArray(sortList(list))
	assert.Equal(t, []int{}, sortedList)
}
