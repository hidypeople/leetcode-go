package binarytree

import (
	"fmt"
	mathInt "leetcode/mathInt"
)

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

// Print tree
func BSTPrint(root *TreeNode) {
	if root == nil {
		fmt.Printf("<nil>\n")
		return
	}
	fmt.Println(printCurr(root, 0))
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
