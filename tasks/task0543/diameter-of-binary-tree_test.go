package task0543

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tree := BSTFromArray([]*int{I(1), I(2), I(3), I(4), I(5)})
	assert.Equal(t, 3, diameterOfBinaryTree(tree))
}

func Test_diameterOfBinaryTree2(t *testing.T) {
	tree := BSTFromArray([]*int{I(1), I(2)})
	assert.Equal(t, 1, diameterOfBinaryTree(tree))
}

func Test_diameterOfBinaryTree3(t *testing.T) {
	tree := BSTFromArray([]*int{I(1)})
	assert.Equal(t, 0, diameterOfBinaryTree(tree))
}
