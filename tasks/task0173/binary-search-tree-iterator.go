package task0173

import . "leetcode/binaryTree"

// Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):
//   BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor.
//                              The pointer should be initialized to a non-existent number smaller than any element in the BST.
//   boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
//   int next() Moves the pointer to the right, then returns the number at the pointer.
// Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.
// You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 10^5].
//   0 <= Node.val <= 10^6
//   At most 10^5 calls will be made to hasNext, and next.
//
// Results:
//   Runtime: 13 ms, faster than 100.00% of Go online submissions for Binary Search Tree Iterator.
//   Memory Usage: 9.7 MB, less than 94.09% of Go online submissions for Binary Search Tree Iterator.
type BSTIterator struct {
	stack []*TreeNode
	curr  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: make([]*TreeNode, 0),
		curr:  root,
	}
}

func (iter *BSTIterator) Next() int {
	for iter.curr != nil {
		// Go to the most left item
		iter.stack = append(iter.stack, iter.curr)
		iter.curr = iter.curr.Left
	}
	// Get top of the stack: this is current value
	iter.curr = iter.stack[len(iter.stack)-1]
	iter.stack = iter.stack[:len(iter.stack)-1]
	val := iter.curr.Val
	// go to the right branch for the next move
	iter.curr = iter.curr.Right
	return val
}

func (iter *BSTIterator) HasNext() bool {
	return iter.curr != nil || len(iter.stack) > 0
}
