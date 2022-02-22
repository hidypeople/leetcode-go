package task0105

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	assert.Equal(t, []int{3, 9, 20, NULL, NULL, 15, 7}, BSTToArrayInts(result))
}

func Test_buildTree2(t *testing.T) {
	preorder := []int{-1}
	inorder := []int{-1}
	result := buildTree(preorder, inorder)
	assert.Equal(t, []int{-1}, BSTToArrayInts(result))
}
