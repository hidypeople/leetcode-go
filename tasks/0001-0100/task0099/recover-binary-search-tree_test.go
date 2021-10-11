package task0099

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_recoverTree(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 3, NULL, NULL, 2})
	recoverTree(tree)
	expected := BSTFromArrayInts([]int{3, 1, NULL, NULL, 2})
	assert.Equal(t, BSTToString(expected), BSTToString(tree))
}

func Test_recoverTree2(t *testing.T) {
	tree := BSTFromArrayInts([]int{3, 1, 4, NULL, NULL, 2})
	recoverTree(tree)
	expected := BSTFromArrayInts([]int{2, 1, 4, NULL, NULL, 3})
	assert.Equal(t, BSTToString(expected), BSTToString(tree))
}
