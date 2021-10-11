package task0098

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	tree := BSTFromArray([]*int{I(2), I(1), I(3)})
	assert.True(t, isValidBST(tree))
}

func Test_isValidBST2(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(5),
		I(1), I(4),
		nil, nil, I(3), I(6)})
	assert.False(t, isValidBST(tree))
}

func Test_isValidBST3(t *testing.T) {
	tree := BSTFromArray([]*int{I(2), I(2), I(2)})
	assert.False(t, isValidBST(tree))
}

func Test_isValidBST4(t *testing.T) {
	tree := BSTFromArray([]*int{I(-2147483648), nil, I(2147483647)})
	assert.True(t, isValidBST(tree))
}
