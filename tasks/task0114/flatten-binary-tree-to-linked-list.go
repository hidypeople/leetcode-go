package task0114

import . "leetcode/binaryTree"

// Given the root of a binary tree, flatten the tree into a "linked list":
// The "linked list" should use the same TreeNode class where the right child
// pointer points to the next node in the list and the left child pointer is always null.
// The "linked list" should be in the same order as a pre-order traversal of the binary tree.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 2000].
//   -100 <= Node.val <= 100
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Flatten Binary Tree to Linked List.
//   Memory Usage: 3 MB, less than 24.77% of Go online submissions for Flatten Binary Tree to Linked List.
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root)
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil {
		// node.Left is empty => go deeper
		traverse(node.Right)
		return
	}
	if node.Right == nil {
		// node.Right is empty => attach left node to right and go deeper
		node.Left, node.Right = node.Right, node.Left
		traverse(node.Right)
		return
	}

	// Left and right nodes aren't empty:
	// 1. Save rightNode reference
	// 2. Attach left node to the right and make left node empty
	// 3. Do right node traverseTree
	// 4. Go to the end of traversed node, attach right node to the end and run traverse again
	rightNode := node.Right
	node.Right = node.Left
	node.Left = nil
	traverse(node.Right)
	current := node.Right
	for current.Right != nil {
		current = current.Right
	}
	current.Right = rightNode
	traverse(node.Right)
}
