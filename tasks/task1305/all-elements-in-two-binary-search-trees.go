package task1305

import . "leetcode/binarytree"

// Given two binary search trees root1 and root2, return a list containing
// all the integers from both trees sorted in ascending order.
//
// Constraints:
//   The number of nodes in each tree is in the range [0, 5000].
//   -10^5 <= Node.val <= 10^5
//
// Results:
//   Runtime: 92 ms, faster than 98.55% of Go online submissions for All Elements in Two Binary Search Trees.
//   Memory Usage: 8.1 MB, less than 89.85% of Go online submissions for All Elements in Two Binary Search Trees.
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result := []int{}
	// 1 - means that we've just added to result element from first tree
	// 2 - means that we've just added to result element from second tree
	elementAddedFromTree := 0

	// inorder tree traversal implementation:
	// 1. go to the left and push items to stack
	// 2. extract top stack item and add its value to results
	// 3. add right node of current node to stack
	curr1, curr2, stack1, stack2, n1, n2 := root1, root2, []*TreeNode{}, []*TreeNode{}, 0, 0
	for len(stack1) > 0 || curr1 != nil || len(stack2) > 0 || curr2 != nil {
		// Go left for both trees:
		if elementAddedFromTree != 2 {
			for curr1 != nil {
				stack1 = append(stack1, curr1)
				curr1 = curr1.Left
			}
		}
		if elementAddedFromTree != 1 {
			for curr2 != nil {
				stack2 = append(stack2, curr2)
				curr2 = curr2.Left
			}
		}
		n1, n2 = len(stack1), len(stack2)
		if n1 == 0 {
			curr2 = stack2[n2-1]
			stack2 = stack2[:n2-1]
			result = append(result, curr2.Val)
			elementAddedFromTree = 2
			curr2 = curr2.Right
		} else if n2 == 0 {
			curr1 = stack1[n1-1]
			stack1 = stack1[:n1-1]
			result = append(result, curr1.Val)
			elementAddedFromTree = 1
			curr1 = curr1.Right
		} else {
			curr1 = stack1[n1-1]
			curr2 = stack2[n2-1]
			if curr1.Val < curr2.Val {
				stack1 = stack1[:n1-1]
				result = append(result, curr1.Val)
				elementAddedFromTree = 1
				curr1 = curr1.Right
			} else {
				stack2 = stack2[:n2-1]
				result = append(result, curr2.Val)
				elementAddedFromTree = 2
				curr2 = curr2.Right
			}
		}
	}
	return result
}
