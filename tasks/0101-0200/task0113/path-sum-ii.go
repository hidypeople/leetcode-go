package task0113

import . "leetcode/binaryTree"

// Given the root of a binary tree and an integer targetSum, return all root-to-leaf
// paths where the sum of the node values in the path equals targetSum. Each path
// should be returned as a list of the node values, not node references.
// A root-to-leaf path is a path starting from the root and ending at any
// leaf node. A leaf is a node with no children.
//
// 	Constraints:
//   The number of nodes in the tree is in the range [0, 5000].
//   -1000 <= Node.val <= 1000
//   -1000 <= targetSum <= 1000
//
// Results:
//   Runtime: 4 ms, faster than 90.56% of Go online submissions for Path Sum II.
//   Memory Usage: 4.4 MB, less than 90.99% of Go online submissions for Path Sum II.
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	path := make([]int, 5000)
	traverse(root, &path, 0, &result, targetSum)
	return result
}

func traverse(node *TreeNode, path *[]int, pathIdx int, result *[][]int, targetSum int) {
	(*path)[pathIdx] = node.Val
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		pathSliceCopy := make([]int, pathIdx+1)
		copy(pathSliceCopy, (*path)[:pathIdx+1])
		*result = append(*result, pathSliceCopy)
		return
	}
	if node.Left != nil {
		traverse(node.Left, path, pathIdx+1, result, targetSum-node.Val)
	}
	if node.Right != nil {
		traverse(node.Right, path, pathIdx+1, result, targetSum-node.Val)
	}
}
