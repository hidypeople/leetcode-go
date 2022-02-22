package task0129

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumNumbers(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3})
	assert.Equal(t, 25, sumNumbers(tree))
}

func Test_sumNumbers2(t *testing.T) {
	tree := BSTFromArrayInts([]int{4, 9, 0, 5, 1})
	assert.Equal(t, 1026, sumNumbers(tree))
}

func Test_sumNumbers3(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 0, 3, NULL, NULL, NULL, 4})
	assert.Equal(t, 1244, sumNumbers(tree))
}
