package task0988

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestFromLeaf(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(25),
		I(1), I(3),
		I(1), I(3), I(0), I(2),
	})
	assert.Equal(t, "adz", smallestFromLeaf(tree))
}

func Test_smallestFromLeaf2(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(2),
		I(2), I(1),
		nil, I(1), I(0), nil,
		nil, nil, I(0),
	})
	assert.Equal(t, "abc", smallestFromLeaf(tree))
}

func Test_smallestFromLeaf3(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(0),
		I(1), I(2),
		I(3), I(4), I(3), I(4),
	})
	assert.Equal(t, "dba", smallestFromLeaf(tree))
}

func Test_smallestFromLeaf4(t *testing.T) {
	assert.Equal(t, "", smallestFromLeaf(nil))
}

func Test_smallestFromLeaf5(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(4),
		I(0), I(1),
		I(1), I(1), nil, nil,
	})
	assert.Equal(t, "bae", smallestFromLeaf(tree))
}

func Test_smallestFromLeaf6(t *testing.T) {
	// [25,1,null,0,0,1,null,null,null,0]
	tree := BSTFromArray([]*int{
		I(25),
		I(1), nil,
		I(0), I(0), nil, nil,
		I(1), nil, nil, nil, nil, nil, nil, nil,
		I(0),
	})
	assert.Equal(t, "ababz", smallestFromLeaf(tree))
}

func Test_smallestFromLeaf7(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 6},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 5},
					},
				},
			},
		},
	}
	assert.Equal(t, "fedcba", smallestFromLeaf(root))
}

func Test_getMinPath(t *testing.T) {
	var a *[]int = &[]int{1, 2, 3, 4}
	var b *[]int = &[]int{1, 2, 3}

	assert.Equal(t, b, getMinPath(a, b))

	a = &[]int{2, 2, 3, 4}
	b = &[]int{1, 2, 3, 4}
	assert.Equal(t, b, getMinPath(a, b))

	a = &[]int{1, 2, 3, 4}
	b = &[]int{1, 2, 3, 4}
	assert.Equal(t, a, getMinPath(a, b))

	a = &[]int{1}
	b = &[]int{3, 2, 1}
	assert.Equal(t, a, getMinPath(a, b))

	a = &[]int{}
	b = &[]int{3, 2, 1}
	assert.Equal(t, a, getMinPath(a, b))

	a = &[]int{}
	b = &[]int{}
	assert.Equal(t, a, getMinPath(a, b))

	b = &[]int{3, 2, 1}
	assert.Equal(t, b, getMinPath(nil, b))

	a = &[]int{0, 1, 2, 3, 4, 5}
	b = &[]int{0, 1, 2, 6}
	assert.Equal(t, a, getMinPath(a, b))
}
