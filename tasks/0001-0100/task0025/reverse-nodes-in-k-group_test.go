package task0025

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	list := LinkedListFromArray([]int{1})
	result := LinkedListToArray(reverseKGroup(list, 1))
	assert.Equal(t, []int{1}, result)
}

func Test_reverseKGroup2(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(reverseKGroup(list, 1))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

func Test_reverseKGroup3(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(reverseKGroup(list, 3))
	assert.Equal(t, []int{3, 2, 1, 4, 5}, result)
}

func Test_reverseKGroup4(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(reverseKGroup(list, 2))
	assert.Equal(t, []int{2, 1, 4, 3, 5}, result)
}

func Test_reverseKGroup5(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2})
	result := LinkedListToArray(reverseKGroup(list, 2))
	assert.Equal(t, []int{2, 1}, result)
}
