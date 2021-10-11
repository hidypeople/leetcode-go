package task0094

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		nil, I(2),
		nil, nil, I(3), nil,
	})
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(tree))
}

func Test_inorderTraversal2(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		I(2),
	})
	assert.Equal(t, []int{2, 1}, inorderTraversal(tree))
}

func Test_inorderTraversal3(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		nil, I(2),
	})
	assert.Equal(t, []int{1, 2}, inorderTraversal(tree))
}
