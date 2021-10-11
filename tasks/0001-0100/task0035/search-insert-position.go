package task0035

// Given a sorted array of distinct integers and a target value,
// return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
//
// Constraints:
//   1 <= nums.length <= 10^4
//   -10^4 <= nums[i] <= 10^4
//   nums contains distinct values sorted in ascending order.
//   -10^4 <= target <= 10^4
//
// Results:
//   Runtime: 0 ms, faster than 100.00% of Go online submissions for Search Insert Position.
//   Memory Usage: 3 MB, less than 13.94% of Go online submissions for Search Insert Position.
func searchInsert(nums []int, target int) int {

	// binary search
	left, right := 0, len(nums)-1
	for left <= right {
		if target < nums[left] {
			return left
		}
		if target > nums[right] {
			return right + 1
		}
		middle := left + (right-left)/2
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return 0
}
