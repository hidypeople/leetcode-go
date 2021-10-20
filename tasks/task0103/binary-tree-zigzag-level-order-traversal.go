package task0103

import . "leetcode/binaryTree"

// Given the root of a binary tree, return the zigzag level order traversal of its nodes'
// values. (i.e., from left to right, then right to left for the next level and alternate between).
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 2000].
//   -100 <= Node.val <= 100
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
// Memory Usage: 2.7 MB, less than 56.28% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{{root.Val}}
	subNodes := []*TreeNode{root}
	isLtr := false
	// Go layer by layer and switch the direction
	for len(subNodes) > 0 {
		subNodesNext := []*TreeNode{}
		for i := len(subNodes) - 1; i >= 0; i-- {
			subnode := subNodes[i]
			if isLtr {
				if subnode.Left != nil {
					subNodesNext = append(subNodesNext, subnode.Left)
				}
				if subnode.Right != nil {
					subNodesNext = append(subNodesNext, subnode.Right)
				}
			} else {
				if subnode.Right != nil {
					subNodesNext = append(subNodesNext, subnode.Right)
				}
				if subnode.Left != nil {
					subNodesNext = append(subNodesNext, subnode.Left)
				}
			}
		}
		if len(subNodesNext) > 0 {
			resultCurrent := make([]int, len(subNodesNext))
			for i, subNodeNext := range subNodesNext {
				resultCurrent[i] = subNodeNext.Val
			}
			result = append(result, resultCurrent)
		}
		subNodes = subNodesNext
		isLtr = !isLtr
	}
	return result
}
