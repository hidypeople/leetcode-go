package task0104

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	tree := BSTFromArray([]*int{I(3), I(9), I(20), nil, nil, I(15), I(7)})
	assert.Equal(t, 3, maxDepth(tree))
}

func Test_maxDepth2(t *testing.T) {
	tree := BSTFromArray([]*int{I(1), nil, I(2)})
	assert.Equal(t, 2, maxDepth(tree))
}

func Test_maxDepth3(t *testing.T) {
	assert.Equal(t, 0, maxDepth(nil))
}

func Test_maxDepth4(t *testing.T) {
	tree := BSTFromArray([]*int{I(1)})
	assert.Equal(t, 1, maxDepth(tree))
}
