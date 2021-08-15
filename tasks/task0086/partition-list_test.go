package task0086

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	assert.Nil(t, Partition(nil, 10))
}

func TestPartition2(t *testing.T) {
	list := LinkedListFromArray([]int{1, 4, 3, 2, 5, 2})
	result := LinkedListToArray(Partition(list, 3))
	assert.Equal(t, []int{1, 2, 2, 4, 3, 5}, result)
}

func TestPartition3(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
	result := LinkedListToArray(Partition(list, 1))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, result)
}

func TestPartition4(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
	result := LinkedListToArray(Partition(list, 8))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, result)
}

func TestPartition5(t *testing.T) {
	list := LinkedListFromArray([]int{2, 1})
	result := LinkedListToArray(Partition(list, 2))
	assert.Equal(t, []int{1, 2}, result)
}

func TestPartition6(t *testing.T) {
	list := LinkedListFromArray([]int{2})
	result := LinkedListToArray(Partition(list, 2))
	assert.Equal(t, []int{2}, result)
}
