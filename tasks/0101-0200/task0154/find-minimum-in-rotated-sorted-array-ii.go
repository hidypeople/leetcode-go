package task0154

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
// For example, the array nums = [0,1,4,4,5,6,7] might become:
//   [4,5,6,7,0,1,4] if it was rotated 4 times.
//   [0,1,4,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in
// the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
// Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.
// You must decrease the overall operation steps as much as possible.
//
// Constraints:
//   n == nums.length
//   1 <= n <= 5000
//   -5000 <= nums[i] <= 5000
//   nums is sorted and rotated between 1 and n times.
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Minimum in Rotated Sorted Array II.
//   Memory Usage: 3.4 MB, less than 11.71% of Go online submissions for Find Minimum in Rotated Sorted Array II.
func findMin(nums []int) int {
	// Binary search: O(lnN)
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		m := (l + r) >> 1
		if nums[r] < nums[m] {
			// if middle > right => [l...m...min...r]
			l = m + 1
		} else if nums[r] > nums[m] {
			// if middle < right => [l...min...m...r]
			r = m
		} else {
			// if middle == right => right border move left
			r--
		}
	}
	return nums[l]
}
