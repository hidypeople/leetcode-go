package task0133

import . "leetcode/graph"

// Given a reference of a node in a connected undirected graph.
// Return a deep copy (clone) of the graph.
// Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.
// class Node {
//    public int val;
//    public List<Node> neighbors;
// }
// Test case format:
// For simplicity, each node's value is the same as the node's index (1-indexed).
// For example, the first node with val == 1, the second node with val == 2, and so on.
// The graph is represented in the test case using an adjacency list.
//
// An adjacency list is a collection of unordered lists used to represent a finite graph.
// Each list describes the set of neighbors of a node in the graph.
// The given node will always be the first node with val = 1. You must return the copy
// of the given node as a reference to the cloned graph.
//
// Constraints:
//   The number of nodes in the graph is in the range [0, 100].
//   1 <= Node.val <= 100
//   Node.val is unique for each node.
//   There are no repeated edges and no self-loops in the graph.
//   The Graph is connected and all nodes can be visited starting from the given node.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Clone Graph.
//   Memory Usage: 3 MB, less than 34.64% of Go online submissions for Clone Graph.
func cloneGraph(root *Node) *Node {
	// Using DFS
	if root == nil {
		return nil
	}
	var result *Node = nil
	nodeToClonedMap := map[*Node]*Node{}
	visited := map[*Node]bool{}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[node] {
			continue
		}
		visited[node] = true
		nodeCloned := &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
		for _, neighbor := range node.Neighbors {
			if neighborCloned, ok := nodeToClonedMap[neighbor]; ok {
				nodeCloned.Neighbors = append(nodeCloned.Neighbors, neighborCloned)
				neighborCloned.Neighbors = append(neighborCloned.Neighbors, nodeCloned)
			} else {
				stack = append(stack, neighbor)
			}
		}
		nodeToClonedMap[node] = nodeCloned
		if result == nil {
			result = nodeCloned
		}
	}
	return result
}
