package task0108

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T) {
	expected := BSTFromArray([]*int{
		I(0),
		I(-2), I(2),
		I(-3), I(-1), I(1), I(3),
	})
	result := sortedArrayToBST([]int{-3, -2, -1, 0, 1, 2, 3})
	assert.Equal(t, expected, result)
}

func TestSortedArrayToBST2(t *testing.T) {
	expected := BSTFromArray([]*int{
		I(-1),
		I(-3), I(2),
		nil, I(-2), I(1), I(3),
	})
	result := sortedArrayToBST([]int{-3, -2, -1, 1, 2, 3})
	assert.Equal(t, expected, result)
}

func TestSortedArrayToBST3(t *testing.T) {
	expected := BSTFromArray([]*int{
		I(-1),
	})
	result := sortedArrayToBST([]int{-1})
	assert.Equal(t, expected, result)
}
