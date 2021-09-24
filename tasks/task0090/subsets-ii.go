package task0090

import "sort"

// Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Constraints:
//   1 <= nums.length <= 10
//   -10 <= nums[i] <= 10
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Subsets II.
//   Memory Usage: 2.5 MB, less than 99.73% of Go online submissions for Subsets II.
func subsetsWithDup(nums []int) [][]int {
	// Solution is the same as in task0078/subsets.go, but it has some small change
	result := [][]int{{}}
	sort.Ints(nums)
	subsetsRecursive(&result, nums, []int{}, 0)
	return result
}

func subsetsRecursive(result *[][]int, nums []int, current []int, start int) {
	currentDepth := len(current) + 1
	n := len(nums)
	// Here we need to check: whether we used current num or not.
	// e.g. if we have 1,2,2,2: we want to use only one "2" on the current position and cut off other possibilities
	used := make([]bool, 21)
	for i := start; i < n; i++ {
		num := nums[i]
		if used[num+10] {
			continue
		}
		used[num+10] = true
		newCurrent := make([]int, currentDepth)
		copy(newCurrent, current)
		newCurrent[currentDepth-1] = num
		*result = append(*result, newCurrent)

		subsetsRecursive(result, nums, newCurrent, i+1)
	}
}
