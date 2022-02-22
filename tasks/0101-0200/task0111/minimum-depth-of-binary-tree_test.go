package task0111

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDepth(t *testing.T) {
	root := BSTFromArrayInts([]int{3, 9, 20, NULL, NULL, 15, 7})
	assert.Equal(t, 2, minDepth(root))
}

func Test_minDepth2(t *testing.T) {
	root := BSTFromArrayInts([]int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6})
	assert.Equal(t, 5, minDepth(root))
}
