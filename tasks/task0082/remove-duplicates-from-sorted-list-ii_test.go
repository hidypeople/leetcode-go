package task0082

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Nil(t, DeleteDuplicates(nil))
}

func TestDeleteDuplicates2(t *testing.T) {
	list := LinkedListFromArray([]int{1, 1, 1})
	assert.Nil(t, DeleteDuplicates(list))
}

func TestDeleteDuplicates3(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 3, 4, 4, 5})
	result := LinkedListToArray(DeleteDuplicates(list))
	assert.Equal(t, []int{1, 2, 5}, result)
}

func TestDeleteDuplicates4(t *testing.T) {
	list := LinkedListFromArray([]int{-1, 1, 1, 1, 2, 3})
	result := LinkedListToArray(DeleteDuplicates(list))
	assert.Equal(t, []int{-1, 2, 3}, result)
}

func TestDeleteDuplicates5(t *testing.T) {
	list := LinkedListFromArray([]int{1, 1, 1, 1, 2, 2, 3, 4, 4})
	result := LinkedListToArray(DeleteDuplicates(list))
	assert.Equal(t, []int{3}, result)
}

func TestDeleteDuplicates6(t *testing.T) {
	list := LinkedListFromArray([]int{1, 1, 1, 1, 2, 2, 3, 3, 4, 4})
	assert.Nil(t, DeleteDuplicates(list))
}
