package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTFromArray(t *testing.T) {
	result := BSTFromArray(nil)
	assert.Nil(t, result)
}

func TestBSTFromArray2(t *testing.T) {
	result := BSTFromArray([]*int{
		I(1),
	})
	assert.Equal(t, result.Val, 1)
	assert.Nil(t, result.Left)
	assert.Nil(t, result.Right)
}

func TestBSTFromArray3(t *testing.T) {
	result := BSTFromArray([]*int{
		I(1), I(2), I(3),
	})
	assert.Equal(t, result.Val, 1)
	assert.Equal(t, result.Left.Val, 2)
	assert.Nil(t, result.Left.Left)
	assert.Nil(t, result.Left.Right)
	assert.Equal(t, result.Right.Val, 3)
	assert.Nil(t, result.Right.Left)
	assert.Nil(t, result.Right.Right)
}

func TestBSTFromArray4(t *testing.T) {
	result := BSTFromArray([]*int{
		I(1), I(2), I(3), nil, I(4), I(5),
	})
	BSTPrint(result)
	assert.Equal(t, 1, result.Val)

	// Left
	assert.Equal(t, 2, result.Left.Val)
	assert.Equal(t, 4, result.Left.Right.Val)
	assert.Nil(t, result.Left.Left)

	// Right
	assert.Equal(t, 3, result.Right.Val)
	assert.Equal(t, 5, result.Right.Left.Val)
	assert.Nil(t, result.Right.Right)
}

func TestBSTFromArray5(t *testing.T) {
	result := BSTFromArray([]*int{
		I(1), nil, I(3), nil, nil, I(5), I(6),
	})
	BSTPrint(result)
	assert.Equal(t, 1, result.Val)

	// Left
	assert.Nil(t, result.Left)

	// Right
	assert.Equal(t, 3, result.Right.Val)
	assert.Equal(t, 5, result.Right.Left.Val)
	assert.Equal(t, 6, result.Right.Right.Val)
}

func TestBSTFromArray6(t *testing.T) {
	result := BSTFromArray([]*int{
		I(0),
		I(-1), I(1),
		nil, I(-2), I(2), I(3),
		nil, nil, nil, nil, nil, I(4),
	})
	assert.Equal(t, 0, result.Val)

	// Left
	assert.Equal(t, result.Left.Val, -1)
	assert.Nil(t, result.Left.Left)
	assert.Equal(t, result.Left.Right.Val, -2)

	// Right
	assert.Equal(t, 1, result.Right.Val)
	assert.Equal(t, 2, result.Right.Left.Val)
	assert.Equal(t, 4, result.Right.Left.Right.Val)
	assert.Equal(t, 3, result.Right.Right.Val)
}

func Test_BSTFromArrayInts(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3})
	assert.Equal(t, 1, tree.Val)
	assert.Equal(t, 2, tree.Left.Val)
	assert.Equal(t, 3, tree.Right.Val)
}

func Test_BSTFromArrayInts2(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, 3, NULL, 100})
	assert.Nil(t, tree.Left.Left)
	assert.NotNil(t, tree.Left.Right)
	assert.Equal(t, 100, tree.Left.Right.Val)
}

func Test_BSTFromArrayInts3(t *testing.T) {
	tree := BSTFromArrayInts([]int{1, 2, NULL, NULL, 200, 2000, 3000})
	assert.Equal(t, 200, tree.Left.Right.Val)
	assert.Equal(t, 2000, tree.Left.Right.Left.Val)
	assert.Equal(t, 3000, tree.Left.Right.Right.Val)
}

func Test_BSTToArrayInts(t *testing.T) {
	var input []int

	input = []int{1, 2, 3}
	assert.Equal(t, input, BSTToArrayInts(BSTFromArrayInts(input)))

	input = []int{1, 2, 3, NULL, 100}
	assert.Equal(t, input, BSTToArrayInts(BSTFromArrayInts(input)))

	input = []int{1, 2, NULL, NULL, 200, 2000, 3000}
	assert.Equal(t, input, BSTToArrayInts(BSTFromArrayInts(input)))
}
