package task0081

// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
// such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
// (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
// Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
// You must decrease the overall operation steps as much as possible.
//
// Constraints:
//   1 <= nums.length <= 5000
//   -10^4 <= nums[i] <= 10^4
//   nums is guaranteed to be rotated at some pivot.
//   -10^4 <= target <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array II.
//   Memory Usage: 3.2 MB, less than 45.83% of Go online submissions for Search in Rotated Sorted Array II.
func search(nums []int, target int) bool {
	// Complexity O(ln(N))
	// Binary search in rotated array:
	// very similar to ordinary binary search, but we need to take care about min/max values that
	// not neccessary on the first and last positions
	left, right := 0, len(nums)-1
	for left < right {
		middle := (left + right) >> 1
		if middle == left {
			return nums[left] == target || nums[right] == target
		}
		numMiddle := nums[middle]
		numLeft, numRight := nums[left], nums[right]
		if target == numMiddle || target == numLeft || target == numRight {
			return true
		}
		if numLeft == numMiddle && numMiddle == numRight {
			for i := left + 1; i <= right-1; i++ {
				if nums[i] != numLeft {
					return nums[i] == target
				}
			}
			return false
		} else if numLeft <= numMiddle && numMiddle <= numRight {
			// Search with zero rotation
			return searchStraight(nums, left, right, target)
		} else if numLeft == numMiddle {
			// [1,1,1,2,3] => [1,2,3]
			left = middle
		} else if numRight == numMiddle {
			// [1,2,3,3,3] => [1,2,3]
			right = middle
		} else if numLeft < numMiddle {
			// [2,3,4,5,1]
			if target < numMiddle && target > numLeft {
				return searchStraight(nums, left, middle, target)
			} else {
				left = middle
			}
		} else {
			// [5,1,2,3,4]
			if target > numMiddle && target < numRight {
				return searchStraight(nums, middle, right, target)
			} else {
				right = middle - 1
			}
		}
	}
	return nums[left] == target || nums[right] == target
}

func searchStraight(nums []int, left, right int, target int) bool {
	if target < nums[left] || target > nums[right] {
		return false
	}
	for left <= right {
		middle := left + (right-left)>>1
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return true
		}
	}
	return nums[left] == target
}
