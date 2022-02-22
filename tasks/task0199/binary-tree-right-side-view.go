package task0199

import . "leetcode/binaryTree"

// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 100].
//   -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
//   Memory Usage: 2.4 MB, less than 23.61% of Go online submissions for Binary Tree Right Side View.
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	levelsDone := map[int]bool{}
	traverseReversed(root, 0, &levelsDone, &result)
	return result
}

func traverseReversed(node *TreeNode, level int, levelsDone *map[int]bool, res *[]int) {
	if node == nil {
		return
	}
	if !(*levelsDone)[level] {
		(*levelsDone)[level] = true
		*res = append(*res, node.Val)
	}
	// Traverse the right side first, because it will reduce the number of levelsDone writes
	traverseReversed(node.Right, level+1, levelsDone, res)
	traverseReversed(node.Left, level+1, levelsDone, res)
}
