package task0102

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	tree := BSTFromArrayInts([]int{3, 9, 20, NULL, NULL, 15, 7})
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(tree))
}

func Test_levelOrder2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1})
	assert.Equal(t, [][]int{{1}}, levelOrder(tree))
}

func Test_levelOrder3(t *testing.T) {
	tree := BSTFromArrayInts([]int{})
	assert.Equal(t, [][]int{}, levelOrder(tree))
}
