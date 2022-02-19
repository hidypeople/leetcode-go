package task1675

import (
	"math"
	"sort"
)

// You are given an array nums of n positive integers.
// You can perform two types of operations on any element of the array any number of times:
//   If the element is even, divide it by 2.
//     For example, if the array is [1,2,3,4], then you can do this operation on the last element, and the array will be [1,2,3,2].
//   If the element is odd, multiply it by 2.
//     For example, if the array is [1,2,3,4], then you can do this operation on the first element, and the array will be [2,2,3,4].
// The deviation of the array is the maximum difference between any two elements in the array.
// Return the minimum deviation the array can have after performing some number of operations.
//
// Constraints:
//   n == nums.length
//   2 <= n <= 10^5
//   1 <= nums[i] <= 10^9
//
// Results:
//   Runtime: 204 ms, faster than 100.00% of Go online submissions for Minimize Deviation in Array.
//   Memory Usage: 9.9 MB, less than 83.33% of Go online submissions for Minimize Deviation in Array.
func minimumDeviation(nums []int) int {
	// Make all items even
	for i, num := range nums {
		if num%2 == 1 {
			nums[i] *= 2
		}
	}

	// Create binary tree set (Binary tree without duplicate elements)
	treeSet := ConstructorTreeSet(nums)
	deviation := math.MaxInt64
	for {
		// Calc current deviation
		maxVal := treeSet.Max()
		deviation = min(deviation, maxVal-treeSet.Min())
		if deviation > 0 && maxVal%2 == 0 {
			// If maxVal is even => remove it from the tree and add maxVal / 2
			treeSet.RemoveMax()
			treeSet.Add(maxVal / 2)
		} else {
			// If maxVal became odd - we cannot make deviation lesser => return
			break
		}
	}
	return deviation
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type TreeSetNode struct {
	Val   int
	Left  *TreeSetNode
	Right *TreeSetNode
}

type TreeSet struct {
	root *TreeSetNode
}

func ConstructorTreeSet(nums []int) TreeSet {
	if len(nums) == 0 {
		return TreeSet{
			root: nil,
		}
	}
	sort.Ints(nums)
	return TreeSet{
		root: TreeSetFromArray(&nums, 0, len(nums)-1),
	}
}

// Build tree set from array of integers
func TreeSetFromArray(nums *[]int, l, r int) *TreeSetNode {
	if l > r {
		return nil
	}
	m := l + (r-l)/2
	num := (*nums)[m]
	root := &TreeSetNode{
		Val:   num,
		Left:  nil,
		Right: nil,
	}
	ml, mr := m-1, m+1
	for ml >= 0 && (*nums)[ml] == num {
		ml--
	}
	for mr < len(*nums) && (*nums)[mr] == num {
		mr++
	}
	root.Left = TreeSetFromArray(nums, l, ml)
	root.Right = TreeSetFromArray(nums, mr, r)
	return root
}

// Remove existing element (it is supposed that this element exist)
func (_self *TreeSet) RemoveMax() {
	if _self.root.Right == nil {
		root := _self.root
		_self.root = root.Left
		root.Left = nil
		return
	}
	prev, curr := _self.root, _self.root.Right
	for curr.Right != nil {
		prev = curr
		curr = curr.Right
	}
	prev.Right = curr.Left
}

func (_self *TreeSet) Add(num int) {
	if _self.root.Val == num {
		return
	}
	if _self.root == nil {
		_self.root = &TreeSetNode{Val: num}
		return
	}
	var prev *TreeSetNode
	curr := _self.root
	isExisting := false
	for curr != nil {
		if num == curr.Val {
			isExisting = true
			break
		} else if num > curr.Val {
			prev = curr
			curr = curr.Right
		} else {
			prev = curr
			curr = curr.Left
		}
	}
	if isExisting {
		return
	}
	newNode := &TreeSetNode{Val: num}
	if num < prev.Val {
		prev.Left = newNode
	} else {
		prev.Right = newNode
	}
}

func (_self *TreeSet) Min() int {
	curr := _self.root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Val
}

func (_self *TreeSet) Max() int {
	curr := _self.root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Val
}
