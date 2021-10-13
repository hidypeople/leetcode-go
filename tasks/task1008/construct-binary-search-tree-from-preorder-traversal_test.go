package task1008

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bstFromPreorder(t *testing.T) {
	tree := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	assert.Equal(t, []int{8, 5, 10, 1, 7, NULL, 12}, BSTToArrayInts(tree))
}

func Test_bstFromPreorder2(t *testing.T) {
	tree := bstFromPreorder([]int{1, 3})
	assert.Equal(t, []int{1, NULL, 3}, BSTToArrayInts(tree))
}

func Test_bstFromPreorder3(t *testing.T) {
	tree := bstFromPreorder([]int{19, 4, 8, 11})
	assert.Equal(t, []int{19, 4, NULL, NULL, 8, NULL, 11}, BSTToArrayInts(tree))
}

func Test_bstFromPreorder4(t *testing.T) {
	tree := bstFromPreorder([]int{19, 4, 1, 100})
	assert.Equal(t, []int{19, 4, 100, 1}, BSTToArrayInts(tree))
}

func Test_bstFromPreorder5(t *testing.T) {
	tree := bstFromPreorder([]int{3, 2, 1})
	assert.Equal(t, []int{3, 2, NULL, 1}, BSTToArrayInts(tree))
}

func Test_bstFromPreorder6(t *testing.T) {
	tree := bstFromPreorder([]int{16, 15, 10, 5, 11})
	assert.Equal(t, []int{16, 15, NULL, 10, NULL, 5, 11}, BSTToArrayInts(tree))
}

func Test_bstFromPreorder7(t *testing.T) {
	tree := bstFromPreorder([]int{19, 2, 1, 17, 11, 12})
	assert.Equal(t, []int{19, 2, NULL, 1, 17, NULL, NULL, 11, NULL, NULL, 12}, BSTToArrayInts(tree))
}
