package task0110

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	tree := BSTFromArray([]*int{I(1), I(2), I(2), I(3), I(3), nil, nil, I(4), I(4)})
	assert.False(t, isBalanced(tree))
}

func Test_isBalanced2(t *testing.T) {
	assert.True(t, isBalanced(nil))
}

func Test_isBalanced3(t *testing.T) {
	tree := BSTFromArray([]*int{I(1), I(2), I(2), I(3), I(3)})
	assert.True(t, isBalanced(tree))
}
