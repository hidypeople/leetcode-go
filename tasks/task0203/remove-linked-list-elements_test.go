package task0203

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	var list *ListNode = nil
	assert.Nil(t, removeElements(list, 3))
}

func TestRemoveElements2(t *testing.T) {
	list := LinkedListFromArray([]int{7, 7, 7, 7})
	result := removeElements(list, 7)
	assert.Nil(t, result)
}

func TestRemoveElements3(t *testing.T) {
	list := LinkedListFromArray([]int{4, 1, 4, 2, 6, 3, 4, 4, 5, 6, 4})
	result := LinkedListToArray(removeElements(list, 4))
	assert.Equal(t, []int{1, 2, 6, 3, 5, 6}, result)
}
