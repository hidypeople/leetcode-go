package task0162

// A peak element is an element that is strictly greater than its neighbors.
// Given an integer array nums, find a peak element, and return its index.
// If the array contains multiple peaks, return the index to any of the peaks.
// You may imagine that nums[-1] = nums[n] = -∞.
// You must write an algorithm that runs in O(log n) time.
//
// Constraints:
//   1 <= nums.length <= 1000
//   -2^31 <= nums[i] <= 2^31 - 1
// nums[i] != nums[i + 1] for all valid i.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Peak Element.
//   Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Find Peak Element.
func findPeakElement(nums []int) int {
	// Binary search: O(lnN)
	l, r := 0, len(nums)-1
	var m1 int
	var m2 int
	for l < r {
		m1 = (r + l - 1) >> 1
		m2 = m1 + 1
		if nums[m1] < nums[m2] {
			l = m2
		} else {
			r = m1
		}
	}
	return l
}
