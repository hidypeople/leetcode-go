package task0538

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertBST(t *testing.T) {
	tree := BSTFromArrayInts([]int{4, 1, 6, 0, 2, 5, 7, NULL, NULL, NULL, 3, NULL, NULL, NULL, 8})
	assert.Equal(t, []int{30, 36, 21, 36, 35, 26, 15, NULL, NULL, NULL, 33, NULL, NULL, NULL, 8}, BSTToArrayInts(convertBST(tree)))
}

func Test_convertBST2(t *testing.T) {
	tree := BSTFromArrayInts([]int{0, NULL, 1})
	assert.Equal(t, []int{1, NULL, 1}, BSTToArrayInts(convertBST(tree)))
}
