package augustleetcodingchallenge2021

import . "leetcode/binaryTree"

////////////////////////////////////////////////////////////////////////////////////////////////
// Solution with recursion
////////////////////////////////////////////////////////////////////////////////////////////////

// Given a binary tree root, a node X in the tree is named good if in the path
// from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.
// Constraints:
//   The number of nodes in the binary tree is in the range [1, 10^5].
//   Each node's value is between [-10^4, 10^4].
func GoodNodes(root *TreeNode) int {
	return goodNodesCurrent(root, root.Val)
}

func goodNodesCurrent(node *TreeNode, level int) int {
	newLevel := level
	result := 0
	if node.Val >= level {
		result++
		newLevel = node.Val
	}
	if node.Left != nil {
		result += goodNodesCurrent(node.Left, newLevel)
	}
	if node.Right != nil {
		result += goodNodesCurrent(node.Right, newLevel)
	}
	return result
}

////////////////////////////////////////////////////////////////////////////////////////////////
// Solution without recursion
////////////////////////////////////////////////////////////////////////////////////////////////

type NodeAndLevel struct {
	node   *TreeNode
	maxVal int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Given a binary tree root, a node X in the tree is named good if in the path
// from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.
// Constraints:
//   The number of nodes in the binary tree is in the range [1, 10^5].
//   Each node's value is between [-10^4, 10^4].
func GoodNodes_NoRecursion(root *TreeNode) int {
	result := 1
	queue := []NodeAndLevel{
		{root, root.Val},
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.node.Left != nil {
			if current.node.Left.Val >= current.maxVal {
				result++
			}
			// Append left to queue with max level
			queue = append(
				queue,
				NodeAndLevel{
					current.node.Left,
					Max(current.maxVal, current.node.Left.Val),
				})
		}
		if current.node.Right != nil {
			if current.node.Right.Val >= current.maxVal {
				result++
			}
			// Append right to queue with max level
			queue = append(
				queue,
				NodeAndLevel{
					current.node.Right,
					Max(current.maxVal, current.node.Right.Val),
				})
		}
	}
	return result
}
