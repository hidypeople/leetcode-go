package task0103

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tree := BSTFromArrayInts([]int{3, 9, 20, NULL, NULL, 15, 7})
	assert.Equal(t, [][]int{{3}, {20, 9}, {15, 7}}, zigzagLevelOrder(tree))
}

func Test_zigzagLevelOrder2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1})
	assert.Equal(t, [][]int{{1}}, zigzagLevelOrder(tree))
}

func Test_zigzagLevelOrder3(t *testing.T) {
	tree := BSTFromArrayInts([]int{})
	assert.Equal(t, [][]int{}, zigzagLevelOrder(tree))
}

func Test_zigzagLevelOrder4(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3, 4, NULL, NULL, 5})
	assert.Equal(t, [][]int{{1}, {3, 2}, {4, 5}}, zigzagLevelOrder(tree))
}
