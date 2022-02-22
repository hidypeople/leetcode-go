package task0133

import (
	. "leetcode/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cloneGraph(t *testing.T) {
	graph := GraphFromNeighborIdxArray([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	clonedGraph := cloneGraph(graph)

	assert.NotSame(t, graph, clonedGraph)
	assert.NotSame(t, graph.Neighbors[0], clonedGraph.Neighbors[0])
	assert.NotSame(t, graph.Neighbors[0].Neighbors[0], clonedGraph.Neighbors[0].Neighbors[0])

	g1 := GraphToNeighborIdxArray(graph)
	g2 := GraphToNeighborIdxArray(clonedGraph)
	assert.Equal(t, g1, g2)
}

func Test_cloneGraph2(t *testing.T) {
	graph := GraphFromNeighborIdxArray([][]int{{}})
	clonedGraph := cloneGraph(graph)
	assert.Equal(t, GraphToNeighborIdxArray(graph), GraphToNeighborIdxArray(clonedGraph))
}

func Test_cloneGraph3(t *testing.T) {
	graph := GraphFromNeighborIdxArray([][]int{})
	clonedGraph := cloneGraph(graph)
	assert.Equal(t, GraphToNeighborIdxArray(graph), GraphToNeighborIdxArray(clonedGraph))
}

func Test_cloneGraph4(t *testing.T) {
	graph := GraphFromNeighborIdxArray([][]int{{2}, {1}})
	clonedGraph := cloneGraph(graph)
	assert.Equal(t, GraphToNeighborIdxArray(graph), GraphToNeighborIdxArray(clonedGraph))
}
