package task0124

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	assert.Equal(t, 6, maxPathSum(BSTFromArrayInts([]int{1, 2, 3})))
}

func Test_maxPathSum2(t *testing.T) {
	assert.Equal(t, 42, maxPathSum(BSTFromArrayInts([]int{-10, 9, 20, NULL, NULL, 15, 7})))
}

func Test_maxPathSum3(t *testing.T) {
	assert.Equal(t, -1, maxPathSum(BSTFromArrayInts([]int{-2, -1})))
}

func Test_maxPathSum4(t *testing.T) {
	assert.Equal(t, 2, maxPathSum(BSTFromArrayInts([]int{2, -1})))
}

func Test_maxPathSum5(t *testing.T) {
	assert.Equal(t, -2, maxPathSum(BSTFromArrayInts([]int{-2, NULL, -3})))
}

func Test_maxPathSum6(t *testing.T) {
	assert.Equal(t, 12, maxPathSum(BSTFromArrayInts([]int{-1, NULL, 9, -6, 3, NULL, NULL, NULL, -2})))
}
