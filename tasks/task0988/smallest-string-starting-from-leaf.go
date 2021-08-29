package task0988

import (
	. "leetcode/binaryTree"
	"strings"
)

var minPath *[]int = nil

// You are given the root of a binary tree where each node has a value in the range [0, 25] representing the letters 'a' to 'z'.
// Return the lexicographically smallest string that starts at a leaf of this tree and ends at the root.
// As a reminder, any shorter prefix of a string is lexicographically smaller.
// For example, "ab" is lexicographically smaller than "aba".
// A leaf of a node is a node that has no children.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 8500].
//   0 <= Node.val <= 25
//
// Results:
//   Runtime: 4 ms, faster than 96.15% of Go online submissions for Smallest String Starting From Leaf.
//   Memory Usage: 6.3 MB, less than 7.69% of Go online submissions for Smallest String Starting From Leaf.
func smallestFromLeaf(node *TreeNode) string {
	if node == nil {
		return ""
	}
	minPath = nil
	traverse(node, []int{})
	builder := new(strings.Builder)
	for i := len(*minPath) - 1; i >= 0; i-- {
		builder.WriteByte(byte((*minPath)[i]) + 'a')
	}
	return builder.String()
}

func traverse(node *TreeNode, currentPath []int) {
	path := currentPath
	if node.Left != nil && node.Right != nil {
		path = make([]int, len(currentPath)+1)
		copy(path, currentPath)
		path[len(path)-1] = node.Val
	} else {
		path = append(path, node.Val)
	}

	if node.Left == nil && node.Right == nil {
		// came to the end - compare with minPath:
		minPath = getMinPath(minPath, &path)
		return
	}
	if node.Left != nil {
		traverse(node.Left, path)
	}
	if node.Right != nil {
		traverse(node.Right, path)
	}
}

// Compare 2 strings lexicographically:
// a > b  => 1
// a == b => 0
// a < b  => -1
func getMinPath(a, b *[]int) *[]int {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	aa, bb := *a, *b
	nA, nB := len(aa), len(bb)
	i, j := nA-1, nB-1
	for i >= 0 && j >= 0 {
		if aa[i] < bb[j] {
			return a
		} else if aa[i] > bb[j] {
			return b
		}
		i--
		j--
	}
	if i == -1 {
		return a
	}
	return b
}
