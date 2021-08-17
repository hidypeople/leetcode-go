package augustleetcodingchallenge2021

import (
	. "leetcode/binaryTree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodNodes(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(3),
		I(3), nil,
		I(4), I(2),
	})
	assert.Equal(t, 3, GoodNodes(tree))
}

func TestGoodNodes2(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(3),
	})
	assert.Equal(t, 1, GoodNodes(tree))
}

func TestGoodNodes3(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(3),
		I(1), I(4),
		I(3), nil, I(1), I(5),
	})
	assert.Equal(t, 4, GoodNodes(tree))
}

func TestGoodNodes4(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(2),
		nil, I(4),
		nil, nil, I(10), I(8),
		nil, nil, nil, nil, nil, nil, I(4), nil,
	})
	assert.Equal(t, 4, GoodNodes(tree))
}

func TestGoodNodes_NoRecursion(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(3),
		I(3), nil,
		I(4), I(2),
	})
	assert.Equal(t, 3, GoodNodes_NoRecursion(tree))
}

func TestGoodNodes_NoRecursion2(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(3),
	})
	assert.Equal(t, 1, GoodNodes_NoRecursion(tree))
}

func TestGoodNodes_NoRecursion3(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(3),
		I(1), I(4),
		I(3), nil, I(1), I(5),
	})
	assert.Equal(t, 4, GoodNodes_NoRecursion(tree))
}

func TestGoodNodes_NoRecursion4(t *testing.T) {
	tree := BSTFromArray([]*int{
		I(2),
		nil, I(4),
		nil, nil, I(10), I(8),
		nil, nil, nil, nil, nil, nil, I(4), nil,
	})
	assert.Equal(t, 4, GoodNodes_NoRecursion(tree))
}
