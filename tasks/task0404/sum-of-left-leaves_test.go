package task0404

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	tree := BSTFromArrayInts([]int{3, 9, 20, NULL, NULL, 15, 7})
	assert.Equal(t, 24, sumOfLeftLeaves(tree))
}

func Test_sumOfLeftLeaves2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1})
	assert.Equal(t, 0, sumOfLeftLeaves(tree))
}
