package task1008

import . "leetcode/binaryTree"

// Given an array of integers preorder, which represents the preorder traversal of a BST
// (i.e., binary search tree), construct the tree and return its root.
// It is guaranteed that there is always possible to find a binary search tree with
// the given requirements for the given test cases.
// A binary search tree is a binary tree where for every node, any descendant of Node.left
// has a value strictly less than Node.val, and any descendant of Node.right has a value
// strictly greater than Node.val.
// A preorder traversal of a binary tree displays the value of the node first, then traverses
// Node.left, then traverses Node.right.
//
// Constraints:
//   1 <= preorder.length <= 100
//   1 <= preorder[i] <= 10^8
//   All the values of preorder are unique.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Construct Binary Search Tree from Preorder Traversal.
//   Memory Usage: 2.5 MB, less than 11.11% of Go online submissions for Construct Binary Search Tree from Preorder Traversal.
func bstFromPreorder(preorder []int) *TreeNode {
	// Complexity O(n)
	// We use the morris traversal technique: unoccupied right nodes contains the link to previous upper level node
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	traverse(&preorder, 1, root)
	return root
}

func traverse(preorder *[]int, preorderIdx int, currentNode *TreeNode) {
	if preorderIdx >= len(*preorder) {
		// Clean up right nodes that were used as a links to the "previous" parent nodes
		for currentNode.Right != nil {
			parent := currentNode.Right
			currentNode.Right = nil
			currentNode = parent
		}
		return
	}
	currentValue := (*preorder)[preorderIdx]
	preorderIdx++

	if currentValue < currentNode.Val {
		// Current value is less than node value -> go left
		// plus keep link to the previous node in the Right
		currentNode.Left = &TreeNode{
			Val:   currentValue,
			Right: currentNode,
		}
		traverse(preorder, preorderIdx, currentNode.Left)
	} else {
		// Current value is greater than node value -> we need to go upper using path and
		// Go upper using Right links and find the proper place
		for currentNode.Right != nil && currentNode.Right.Val < currentValue {
			temp := currentNode
			currentNode = currentNode.Right
			temp.Right = nil
		}
		prevNode := currentNode.Right
		currentNode.Right = &TreeNode{
			Val:   currentValue,
			Right: prevNode,
		}
		traverse(preorder, preorderIdx, currentNode.Right)
	}
}
