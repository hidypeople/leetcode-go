package task0173

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTIterator_Next(t *testing.T) {
	tree := BSTFromArrayInts([]int{7, 3, 15, NULL, NULL, 9, 20})
	iterator := Constructor(tree)
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 3, iterator.Next())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 7, iterator.Next())
	assert.Equal(t, 9, iterator.Next())
	assert.Equal(t, 15, iterator.Next())
	assert.Equal(t, 20, iterator.Next())
	assert.False(t, iterator.HasNext())
}
