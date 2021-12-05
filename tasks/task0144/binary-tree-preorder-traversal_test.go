package task0144

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(BSTFromArrayInts([]int{1, NULL, 2, 3})))
}

func Test_preorderTraversal2(t *testing.T) {
	assert.Equal(t, []int{}, preorderTraversal(nil))
}

func Test_preorderTraversal3(t *testing.T) {
	assert.Equal(t, []int{1}, preorderTraversal(BSTFromArrayInts([]int{1})))
}

func Test_preorderTraversal4(t *testing.T) {
	assert.Equal(t, []int{1, 2}, preorderTraversal(BSTFromArrayInts([]int{1, 2})))
}

func Test_preorderTraversal5(t *testing.T) {
	assert.Equal(t, []int{1, 2}, preorderTraversal(BSTFromArrayInts([]int{1, NULL, 2})))
}
