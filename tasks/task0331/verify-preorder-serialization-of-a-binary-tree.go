package task0331

import "strings"

// One way to serialize a binary tree is to use preorder traversal. When we encounter a non-null node,
// we record the node's value. If it is a null node, we record using a sentinel value such as '#'.
//     9
//    / \
//   3   2
//  / \   \
// 4   1   6
// For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", where '#' represents a null node.
// Given a string of comma-separated values preorder, return true if it is a correct preorder traversal serialization of a binary tree.
// It is guaranteed that each comma-separated value in the string must be either an integer or a character '#' representing null pointer.
// You may assume that the input format is always valid.
// For example, it could never contain two consecutive commas, such as "1,,3".
// Note: You are not allowed to reconstruct the tree.
//
// Constraints:
//   1 <= preorder.length <= 104
//   preorder consist of integers in the range [0, 100] and '#' separated by commas ','.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Verify Preorder Serialization of a Binary Tree.
//   Memory Usage: 2.8 MB, less than 84.00% of Go online submissions for Verify Preorder Serialization of a Binary Tree.
func isValidSerialization(preorder string) bool {
	s := strings.Split(preorder, ",")
	resultIdx, ok := traverse(s, 0, 0)
	return ok && resultIdx == len(s)-1
}

func traverse(s []string, currIdx int, depth int) (int, bool) {
	if currIdx > len(s)-1 {
		return 0, false
	}
	isNil := s[currIdx] == "#"
	if !isNil {
		resultIdx, ok := traverse(s, currIdx+1, depth+1)
		if !ok {
			return 0, false
		} else {
			resultIdx2, ok2 := traverse(s, resultIdx+1, depth+1)
			if !ok2 {
				return 0, false
			}
			return resultIdx2, true
		}
	} else {
		return currIdx, true
	}
}
