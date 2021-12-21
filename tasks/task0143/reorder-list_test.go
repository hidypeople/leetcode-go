package task0143

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorderList(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4})
	reorderList(list)
	result := LinkedListToArray(list)
	assert.Equal(t, []int{1, 4, 2, 3}, result)
}

func Test_reorderList2(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	reorderList(list)
	result := LinkedListToArray(list)
	assert.Equal(t, []int{1, 5, 2, 4, 3}, result)
}

func Test_reorderList3(t *testing.T) {
	list := LinkedListFromArray([]int{1})
	reorderList(list)
	result := LinkedListToArray(list)
	assert.Equal(t, []int{1}, result)
}
