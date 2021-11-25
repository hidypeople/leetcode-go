package graph

import "sort"

type Node struct {
	Val       int
	Neighbors []*Node
}

func GraphFromNeighborIdxArray(neighborArr [][]int) *Node {
	n := len(neighborArr)
	if n == 0 {
		return nil
	}
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{Val: i + 1, Neighbors: []*Node{}}
	}
	for i := 0; i < n; i++ {
		node := nodes[i]
		for _, neighborIdx := range neighborArr[i] {
			node.Neighbors = append(node.Neighbors, nodes[neighborIdx-1])
		}
	}
	return nodes[0]
}

func GraphToNeighborIdxArray(root *Node) [][]int {
	edges := GraphToEdgeArray(root)
	vertsMap := map[int]bool{}
	for _, edge := range edges {
		vertsMap[edge[0]] = true
		vertsMap[edge[1]] = true
	}
	verts := make([]int, len(vertsMap))
	i := 0
	for vert := range verts {
		verts[i] = vert
		i++
	}
	sort.Ints(verts)
	result := make([][]int, len(verts))
	for i = 1; i <= len(result); i++ {
		neighbors := []int{}
		for _, edge := range edges {
			if edge[0] == i {
				neighbors = append(neighbors, edge[1])
			}
			if edge[1] == i {
				neighbors = append(neighbors, edge[0])
			}
		}
		sort.Ints(neighbors)
		result[i-1] = neighbors
	}
	return result
}

func GraphFromEdgeArray(edges [][]int) *Node {
	if len(edges) == 0 {
		return nil
	}
	var root *Node = nil
	done := map[int]*Node{}
	for _, edge := range edges {
		if len(edge) != 2 {
			continue
		}
		node1, ok1 := done[edge[0]]
		node2, ok2 := done[edge[1]]
		if !ok1 {
			node1 = &Node{
				Val:       edge[0],
				Neighbors: []*Node{},
			}
			if root == nil {
				root = node1
			}
			done[node1.Val] = node1
		}
		if !ok2 {
			node2 = &Node{
				Val:       edge[1],
				Neighbors: []*Node{},
			}
			done[node2.Val] = node2
		}
		node1.Neighbors = append(node1.Neighbors, node2)
		node2.Neighbors = append(node2.Neighbors, node1)
	}
	return root
}

func GraphToEdgeArray(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	done := map[int]bool{}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range node.Neighbors {
			if !done[neighbor.Val] {
				result = append(result, []int{node.Val, neighbor.Val})
				stack = append(stack, neighbor)
			}
		}
		done[node.Val] = true
	}
	return result
}

func GraphEqual(node1, node2 *Node) bool {
	return GraphArrayEqual(GraphToEdgeArray(node1), GraphToEdgeArray(node2))
}

func GraphArrayEqual(arr1, arr2 [][]int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	if len(arr1) == 0 {
		return true
	}
	edgesLeft := make([][]int, len(arr2))
	for i, a2 := range arr2 {
		edgesLeft[i] = []int{a2[0], a2[1]}
	}
	for _, a1 := range arr1 {
		foundIndex := -1
		for i, a2 := range edgesLeft {
			if (a1[0] == a2[0] && a1[1] == a2[1]) || (a1[0] == a2[1] && a1[1] == a2[0]) {
				foundIndex = i
				break
			}
		}
		if foundIndex < 0 {
			return false
		}
		edgesLeft = append(edgesLeft[:foundIndex], edgesLeft[foundIndex+1:]...)
	}
	return true
}
