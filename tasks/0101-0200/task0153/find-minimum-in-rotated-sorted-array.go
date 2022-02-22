package task0153

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
// For example, the array nums = [0,1,2,4,5,6,7] might become:
//   [4,5,6,7,0,1,2] if it was rotated 4 times.
//   [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
// You must write an algorithm that runs in O(log n) time.
//
// Constraints:
//   n == nums.length
//   1 <= n <= 5000
//   -5000 <= nums[i] <= 5000
//   All the integers of nums are unique.
//   nums is sorted and rotated between 1 and n times.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Minimum in Rotated Sorted Array.
//   Memory Usage: 2.5 MB, less than 25.77% of Go online submissions for Find Minimum in Rotated Sorted Array.
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		left, right, middle := nums[l], nums[r], nums[m]
		if left > middle {
			// move right
			r = m
		} else {
			if middle < right {
				// special case (order without shift): [1,2,3,4,5]
				return left
			}

			// move left
			if l != m {
				l = m + 1
				continue
			}

			// if l == m: we have 2 or 1 elements
			if left < right {
				return left
			} else {
				return right
			}
		}
	}
	return nums[r]
}
