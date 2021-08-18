package task0876

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	assert.Nil(t, middleNode(nil))
}

func TestMiddleNode2(t *testing.T) {
	l := LinkedListFromArray([]int{1})
	result := middleNode(l)
	assert.Equal(t, result.Val, 1)
}

func TestMiddleNode3(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := middleNode(l)
	assert.Equal(t, result.Val, 3)
}

func TestMiddleNode4(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6})
	result := middleNode(l)
	assert.Equal(t, result.Val, 4)
}
