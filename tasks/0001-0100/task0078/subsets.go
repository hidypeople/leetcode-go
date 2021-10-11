package task0078

// Given an integer array nums of unique elements, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Constraints:
//   1 <= nums.length <= 10
//   -10 <= nums[i] <= 10
//   All the numbers of nums are unique.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Subsets.
//   Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Subsets.
func subsets(nums []int) [][]int {
	result := [][]int{{}}
	subsetsRecursive(&result, nums, []int{}, 0)
	return result
}

func subsetsRecursive(result *[][]int, nums []int, current []int, start int) {
	currentDepth := len(current) + 1
	n := len(nums)
	for i := start; i < n; i++ {
		newCurrent := make([]int, currentDepth)
		copy(newCurrent, current)
		newCurrent[currentDepth-1] = nums[i]
		*result = append(*result, newCurrent)
		subsetsRecursive(result, nums, newCurrent, i+1)
	}
}
