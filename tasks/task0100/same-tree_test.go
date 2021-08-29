package task0100

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSameTree(t *testing.T) {
	p := BSTFromArray([]*int{
		I(1),
		I(2), I(3),
	})
	q := BSTFromArray([]*int{
		I(1),
		I(2), I(3),
	})
	assert.True(t, isSameTree(p, q))
}

func Test_isSameTree2(t *testing.T) {
	p := BSTFromArray([]*int{
		I(1),
		I(2), I(3),
		I(4),
	})
	q := BSTFromArray([]*int{
		I(1),
		I(2), I(3),
	})
	assert.False(t, isSameTree(p, q))
}

func Test_isSameTree3(t *testing.T) {
	p := BSTFromArray([]*int{
		I(1),
		I(2), nil,
	})
	q := BSTFromArray([]*int{
		I(1),
		nil, I(2),
	})
	assert.False(t, isSameTree(p, q))
}
