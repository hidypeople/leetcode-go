package task0111

import . "leetcode/binaryTree"

// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
// Note: A leaf is a node with no children.
//
// Constraints:
//   The number of nodes in the tree is in the range [0, 10^5].
//   -1000 <= Node.val <= 1000
//
// Results:
//   Runtime: 180 ms, faster than 98.35% of Go online submissions for Minimum Depth of Binary Tree.
//   Memory Usage: 19.1 MB, less than 73.27% of Go online submissions for Minimum Depth of Binary Tree.
type QueueItem struct {
	Node  *TreeNode
	Level int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []QueueItem{{root, 1}}
	for len(queue) > 0 {
		queueItem := queue[0]
		queue = queue[1:]
		if queueItem.Node.Left == nil && queueItem.Node.Right == nil {
			return queueItem.Level
		}
		if queueItem.Node.Left != nil {
			queue = append(queue, QueueItem{queueItem.Node.Left, queueItem.Level + 1})
		}
		if queueItem.Node.Right != nil {
			queue = append(queue, QueueItem{queueItem.Node.Right, queueItem.Level + 1})
		}
	}
	return 0
}
