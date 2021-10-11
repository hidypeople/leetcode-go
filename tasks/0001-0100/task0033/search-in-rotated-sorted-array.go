package task0033

// There is an integer array nums sorted in ascending order (with distinct values).
// Prior to being passed to your function, nums is rotated at an unknown pivot index k
// (0 <= k < nums.length) such that the resulting array is
// [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
// Given the array nums after the rotation and an integer target, return the index
// of target if it is in nums, or -1 if it is not in nums.
// You must write an algorithm with O(log n) runtime complexity.
//
// Constraints:
//   1 <= nums.length <= 5000
//   -10^4 <= nums[i] <= 10^4
//   All values of nums are unique.
//   nums is guaranteed to be rotated at some pivot.
//   -10^4 <= target <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
//   Memory Usage: 2.6 MB, less than 14.60% of Go online submissions for Search in Rotated Sorted Array.
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	} else if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// (n >= 2)
	left, right := 0, n-1
	for left < right {
		middle := left + (right-left)/2
		if nums[left] < nums[middle] && nums[middle] < nums[right] {
			return binarySearch(nums, target)
		} else if nums[middle] >= nums[left] {
			// Left part is increasing, right part is mixed
			if target >= nums[left] && target <= nums[middle] {
				// target inside of increasing left part: simple search
				res := binarySearch(nums[:middle+1], target)
				if res >= 0 {
					return res
				} else {
					return -1
				}
			} else {
				// target inside of mixed right part: recursive search
				res := search(nums[middle+1:], target)
				if res >= 0 {
					return middle + 1 + res
				} else {
					return -1
				}
			}
		} else {
			// Left part is mixed, right part is increasing
			if target >= nums[middle] && target <= nums[right] {
				// target inside of increasing right part: simple search
				res := binarySearch(nums[middle:], target)
				if res >= 0 {
					return middle + res
				} else {
					return -1
				}
			} else {
				// target inside of mixed left part: recursive search
				res := search(nums[:middle], target)
				if res >= 0 {
					return res
				} else {
					return -1
				}
			}
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if len(nums) == 0 {
		return -1
	}
	for l < r {
		m := l + (r-l)/2
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			return m
		}
	}
	if r < 0 || l >= len(nums) {
		return -1
	}
	if nums[l] == target {
		return l
	} else if nums[r] == target {
		return r
	} else {
		return -1
	}
}
