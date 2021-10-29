package task0107

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the bottom-up level order traversal of its nodes'
// values. (i.e., from left to right, level by level from leaf to root).
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 2000].
//   -1000 <= Node.val <= 1000
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
//   Memory Usage: 3.1 MB, less than 63.73% of Go online submissions for Binary Tree Level Order Traversal II.
func levelOrderBottom(root *TreeNode) [][]int {
	layers := [][]int{}
	if root == nil {
		return layers
	}
	// traverse tree and collect layers from top to bottom
	traverse(&layers, root, 0)

	// Reverse the array
	reversedLayers := make([][]int, len(layers))
	for i := 0; i < len(layers); i++ {
		reversedLayers[len(layers)-i-1] = layers[i]
	}
	return reversedLayers
}

func traverse(layers *[][]int, node *TreeNode, level int) {
	if len(*layers) < level+1 {
		*layers = append(*layers, []int{})
	}
	(*layers)[level] = append((*layers)[level], node.Val)
	if node.Left != nil {
		traverse(layers, node.Left, level+1)
	}
	if node.Right != nil {
		traverse(layers, node.Right, level+1)
	}
}
