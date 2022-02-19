package binarytree

import (
	"fmt"
	mathInt "leetcode/mathInt"
)

var NULL = -1 << 63

// Convert int -> *int
func I(val int) *int {
	return &val
}

// binary tree with int values
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BST to string
func BSTToString(root *TreeNode) string {
	if root == nil {
		return "<nil>"
	}
	return printCurr(root, 0)
}

// Print tree
func BSTPrint(root *TreeNode) {
	fmt.Println(BSTToString(root))
}

// Convert array into height balanced BST
// [0, -3, 9, -10, nil, 5]
// (level0: 0 -> level1: -3,9 -> level2: -10,nil,5):
//           0
//          / \
//        -3   9
//        /   /
//      -10  5
func BSTFromArray(arr []*int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	return convertNode(arr, 0, 0)
}

func printCurr(node *TreeNode, level int) string {
	var leftResult, rightResult string = "", ""
	if node.Left != nil {
		leftResult = printCurr(node.Left, level+1)
	}
	if node.Right != nil {
		rightResult = printCurr(node.Right, level+1)
	}
	if node.Left != nil && node.Right != nil {
		return fmt.Sprintf("%v(l=%v | r=%v)", node.Val, leftResult, rightResult)
	} else if node.Left != nil {
		return fmt.Sprintf("%v(l=%v)", node.Val, leftResult)
	} else if node.Right != nil {
		return fmt.Sprintf("%v(r=%v)", node.Val, rightResult)
	} else {
		return fmt.Sprintf("%v", node.Val)
	}
}

func convertNode(arr []*int, i int, level int) *TreeNode {
	result := &TreeNode{
		Val:   *arr[i],
		Left:  nil,
		Right: nil,
	}
	currGroupSize := mathInt.Pow(2, level)
	levelIndex := i - (currGroupSize - 1)
	leftIndex := levelIndex*2 + (currGroupSize*2 - 1)
	rightIndex := leftIndex + 1
	if leftIndex < len(arr) && arr[leftIndex] != nil {
		result.Left = convertNode(arr, leftIndex, level+1)
	}
	if rightIndex < len(arr) && arr[rightIndex] != nil {
		result.Right = convertNode(arr, rightIndex, level+1)
	}
	return result
}

// Transorm array into binary tree, the empty children are NULL
func BSTFromArrayInts(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Transform binary tree to array of ints, the empty children will be NULL
func BSTToArrayInts(root *TreeNode) []int {
	res := make([]int, 0, 1024)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}
	return res[:i]
}

// Traverse through the binary tree without recursion example
func BSTPrintSortedArray(root *TreeNode) {
	// Morris traversal algorithm:
	// https://en.wikipedia.org/wiki/Tree_traversal#Morris_in-order_traversal_using_threading

	var current *TreeNode

	for root != nil {
		if root.Left != nil {
			// Look into left branch
			current = root.Left

			// Try to go right as deep as we can
			// If current tree branch has already been threaded: it will stop on the pre-root item,
			// otherwise it will stop on the deepest right leaf
			for current.Right != nil && current.Right != root {
				current = current.Right
			}

			if current.Right != nil {
				// Thread exist (current.Right == root):
				// break the thread and return go to the root.Right
				current.Right = nil

				// Do the logic:
				print(root.Val, " ")

				root = root.Right
			} else {
				// Make thread from the deepest right node to the current root
				// For BST: [..., current.Val, root.Val, ...]
				current.Right = root

				// Go to the left
				root = root.Left
			}
		} else {

			// Do the logic:
			print(root.Val, " ")

			// Go to the right
			root = root.Right
		}
	}
	println()
}
