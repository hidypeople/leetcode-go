package task0450

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	tree := BSTFromArrayInts([]int{5, 3, 6, 2, 4, NULL, 7})
	newTree := deleteNode(tree, 0)
	assert.Equal(t, []int{5, 3, 6, 2, 4, NULL, 7}, BSTToArrayInts(newTree))
}

func Test_deleteNode2(t *testing.T) {
	tree := BSTFromArrayInts([]int{})
	newTree := deleteNode(tree, 0)
	assert.Equal(t, []int{}, BSTToArrayInts(newTree))
}

func Test_deleteNode3(t *testing.T) {
	tree := BSTFromArrayInts([]int{2, 1, 3})
	newTree := deleteNode(tree, 1)
	assert.Equal(t, []int{2, NULL, 3}, BSTToArrayInts(newTree))
}

func Test_deleteNode4(t *testing.T) {
	tree := BSTFromArrayInts([]int{5, 3, 6, 2, 4, NULL, 7})
	newTree := deleteNode(tree, 3)
	assert.Equal(t, []int{5, 4, 6, 2, NULL, NULL, 7}, BSTToArrayInts(newTree))
}

func Test_deleteNode5(t *testing.T) {
	tree := BSTFromArrayInts([]int{5, 3, 6, 2, 4, NULL, 7})
	newTree := deleteNode(tree, 5)
	assert.Equal(t, []int{6, 3, 7, 2, 4}, BSTToArrayInts(newTree))
}

func Test_deleteNode6(t *testing.T) {
	tree := BSTFromArrayInts([]int{50, 30, 70, NULL, 40, 60, 80})
	newTree := deleteNode(tree, 50)
	assert.Equal(t, []int{60, 30, 70, NULL, 40, NULL, 80}, BSTToArrayInts(newTree))
}
