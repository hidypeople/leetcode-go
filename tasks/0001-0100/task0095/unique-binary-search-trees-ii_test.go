package task0095

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateTrees(t *testing.T) {
	tree := generateTrees(1)
	assert.Equal(t, 1, len(tree))
	assert.Equal(t, 1, tree[0].Val)
	assert.Nil(t, tree[0].Left)
	assert.Nil(t, tree[0].Right)
}

func Test_generateTrees2(t *testing.T) {
	// [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
	tree := generateTrees(3)
	assert.Equal(t, 5, len(tree))
}

func Test_generateTrees3(t *testing.T) {
	// [[1, null, 2],[2, 1]]
	tree := generateTrees(2)
	assert.Equal(t, 2, len(tree))
}

func Test_generateTrees4(t *testing.T) {
	tree := generateTrees(7)
	assert.Equal(t, 429, len(tree))
}

func Test_generateTrees5(t *testing.T) {
	tree := generateTrees(8)
	assert.Equal(t, 1430, len(tree))
}
