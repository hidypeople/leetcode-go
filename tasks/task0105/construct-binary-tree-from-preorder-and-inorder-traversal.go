package task0105

import . "leetcode/binaryTree"

// Given two integer arrays preorder and inorder where preorder is the preorder
// traversal of a binary tree and inorder is the inorder traversal of the same
// tree, construct and return the binary tree.
//
// Constraints:
//   1 <= preorder.length <= 3000
//   inorder.length == preorder.length
//   -3000 <= preorder[i], inorder[i] <= 3000
//   preorder and inorder consist of unique values.
//   Each value of inorder also appears in preorder.
//   preorder is guaranteed to be the preorder traversal of the tree.
//   inorder is guaranteed to be the inorder traversal of the tree.
//
// Results:
//   Runtime: 4 ms, faster than 91.78% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
//   Memory Usage: 4.3 MB, less than 97.26% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Algorithm Inorder(tree)
	// 1. Traverse the left subtree, i.e., call Inorder(left-subtree)
	// 2. Visit the root.
	// 3. Traverse the right subtree, i.e., call Inorder(right-subtree)

	// Algorithm Preorder(tree)
	// 1. Visit the root.
	// 2. Traverse the left subtree, i.e., call Preorder(left-subtree)
	// 3. Traverse the right subtree, i.e., call Preorder(right-subtree)
	return traverse(0, 0, len(preorder), len(inorder), &preorder, &inorder)
}

func traverse(l1, l2, r1, r2 int, preorder, inorder *[]int) *TreeNode {
	// preorder is: l1->[root value, left values..., ... rightValues]<-r1
	//  inorder is: l2->[left values..., root value, ... rightValues]<-r2
	// The first value in preorder tree is root value
	val := (*preorder)[l1]

	// Find the root index in inorder value
	idx := 0
	for i := l2; i < r2; i++ {
		if (*inorder)[i] == val {
			break
		}
		idx++
	}

	// We can split the tree into left and right subtrees when we know the index
	var left *TreeNode = nil
	var right *TreeNode = nil
	if idx > 0 {
		left = traverse(l1+1, l2, l1+idx+1, l2+idx, preorder, inorder)
	}
	if l1+idx+1 < r1 {
		right = traverse(l1+idx+1, l2+idx+1, r1, r2, preorder, inorder)
	}
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}
