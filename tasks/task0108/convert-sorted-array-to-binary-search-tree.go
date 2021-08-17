package task0108

import . "leetcode/binaryTree"

type Slice struct {
	l    int
	r    int
	node *TreeNode
}

// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree. A height-balanced binary
// tree is a binary tree in which the depth of the two subtrees of every node
// never differs by more than one.
// Constraints:
//   1 <= nums.length <= 104
//   -104 <= nums[i] <= 104
//   nums is sorted in a strictly increasing order.
func sortedArrayToBST(nums []int) *TreeNode {
	// We'll solve it using queue (also it is possible to use recursive function calls)
	slice := Slice{
		0,
		len(nums) - 1,
		&TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	rootNode := slice.node

	queue := []Slice{slice}
	for len(queue) > 0 {
		// dequeue
		currentSlice := queue[0]
		queue = queue[1:]

		// calc middle (always exists)
		middle := (currentSlice.r-currentSlice.l)/2 + currentSlice.l

		// Set up curr val
		currentSlice.node.Val = nums[middle]

		// Enqueue left and right parts
		if currentSlice.l < middle {
			leftNode := &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}
			currentSlice.node.Left = leftNode
			queue = append(queue, Slice{currentSlice.l, middle - 1, leftNode})
		}
		if middle < currentSlice.r {
			rightNode := &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}
			currentSlice.node.Right = rightNode
			queue = append(queue, Slice{middle + 1, currentSlice.r, rightNode})
		}
	}
	return rootNode
}
