package task0034

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.
//
// Constraints:
//   0 <= nums.length <= 10^5
//   -10^9 <= nums[i] <= 10^9
//   nums is a non-decreasing array.
//   -10^9 <= target <= 10^9
//
// Results:
//   Runtime: 4 ms, faster than 98.77% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//   Memory Usage: 4 MB, less than 19.70% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	low := findLow(nums, target)
	if low < 0 {
		return []int{-1, -1}
	}
	return []int{low, findHigh(nums, target, low)}
}

// Find left index of slice [target, target, ..., target]
func findLow(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
			if left > len(nums)-1 {
				break
			}
		} else if nums[middle] >= target {
			right = middle - 1
		}
		if nums[left] == target {
			result = left
		}
	}
	return result
}

// Find right index of slice [target, target, ..., target]
func findHigh(nums []int, target int, startIdx int) int {
	left, right := startIdx, len(nums)-1
	result := -1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] <= target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		}
		if nums[right] == target {
			result = right
		}
	}
	return result
}
