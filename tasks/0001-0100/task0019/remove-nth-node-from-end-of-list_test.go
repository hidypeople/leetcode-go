package task0019

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd1(t *testing.T) {
	assert.Nil(t, removeNthFromEnd(nil, 1))
}

func TestRemoveNthFromEnd2(t *testing.T) {
	input := LinkedListFromArray([]int{1})
	assert.Nil(t, removeNthFromEnd(input, 1))
}

func TestRemoveNthFromEnd3(t *testing.T) {
	input := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(removeNthFromEnd(input, 2))
	assert.Equal(t, result, []int{1, 2, 3, 5})
}

func TestRemoveNthFromEnd4(t *testing.T) {
	input := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(removeNthFromEnd(input, 5))
	assert.Equal(t, result, []int{2, 3, 4, 5})
}
