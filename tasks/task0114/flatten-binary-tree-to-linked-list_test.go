package task0114

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_flatten(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 5, 3, 4, NULL, 6})
	flatten(tree)
	assert.Equal(t, []int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5, NULL, 6}, BSTToArrayInts(tree))
}

func Test_flatten2(t *testing.T) {
	tree := BSTFromArrayInts([]int{})
	flatten(tree)
	assert.Equal(t, []int{}, BSTToArrayInts(tree))
}

func Test_flatten3(t *testing.T) {
	tree := BSTFromArrayInts([]int{1})
	flatten(tree)
	assert.Equal(t, []int{1}, BSTToArrayInts(tree))
}

func Test_flatten4(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, NULL, 3})
	flatten(tree)
	assert.Equal(t, []int{1, NULL, 2, NULL, 3}, BSTToArrayInts(tree))
}
