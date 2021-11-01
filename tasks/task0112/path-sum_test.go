package task0112

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	tree := BSTFromArrayInts([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, NULL, 1})
	assert.Equal(t, true, hasPathSum(tree, 22))
}

func Test_hasPathSum2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3})
	assert.Equal(t, false, hasPathSum(tree, 5))
}

func Test_hasPathSum3(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2})
	assert.Equal(t, false, hasPathSum(tree, 1))
}

func Test_hasPathSum4(t *testing.T) {
	assert.Equal(t, false, hasPathSum(nil, 0))
}
