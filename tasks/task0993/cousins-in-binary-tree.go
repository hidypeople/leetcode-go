package task0993

import . "leetcode/binaryTree"

// Given the root of a binary tree with unique values and the values of two different nodes of the
// tree x and y, return true if the nodes corresponding to the values x and y in the tree are
// cousins, or false otherwise.
// Two nodes of a binary tree are cousins if they have the same depth with different parents.
// Note that in a binary tree, the root node is at the depth 0, and children of each depth k
// node are at the depth k + 1.
//
// Constraints:
//   The number of nodes in the tree is in the range [2, 100].
//   1 <= Node.val <= 100
//   Each node has a unique value.
//   x != y
//   x and y are exist in the tree.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Cousins in Binary Tree.
//   Memory Usage: 2.7 MB, less than 47.06% of Go online submissions for Cousins in Binary Tree.
func isCousins(root *TreeNode, x int, y int) bool {
	dXY := []int{-1, -1}
	if !findNodes(root, x, y, 0, &dXY) {
		return false
	}
	return dXY[0] == dXY[1]
}

func findNodes(node *TreeNode, x, y, depth int, dXY *[]int) bool {
	if node == nil || (*dXY)[0] >= 0 && (*dXY)[1] >= 0 {
		return true
	} else if node != nil && node.Left != nil && node.Right != nil {
		if (node.Left.Val == x && node.Right.Val == y) || (node.Left.Val == y && node.Right.Val == x) {
			// brothers
			return false
		}
	}
	if node.Val == x {
		(*dXY)[0] = depth
	} else if node.Val == y {
		(*dXY)[1] = depth
	} else {
		if !findNodes(node.Left, x, y, depth+1, dXY) {
			return false
		}
		if !findNodes(node.Right, x, y, depth+1, dXY) {
			return false
		}
	}
	return true
}
