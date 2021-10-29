package task0107

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrderBottom(t *testing.T) {
	tree := BSTFromArrayInts([]int{3, 9, 20, NULL, NULL, 15, 7})
	assert.Equal(t, [][]int{{15, 7}, {9, 20}, {3}}, levelOrderBottom(tree))
}

func Test_levelOrderBottom2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1})
	assert.Equal(t, [][]int{{1}}, levelOrderBottom(tree))
}

func Test_levelOrderBottom3(t *testing.T) {
	tree := BSTFromArrayInts([]int{})
	assert.Equal(t, [][]int{}, levelOrderBottom(tree))
}
