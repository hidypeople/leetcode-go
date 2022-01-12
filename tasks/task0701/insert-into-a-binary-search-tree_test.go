package task0701

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertIntoBST(t *testing.T) {
	tree := BSTFromArrayInts([]int{4, 2, 7, 1, 3})
	result := insertIntoBST(tree, 5)
	assert.Equal(t, []int{4, 2, 7, 1, 3, 5}, BSTToArrayInts(result))
}

func Test_insertIntoBST2(t *testing.T) {
	tree := BSTFromArrayInts([]int{40, 20, 60, 10, 30, 50, 70})
	result := insertIntoBST(tree, 25)
	assert.Equal(t, []int{40, 20, 60, 10, 30, 50, 70, NULL, NULL, 25}, BSTToArrayInts(result))
}

func Test_insertIntoBST3(t *testing.T) {
	tree := BSTFromArrayInts([]int{4, 2, 7, 1, 3})
	result := insertIntoBST(tree, 5)
	assert.Equal(t, []int{4, 2, 7, 1, 3, 5}, BSTToArrayInts(result))
}

func Test_insertIntoBST4(t *testing.T) {
	tree := BSTFromArrayInts([]int{55, 28, 92, 26, 43})
	result := insertIntoBST(tree, 1)
	assert.Equal(t, []int{55, 28, 92, 26, 43, NULL, NULL, 1}, BSTToArrayInts(result))
}
