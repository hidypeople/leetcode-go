package task0993

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCousins(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3, 4})
	assert.Equal(t, false, isCousins(tree, 4, 3))
}

func Test_isCousins2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3, NULL, 4, NULL, 5})
	assert.Equal(t, true, isCousins(tree, 5, 4))
}

func Test_isCousins3(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3, NULL, 4})
	assert.Equal(t, false, isCousins(tree, 2, 3))
}

func Test_isCousins4(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, NULL, 3, 4, NULL, NULL, 5})
	assert.Equal(t, false, isCousins(tree, 2, 4))
}
