package task0450

import . "leetcode/binaryTree"

// Given a root node reference of a BST and a key, delete the node with the
// given key in the BST. Return the root node reference (possibly updated) of the BST.
// Basically, the deletion can be divided into two stages:
//   Search for a node to remove.
//   If the node is found, delete the node.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 10^4].
//   -10^5 <= Node.val <= 10^5
//   Each node has a unique value.
//   root is a valid binary search tree.
//   -10^5 <= key <= 10^5
//
// Results:
//   Runtime: 8 ms, faster than 98.43% of Go online submissions for Delete Node in a BST.
//   Memory Usage: 7.1 MB, less than 96.85% of Go online submissions for Delete Node in a BST.
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	superRoot := &TreeNode{Left: root}

	// Search node (in addition, we need to keep a parent node and parent-child direction)
	var foundNode *TreeNode = nil
	var foundNodeParent *TreeNode = superRoot
	var foundNodeIsLeftNode bool = true
	current := root
	for current != nil {
		if current.Val == key {
			foundNode = current
			break
		}
		foundNodeParent = current
		if key < current.Val {
			foundNodeIsLeftNode = true
			current = current.Left
		} else {
			foundNodeIsLeftNode = false
			current = current.Right
		}
	}
	if foundNode == nil {
		// If node not found => return root node as it is
		return root
	}

	var nodeForReplace *TreeNode = foundNode.Left
	if foundNode.Right != nil {
		// If we have right subtree => find the most left node. It'll be the node instead of found
		tempParent := foundNode
		temp := foundNode.Right
		for temp.Left != nil {
			tempParent = temp
			temp = temp.Left
		}
		if tempParent == foundNode {
			foundNode.Val = temp.Val
			foundNode.Right = temp.Right
			return superRoot.Left
		} else {
			tempParent.Left = temp.Right
			temp.Left = foundNode.Left
			temp.Right = foundNode.Right
		}
		nodeForReplace = temp
	}
	if foundNodeIsLeftNode {
		foundNodeParent.Left = nodeForReplace
	} else {
		foundNodeParent.Right = nodeForReplace
	}
	return superRoot.Left
}
