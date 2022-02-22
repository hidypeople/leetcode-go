package task0145

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, postorderTraversal(BSTFromArrayInts([]int{1, NULL, 2, 3})))
}

func Test_postorderTraversal2(t *testing.T) {
	assert.Equal(t, []int{}, postorderTraversal(nil))
}

func Test_postorderTraversal3(t *testing.T) {
	assert.Equal(t, []int{1}, postorderTraversal(BSTFromArrayInts([]int{1})))
}

func Test_postorderTraversal4(t *testing.T) {
	assert.Equal(t, []int{2, 1}, postorderTraversal(BSTFromArrayInts([]int{1, 2})))
	assert.Equal(t, []int{2, 1}, postorderTraversal(BSTFromArrayInts([]int{1, NULL, 2})))
}
