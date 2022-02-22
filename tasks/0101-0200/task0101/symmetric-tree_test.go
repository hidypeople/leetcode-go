package task0101

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		I(2), I(2),
		I(3), I(4), I(4), I(3),
	})
	assert.True(t, isSymmetric(tree))
}

func Test_isSymmetric2(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		I(2), I(2),
		I(3), I(1), I(4), I(3),
	})
	assert.False(t, isSymmetric(tree))
}

func Test_isSymmetric3(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		I(1), I(2),
		I(3), I(1), I(4), I(3),
	})
	assert.False(t, isSymmetric(tree))
}

func Test_isSymmetric4(t *testing.T) {
	assert.True(t, isSymmetric(nil))
}

func Test_isSymmetric5(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		I(1), I(2),
		nil, nil, I(4), I(3),
	})
	assert.False(t, isSymmetric(tree))
}

func Test_isSymmetric6(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(1),
		I(1), nil,
		I(4), I(3),
	})
	assert.False(t, isSymmetric(tree))
}
