package task0938

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rangeSumBST(t *testing.T) {
	tree := BSTFromArrayInts([]int{10, 5, 15, 3, 7, NULL, 18})
	assert.Equal(t, 32, rangeSumBST(tree, 7, 15))
}

func Test_rangeSumBST2(t *testing.T) {
	tree := BSTFromArrayInts([]int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6})
	assert.Equal(t, 23, rangeSumBST(tree, 6, 10))
}
