package task0106

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	assert.Equal(t, []int{3, 9, 20, NULL, NULL, 15, 7}, BSTToArrayInts(buildTree(inorder, postorder)))
}

func Test_buildTree2(t *testing.T) {
	inorder := []int{-1}
	postorder := []int{-1}
	assert.Equal(t, []int{-1}, BSTToArrayInts(buildTree(inorder, postorder)))
}
