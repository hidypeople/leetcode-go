package task0199

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rightSideView(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3, NULL, 5, NULL, 4})
	assert.Equal(t, []int{1, 3, 4}, rightSideView(tree))
}

func Test_rightSideView2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, NULL, 3})
	assert.Equal(t, []int{1, 3}, rightSideView(tree))
}

func Test_rightSideView3(t *testing.T) {
	assert.Equal(t, []int{}, rightSideView(nil))
}
