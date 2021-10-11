package task0092

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	assert.Nil(t, reverseBetween(nil, 1, 2))
}

func TestReverseBetween2(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7})
	result := LinkedListToArray(reverseBetween(list, 2, 4))
	assert.Equal(t, []int{1, 4, 3, 2, 5, 6, 7}, result)
}

func TestReverseBetween3(t *testing.T) {
	list := LinkedListFromArray([]int{4})
	result := LinkedListToArray(reverseBetween(list, 1, 1))
	assert.Equal(t, []int{4}, result)
}

func TestReverseBetween4(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7})
	result := LinkedListToArray(reverseBetween(list, 1, 7))
	assert.Equal(t, []int{7, 6, 5, 4, 3, 2, 1}, result)
}

func TestReverseBetween5(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7})
	result := LinkedListToArray(reverseBetween(list, 1, 1))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, result)
}

func TestReverseBetween6(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7})
	result := LinkedListToArray(reverseBetween(list, 7, 7))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, result)
}
