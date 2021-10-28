package task0106

import . "leetcode/binaryTree"

// Given two integer arrays inorder and postorder where inorder is the inorder traversal of
// a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.
//
// Constraints:
//   1 <= inorder.length <= 3000
//   postorder.length == inorder.length
//   -3000 <= inorder[i], postorder[i] <= 3000
//   inorder and postorder consist of unique values.
//   Each value of postorder also appears in inorder.
//   inorder is guaranteed to be the inorder traversal of the tree.
//   postorder is guaranteed to be the postorder traversal of the tree.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Construct Binary Tree from Inorder and Postorder Traversal.
//   Memory Usage: 4.4 MB, less than 65.65% of Go online submissions for Construct Binary Tree from Inorder and Postorder Traversal.
func buildTree(inorder []int, postorder []int) *TreeNode {
	// Algorithm Inorder(tree)
	// 1. Traverse the left subtree, i.e., call Inorder(left-subtree)
	// 2. Visit the root.
	// 3. Traverse the right subtree, i.e., call Inorder(right-subtree)

	// Algorithm Postorder(tree)
	// 1. Traverse the left subtree, i.e., call Preorder(left-subtree)
	// 2. Traverse the right subtree, i.e., call Preorder(right-subtree)
	// 3. Visit the root.
	indexCache := make(map[int]int, len(inorder))
	for i, v := range inorder {
		indexCache[v] = i
	}
	return traverse(0, 0, len(inorder), len(postorder), &inorder, &postorder, &indexCache)
}

func traverse(l1, l2, r1, r2 int, inorder, postorder *[]int, indexCache *map[int]int) *TreeNode {
	// inorder is:  l1->[left values..., root value, ... rightValues]<-r2
	// preorder is: l2->[left values..., ... rightValues, root value]<-r2
	// The first value in preorder tree is root value
	val := (*postorder)[r2-1]

	// Find the root index in inorder value
	idx := (*indexCache)[val] - l1

	// We can split the tree into left and right subtrees when we know the index
	var left *TreeNode = nil
	var right *TreeNode = nil
	if idx > 0 {
		left = traverse(l1, l2, l1+idx, l2+idx, inorder, postorder, indexCache)
	}
	if l1+idx+1 < r1 {
		right = traverse(l1+idx+1, l2+idx, r1, r2-1, inorder, postorder, indexCache)
	}
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}
