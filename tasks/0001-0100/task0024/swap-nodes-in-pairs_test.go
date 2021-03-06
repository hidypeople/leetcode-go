package task0024

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	assert.Nil(t, swapPairs(nil))
}

func TestSwapPairs2(t *testing.T) {
	l := LinkedListFromArray([]int{1})
	assert.Equal(t, l, swapPairs(l))
}

func TestSwapPairs3(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4})
	result := LinkedListToArray(swapPairs(l))
	assert.Equal(t, []int{2, 1, 4, 3}, result)
}

func TestSwapPairs4(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(swapPairs(l))
	assert.Equal(t, []int{2, 1, 4, 3, 5}, result)
}
