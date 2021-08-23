package task0226

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {
	result := invertTree(BSTFromArray([]*int{
		I(4),
		I(2), I(7),
		I(1), I(3), I(6), I(9),
	}))
	expected := BSTFromArray([]*int{
		I(4),
		I(7), I(2),
		I(9), I(6), I(3), I(1),
	})
	assert.Equal(t, expected, result)
}
