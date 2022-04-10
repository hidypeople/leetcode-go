package task1305

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getAllElements(t *testing.T) {
	tree1 := BSTFromArrayInts([]int{2, 1, 4})
	tree2 := BSTFromArrayInts([]int{1, 0, 3})
	assert.Equal(t, []int{0, 1, 1, 2, 3, 4}, getAllElements(tree1, tree2))
}

func Test_getAllElements2(t *testing.T) {
	tree1 := BSTFromArrayInts([]int{1, NULL, 8})
	tree2 := BSTFromArrayInts([]int{8, 1})
	assert.Equal(t, []int{1, 1, 8, 8}, getAllElements(tree1, tree2))
}
