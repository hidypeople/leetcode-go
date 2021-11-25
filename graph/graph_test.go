package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphFromEdgeArray(t *testing.T) {
	var node *Node

	node = GraphFromEdgeArray([][]int{})
	assert.Nil(t, node)

	node = GraphFromEdgeArray([][]int{{1, 2}})
	assert.Equal(t, 1, node.Val)
	assert.Equal(t, 1, len(node.Neighbors))
	assert.Equal(t, 2, node.Neighbors[0].Val)
	assert.Equal(t, 1, len(node.Neighbors[0].Neighbors))
	assert.Equal(t, node, node.Neighbors[0].Neighbors[0])

	node = GraphFromEdgeArray([][]int{{1, 2}, {2, 3}, {1, 3}})
	assert.Equal(t, 1, node.Val)
	assert.Equal(t, 2, len(node.Neighbors))

	assert.Equal(t, 2, node.Neighbors[0].Val)
	assert.Equal(t, 2, len(node.Neighbors[0].Neighbors))

	assert.Equal(t, 2, len(node.Neighbors[1].Neighbors))
	assert.Equal(t, 3, node.Neighbors[1].Val)
}

func TestGraphToEdgeArray(t *testing.T) {
	var node *Node

	node = GraphFromEdgeArray([][]int{{1, 2}})
	assert.True(t, GraphArrayEqual([][]int{{1, 2}}, GraphToEdgeArray(node)))

	node = GraphFromEdgeArray([][]int{{1, 2}, {2, 3}, {3, 4}})
	assert.True(t, GraphArrayEqual([][]int{{1, 2}, {2, 3}, {3, 4}}, GraphToEdgeArray(node)))

	node = GraphFromEdgeArray([][]int{{1, 2}, {1, 4}, {2, 3}, {3, 4}})
	assert.True(t, GraphArrayEqual([][]int{{1, 2}, {1, 4}, {2, 3}, {3, 4}}, GraphToEdgeArray(node)))
}
